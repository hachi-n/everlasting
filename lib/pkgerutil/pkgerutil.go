package pkgerutil

import (
	"fmt"
	"github.com/markbates/pkger"
	"io"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

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

func Exec(executeFilepath string) {
	const receiverCommand = "bash"

	cmdPipeGenerator := func(cmd *exec.Cmd) (io.WriteCloser, io.ReadCloser){
		wCloser, err := cmd.StdinPipe()
		if err != nil {
			fmt.Println(err)
		}

		rCloser, err := cmd.StdoutPipe()
		if err != nil {
			fmt.Println(err)
		}
		return wCloser, rCloser
	}

	f, err := pkger.Open(executeFilepath)
	if err != nil {
		fmt.Println(err)
	}

	cmd := exec.Command(receiverCommand)
	wCloser, rCloser := cmdPipeGenerator(cmd)

	// Execute command
	func () {
		defer wCloser.Close()

		err = cmd.Start()
		if err != nil {
			fmt.Println(err)
		}
		_, err = io.Copy(wCloser, f)
		if err != nil {
			fmt.Println(err)
		}
	}()

	// output result
	_, err = io.Copy(os.Stdout, rCloser)
	if err != nil {
		fmt.Println(err)
	}
	syscall.ForkExec()
}
