package main

import (
//	"fmt"
	"os"
	"os/exec"
	"log"
)


func isInstalled(packageName string) bool {
	dpkgL := exec.Command("dpkg", "-S", packageName)
	err := dpkgL.Run()
	if err != nil {
		log.Printf("%s not installed\n\n", packageName)
		return false
	}
	log.Printf("%s already installed\n\n", packageName)
	return true
}

func aptInstall(tools []string) {
	for _, tool := range tools {
		log.Printf("Start %s installation\n", tool)
		if isInstalled(tool) {
			continue
		}
		install := exec.Command("apt", "install", "-y", tool)
		install.Stdout = os.Stdout
		install.Stderr = os.Stderr
		err := install.Run()
		if err != nil {
			log.Fatalf("%s installation failed with error: %s\n\n", tool, err)
		}
		log.Printf("%s succesfully installed\n\n", tool)
	}
}

func installK3d() {
	log.Println("Start k3d installation!")
	installed := exec.Command("k3d", "version")
	err := installed.Run()
	if err == nil {
		log.Println("k3d already installed\n")
		return
	}
	download := exec.Command("wget", "-q", "-O", "-", "https://raw.githubusercontent.com/k3d-io/k3d/main/install.sh")
	execute := exec.Command("bash")
	execute.Stdin, _ = download.StdoutPipe()
	err = download.Run()
	if err != nil {
		log.Fatalln("Download failed with error: \n", err )
	}
	err = execute.Run()
	if err != nil {
		log.Fatalln("Execute failed with error: \n", err )
	}
	log.Println("k3d succesfully installed!\n")
}

func main() {
	log.Println("Start setup part 3 of Inception Of Things!\n")
	aptInstall([]string{"docker.io", "kubectl", "wget"})
	installK3d()
	log.Println("Setup of third part of Inception Of Things Complete!")
}
