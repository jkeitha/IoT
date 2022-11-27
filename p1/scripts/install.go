package main

import (
	"fmt"
	"os"
	"os/exec"
)

//func strerror(str string, err error) {
//	panic(fmt.Sprintf("%s%s\n", str, err))
//}

type DataStore struct {
	nodeType         string
	sshPublicKeyPath string
	nodeConfigPath   string
}

func NewDataStore(args []string) *DataStore {
	argsLen := len(args)

	if argsLen != 4 {
		panic(fmt.Sprintf("Wrong number of args, number of args is %d!n", argsLen))
	}
	if args[3] == "" {
		panic("Node config path can't be empty")
	}
	if args[2] == "" {
		panic("SSH public key path can't be empty")
	}
	if args[1] == "" {
		panic("Node type can't be empty")
	}
	return &DataStore{nodeType: args[1], sshPublicKeyPath: args[2], nodeConfigPath: args[3]}
}

func main() {
	fmt.Println("Install script running!")
	dataStore := NewDataStore(os.Args)
	addPublicKey(dataStore.sshPublicKeyPath)
}
