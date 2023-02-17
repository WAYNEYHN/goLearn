package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
)

func main() {

	for _, req := range []string{"", "hello"} {
		resp, err := echo(req)
		if err != nil {
			fmt.Printf("err: %s\n", err)
			judgeError(err)
			continue
		} else {
			fmt.Printf("%s\n", resp)
		}
	}

}

func echo(request string) (response string, myError error) {
	if request == "" {
		myError = errors.New("empty request")
		fmt.Errorf("%s")
		return
	}
	response = fmt.Sprintf("echo: %s", request)
	return
}

func judgeError(err error) error {
	switch err := err.(type) {
	case *os.PathError:
		fmt.Println(err)
		return err.Err
	case *os.LinkError:
		return err.Err
	case *os.SyscallError:
		return err.Err
	case *exec.Error:
		return err.Err
	}
	return err

}
