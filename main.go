package main

import (
	"os"
	"path/filepath"
	"strings"
	"fmt"
)

type FileNameChecker struct {
	Pre, Post bool
	Name string
	
	Finds []string
}

func (c *FileNameChecker) Parse(arg string) {
	if strings.HasPrefix(arg, "*") {
		c.Pre = true;
	}

	if strings.HasSuffix(arg, "*") {
		c.Post = true;
	}

	c.Name = strings.Trim(arg, "*");
}

func (c *FileNameChecker) Add(file string) {
	c.Finds = append(c.Finds, file);
}

func (c *FileNameChecker) Matches(name string) bool {
	if checker.Pre && checker.Post {
		// If contains add
		if strings.Contains(name, c.Name) {
			return true;
		}
	} else if checker.Pre {
		// If Trails add
		if strings.HasSuffix(name, c.Name) {
			return true;
		}
	} else if checker.Post {
		// If Begins with add
		if strings.HasPrefix(name, c.Name) {
			return true;
		}
	} else if name == c.Name {
		return true;
	}

	return false;
}	

func Walker(path string, info os.FileInfo, err error) error {
	if info.IsDir() {
		return nil
	}

	if checker.Matches(info.Name()) {
		checker.Add(path + info.Name());
	}
	
	return nil;
}

var checker *FileNameChecker;

func main() {

	// Step1: Grab the first arg
	// Step2: parse it for begining or ending *
	// Step3: set up filewalker
	// Step4: print findings and wait for input
	// Step5: open file, I will need to have a list of types I want to always open in vim

	if len(os.Args) < 2 { 
		fmt.Println("Must input file name to find"); // TODO: make this more helpful
		return;
	}

   	working_dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println(err.Error());
	}

	fileToFind := os.Args[1];

	checker = new(FileNameChecker);
	checker.Parse(fileToFind);

	err = filepath.Walk(working_dir, Walker);
	if err != nil {
		fmt.Println(err.Error());
	}

	for _, e := range checker.Finds {
		fmt.Println(e);
	}
}
