package auth

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func ClearScreen(){
	switch runtime.GOOS {
		case "windows" :
			cmd := exec.Command("cmd", "/c", "cls")
			cmd.Stdout = os.Stdout
			cmd.Run()
		default :
			fmt.Print("\033[H\033[2J\033[3J")
	}
}

func ReadString(input string) string{
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(input)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func WaitForKey(mess string) {
	fmt.Printf("\n%s", mess)
	bufio.NewReader(os.Stdin).ReadString('\n')
}

