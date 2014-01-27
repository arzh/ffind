package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
//	"time"
)

type FileNameChecker struct {
	Name string

	Finds []string
}

func (c *FileNameChecker) Parse(arg string) {
	c.Name = arg
}

func (c *FileNameChecker) Add(file string) {
	c.Finds = append(c.Finds, file)
}

func Walker(path string, info os.FileInfo, err error) error {
	if info.IsDir() {
		return nil
	}

	//if checker.Matches(info.Name()) {
	//	checker.Add(path + info.Name());
	//}

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

	//fmt.Println(working_dir);

	fileToFind := os.Args[1]

	checker = new(FileNameChecker)
	checker.Parse(fileToFind)

	err = filepath.Walk(working_dir, Walker)
	if err != nil {
		fmt.Println(err.Error())
	}

	for i, e := range checker.Finds {
		str := strings.TrimLeft(e, working_dir)
		checker.Finds[i] = str;
		fmt.Fprintf(os.Stdout, "[%d]\t%s\n", i, str)
	}

	var fileIndex int
	_, err = fmt.Scanln(&fileIndex)
	if err != nil {
		fmt.Println("Failed to get input:", err.Error())
		return
	}

	fmt.Println("Index:",  fileIndex)

	if fileIndex > len(checker.Finds) {
		fmt.Println("Index out of range")
		return
	}

	openstr := checker.Finds[fileIndex]
	fmt.Println(openstr)

	cmd := exec.Command("cmd", "/C", openstr)
//	inpipe, err := cmd.StdinPipe()
//	if err != nil {
//		fmt.Println("Error getting stdin pipe:", err.Error())
//		return
//	}
	cmd.Dir = working_dir
	err = cmd.Start()
//	if err != nil {
//		fmt.Println("failed on start:", err.Error())
//		return
//	}

//	fmt.Println("cmd:", cmd)


	//inpipe.Write([]byte(openstr))
	//inpipe.Write([]byte("exit\n"))
	
//	n, err := fmt.Fprintln(inpipe, "main.go")
//	if err != nil {
//		fmt.Println("Failed printing to cmd inpipe", err.Error())
//		return;
//	}
//	fmt.Println("Bytes written:", n)
	
//	fmt.Println("Process state:", cmd.ProcessState)
	err = cmd.Process.Release()
	if err != nil {
		fmt.Println("Error on Kill", err.Error())
	}
//	err = cmd.Wait()
//	if err != nil {
//		fmt.Println("Wait failed:", err.Error())
//	}
}
