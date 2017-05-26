package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	vimOrTreat()
}

func vimOrTreat() {
	sc := bufio.NewScanner(os.Stdin)

	fmt.Print("Vim or Treat?: ")
	sc.Scan()
	answer := sc.Text()

	if answer == "Happy Halloween!" {
		fmt.Println("Thanks!")
		return
	}

	err := startVim()
	if err != nil {
		fmt.Printf("failed starting vim: %s\n", err.Error())
	}
}

func startVim() error {
	vim := exec.Command("vim", os.Args[1:]...)
	vim.Stdin = os.Stdin
	vim.Stdout = os.Stdout
	vim.Stderr = os.Stderr

	return vim.Run()
}
