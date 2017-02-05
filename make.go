package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var (
	emacsFile = filepath.Join(os.Getenv("HOME"), ".emacs")
	emacsDir  = filepath.Join(os.Getenv("HOME"), ".emacs.d", "lehmrob")

	// ErrNoFile There is no emacs file availible
	ErrNoFile = errors.New("No emacs file availible")
)

func isFileInstalled() bool {
	finfo, err := os.Stat(emacsFile)
	if err != nil {
		return false
	}

	return !finfo.IsDir()
}

func isDirInstalled() bool {
	dinfo, err := os.Stat(emacsDir)
	if err != nil {
		return false
	}

	return dinfo.IsDir()
}

func prepareForInstall() error {
	if isFileInstalled() {
		err := os.Remove(emacsFile)
		if err != nil {
			return nil
		}
	}

	if isDirInstalled() {
		err := os.RemoveAll(emacsDir)
		if err != nil {
			return nil
		}
	}

	return nil
}

func installFile() error {
	curDir, err := os.Getwd()
	if err != nil {
		return err
	}

	f := filepath.Join(curDir, "emacs")
	fInfo, err := os.Stat(f)
	if err != nil {
		return err
	}

	if fInfo.IsDir() {
		return ErrNoFile
	}

	content, err := ioutil.ReadFile(f)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(emacsFile, content, fInfo.Mode())
	if err != nil {
		return err
	}

	return nil
}

func copyFile(src, dest string) error {
	content, err := ioutil.ReadFile(src)
	if err != nil {
		return err
	}

	info, err := os.Stat(src)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(dest, content, info.Mode())
	if err != nil {
		return err
	}

	return nil
}

func installDir() error {
	err := os.MkdirAll(emacsDir, 0771)
	if err != nil {
		return err
	}

	curDir, err := os.Getwd()
	if err != nil {
		return err
	}

	cpySrc := filepath.Join(curDir, "emacs.d", "lehmrob")
	content, err := ioutil.ReadDir(cpySrc)
	if err != nil {
		return err
	}

	for _, elem := range content {
		pSrc := filepath.Join(cpySrc, elem.Name())
		pDest := filepath.Join(emacsDir, elem.Name())

		err = copyFile(pSrc, pDest)
		if err != nil {
			return err
		}
	}

	return nil
}

func install() error {
	err := prepareForInstall()
	if err != nil {
		return err
	}

	err = installFile()
	if err != nil {
		return err
	}

	err = installDir()
	if err != nil {
		return err
	}

	return nil
}

func grepElFiles() ([]string, error) {
	var elFiles []string

	err := filepath.Walk(emacsDir, func(p string, i os.FileInfo, e error) error {
		if strings.Contains(i.Name(), ".el") {
			elFiles = append(elFiles, i.Name())
		}
		return nil
	})

	return elFiles, err
}

func updateFiles(files []string) error {
	curDir, err := os.Getwd()
	if err != nil {
		return err
	}

	for _, file := range files {
		src := filepath.Join(emacsDir, file)
		dest := filepath.Join(curDir, "emacs.d/lehmrob", file)
		err = copyFile(src, dest)
		if err != nil {
			return err
		}
	}

	return nil
}

func update() error {
	elFiles, err := grepElFiles()
	if err != nil {
		return nil
	}

	for _, f := range elFiles {
		fmt.Println(f)
	}

	err = updateFiles(elFiles)
	if err != nil {

	}

	return nil
}

func main() {
	if len(os.Args) <= 1 {
		log.Fatal("Not enough arguments")
	}

	switch os.Args[1] {
	case "install":
		err := install()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Can't install emacs files:%s\n", err.Error())
		}
	case "update":
		err := update()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Can't update dotfiles:%s\n", err.Error())
		}
	}
}
