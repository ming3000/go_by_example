package main

import (
	"errors"
	"fmt"
	"strings"
)

// ------------------------------------
func errorCmd(cmd string) (string, error) {
	if cmd == "error" {
		return cmd, errors.New("i love the way you lie")
	}
	return cmd, nil
}

// ------------------------------------
type myError struct {
	code int
	message string
}

func (e * myError) Error() string {
	return fmt.Sprintf("error code: %d, error message: %s", e.code, e.message)
}

func errorCmd2(cmd string) (string , error) {
	if cmd == "error" {
		errObj := myError{code:500, message:"pressure, pushing down on me"}
		return cmd, &errObj
	}
	return cmd, nil
}

// ------------------------------------
func main() {
	_, err := errorCmd("haha")
	if err != nil {
		fmt.Println("errorCmd return error:", err)
	} else {
		fmt.Println("errorCmd return no error")
	}
	fmt.Println(strings.Repeat("-", 20))

	// best practice
	if _, err = errorCmd("error"); err != nil {
		fmt.Println("errorCmd return error:", err)
	} else {
		fmt.Println("errorCmd return no error")
	}
	fmt.Println(strings.Repeat("-", 20))

	_, err = errorCmd2("haha")
	if err != nil {
		fmt.Println("errorCmd2 return error:", err)
	} else {
		fmt.Println("errorCmd2 return no error")
	}
	fmt.Println(strings.Repeat("-", 20))

	// best practice
	if _, err = errorCmd2("error"); err != nil {
		fmt.Println("errorCmd2 return error:", err)
	} else {
		fmt.Println("errorCmd2 return no error")
	}
	fmt.Println(strings.Repeat("-", 20))
}
