package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

/*
=== Взаимодействие с ОС ===

# Необходимо реализовать собственный шелл

встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды
конвеер на пайпах

Реализовать утилиту netcat (nc) клиент
принимать данные из stdin и отправлять в соединение (tcp/udp) "error"
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

const (
	cdCommand   = "cd"
	lsCommand   = "ls"
	pwdCommand  = "pwd"
	echoCommand = "echo"
	killCommand = "kill"
	quitCommand = "quit"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("$ ")
		cmdString, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		cmdString = strings.TrimSpace(cmdString)
		if cmdString == "" {
			continue
		}
		if cmdString == quitCommand {
			break
		}

		commands := strings.Split(cmdString, "|")
		fmt.Println(commands)
		err = runCommand(commands)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
	os.Exit(0)
}
func runCommand(commands []string) error {
	if len(commands) == 1 {

		args := strings.Fields(strings.TrimSpace(commands[0]))

		switch args[0] {
		case cdCommand:
			CdCommand(args)

		case lsCommand:
			LsCommand()
		case pwdCommand:
			PwdCommand()
		case echoCommand:
			EchoCommand(args)
		case killCommand:
			KillCommand(args)
		default:
			cmd := exec.Command(args[0], args[1:]...)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				return err
			}
		}

	}
	if len(commands) > 1 {
		err := executePipeline(commands)
		fmt.Fprintln(os.Stderr, err)
	}
	return nil
}
func executePipeline(commands []string) error {
	var prevOutput io.Reader
	var err error
	for _, command := range commands {
		args := strings.Fields(strings.TrimSpace(command))
		cmd := exec.Command(args[0], args[1:]...)

		if prevOutput != nil {
			cmd.Stdin = prevOutput
		}
		var output []byte
		output, err = cmd.Output()

		if err != nil {
			return err
		}

		fmt.Print(string(output))

		prevOutput = bytes.NewReader(output)
	}
	return nil
}
func CdCommand(args []string) {
	if len(args) == 1 {
		fmt.Print("usage: cd <dir>")
	}
	if len(args) == 2 {
		err := os.Chdir(args[1])
		if err != nil {
			fmt.Printf("cd: no such file or directory: %s\n", args[1])
		}
	}
}
func LsCommand() {
	dirs, err := os.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}
	for _, dir := range dirs {
		fmt.Print(dir.Name(), "        ")
	}
	fmt.Println("")
}
func PwdCommand() {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(path)
}
func EchoCommand(args []string) {
	fmt.Println(strings.Join(args[1:], " "))
}
func KillCommand(args []string) {
	if len(args) == 1 {
		fmt.Print("usage: kill <process>\n")
	}
	pid, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Print("usage: kill <process>\n")
	}
	p, err := os.FindProcess(pid)
	if err != nil {
		fmt.Printf("kill: %s failed: no such process\n", args[1])
	}
	err = p.Kill()
	if err != nil {
		fmt.Printf("kill: %s failed: %s\n", args[1], err.Error())
	}
}
