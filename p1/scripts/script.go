package main

import (
	"fmt"
	"os"
	"os/exec"
)

func addPublicKey(publicKey string) {
	exec.Command("cat", publicKey, ">>", "~/.ssh/authorized_keys")
}

func main() {
	fmt.Println("Hello vagrant!")
	addPublicKey(os.Args[1])
}
