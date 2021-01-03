package pkgerutil

import (
	"fmt"
	"github.com/markbates/pkger"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func Root(realFileName string) string {
	splited := strings.SplitN(realFileName,"/", 3)
	return filepath.Join("/", splited[0], splited[1])
}

func RealPath(pkgerFileName string) string {
	splited := strings.Split(pkgerFileName, ":")
	return splited[1]
}

func ReadDir(realFileName string) ([]os.FileInfo, error) {
	getCurrentDepth := func(filePath string) int {
		splited := strings.Split(filePath, "/")
		return len(splited)
	}

	expectDepth := getCurrentDepth(realFileName) + 1

	var fileinfos []os.FileInfo
	err := pkger.Walk(realFileName, func(path string, info os.FileInfo, err error) error {
		currentDepth := getCurrentDepth(RealPath(path))
		if currentDepth == expectDepth {
			fileinfos = append(fileinfos, info)
		}
		return nil
	})

	return fileinfos, err
}

const dumpRootPath = "/tmp"

func Dump(realFileName string) error {

	return pkger.Walk(realFileName, func(path string, info os.FileInfo, err error) error {
		dumpPath := filepath.Join(dumpRootPath, RealPath(path))
		if info.IsDir() {
			if err := os.MkdirAll(dumpPath, 0755); err != nil {
				return err
			}
		} else {
			f, err := os.OpenFile(dumpPath, os.O_WRONLY|os.O_CREATE, 0755)
			if err != nil {
				return err
			}

			packedFile, err := pkger.Open(path)
			if err != nil {
				return err
			}
			io.Copy(f, packedFile)
		}
		return nil
	})
}

func ClearDump(realFileName string) error {
	dumpPath := filepath.Join(dumpRootPath, realFileName)
	return os.RemoveAll(dumpPath)
}

func Exec(executeFilepath string) {
	rootPath := Root(executeFilepath)
	err := Dump(rootPath)
	if err != nil {
		fmt.Println(err)
	}

	tempPath := filepath.Join(dumpRootPath, executeFilepath)

	// Execute command
	cmd := exec.Command(tempPath)
	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(output))

	err = ClearDump(rootPath)
	if err != nil {
		fmt.Println(err)
	}
}
