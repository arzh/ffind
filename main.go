package main

import (
	"fmt"
	"os"
	"os/exec"
	//"path"
	"path/filepath"
	"strconv"
	"strings"
	//	"time"
)

type FileNameChecker struct {
	Name string

	Finds []string
}

func (c *FileNameChecker) Parse(arg string) {
	c.Name = arg
	c.Finds = make([]string, 0)
}

func (c *FileNameChecker) Add(file string) {
	c.Finds = append(c.Finds, file)
}

func Walker(path string, info os.FileInfo, err error) error {
	if info.IsDir() {
		return nil
	}

	b, err := filepath.Match(checker.Name, info.Name())
	if b {
		checker.Add(path)
	} else if err != nil {
		return err
	}

	return nil
}

var checker *FileNameChecker

func main() {

	// Step1: Grab the first arg
	// Step2: parse it for begining or ending *
	// Step3: set up filewalker
	// Step4: print findings and wait for input
	// Step5: open file, I will need to have a list of types I want to always open in vim

	if len(os.Args) < 2 {
		fmt.Println("Must input file name to find") // TODO: make this more helpful
		return
	}

	working_dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println(err.Error())
	}

	fileToFind := os.Args[1]

	checker = new(FileNameChecker)
	checker.Parse(fileToFind)

	err = filepath.Walk(working_dir, Walker)
	if err != nil {
		fmt.Println(err.Error())
	}

	for i, e := range checker.Finds {
		str := strings.TrimPrefix(e, working_dir)
		//checker.Finds[i] = str
		fmt.Fprintf(os.Stdout, "[%d]\t%s\n", i, str)
	}

	// Now we wait...
	var scan1 string
	var scan2 string
	cmdName := "open"
	var index int64
	n, err := fmt.Scanln(&scan1, &scan2)
	fmt.Printf("n:%d scan1:%s scan2:%s \n", n, scan1, scan2)
	if UnexpectedError(err) {
		fmt.Println("Failed to get input:", err.Error())
		return
	}

	if n == 2 {
		cmdName = scan1
		index, _ = strconv.ParseInt(scan2, 10, 0)
	} else if n == 1 {
		index, _ = strconv.ParseInt(scan1, 10, 0)
	}

	fileindex := int(index)
	if fileindex > len(checker.Finds) {
		fmt.Println("File Index out of range")
		return
	}

	openstr := checker.Finds[fileindex]
	if cmdName == "cd" {
		// If cd is given we want to go to
		// the directory of the file
		fmt.Println(openstr)
		openstr, _ = filepath.Split(openstr)
		fmt.Println(openstr)
	}

	fmt.Println(cmdName, openstr)

	cmdProc := exec.Command(cmdName, openstr)
	cmdProc.Stdin = os.Stdin
	cmdProc.Stdout = os.Stdout
	cmdProc.Stderr = os.Stderr
	cmdProc.Run()

}

func UnexpectedError(err error) bool {
	if err == nil || err.Error() == "unexpected newline" {
		return false
	}

	return true
}
