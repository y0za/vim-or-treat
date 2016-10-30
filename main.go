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
		fmt.Println("failed starting vim: %s", err.Error())
	}
}

func startVim() error {
	vim := exec.Command("vim")
	vim.Stdin = os.Stdin
	vim.Stdout = os.Stdout

	return vim.Run()
}
