package main

import (
    "errors"
    "fmt"
    "io/ioutil"
	"os"
	"path/filepath"
)

var (
    emacsFile string = filepath.Join(os.Getenv("HOME"), ".emacs")
    emacsDir string = filepath.Join(os.Getenv("HOME"), ".emacs.d")
    ErrNoFile error = errors.New("No emacs file availible")
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
        return error
    }

    destStat, err := os.Stat(dest)
    if err != nil {
        return err
    }

    err = ioutil.WriteFile(dest, content, destStat.Mode())
    if err != nil {
        return err
    }

    return nil
}

func installDir() error {
    completeDir := filepath.Join(emacsDir, "lehmrob")
    err := os.MkdirAll(completeDir, 0771)
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
        pSrc := filepath.Join(cpySrc, elem)
        pDest := filepath.Join(completeDir, elem)

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

func main() {
	switch os.Args[1] {
	case "install":
        err := install()
        if err != nil {
            fmt.Fprintf(os.Stderr, "Can't install emacs files:%s\n", err.Error())
        }
	}
}
