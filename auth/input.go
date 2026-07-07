package auth

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadString(input string) string{
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(input)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func WaitForKey() {
	fmt.Print("\nTekan Enter untuk kembali ke menu...")
	bufio.NewReader(os.Stdin).ReadString('\n')
}