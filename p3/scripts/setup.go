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
		log.Fatalf("Download failed with error: %s\n\n", err)
	}
	err = execute.Run()
	if err != nil {
		log.Fatalf("Execute failed with error: %s\n\n", err)
	}
	log.Println("k3d succesfully installed!\n")
}

func createCluster(clusterName string) {
	log.Println("Start k3s cluster creation!")
	create := exec.Command("k3d", "cluster", "create", clusterName)
	err := create.Run()
	if err != nil {
		log.Fatalf("k3d cluster creation failed with error: %s\n\n", err)
	}
	log.Println("k3s cluster succesfully installed!\n")
}

func createNamespaces(namespaces []string) {
	log.Println("Start k3s namespaces creation!")
	for _, ns := range namespaces {
		create := exec.Command("kubectl", "create", "namespace", ns)
		err := create.Run()
		if err != nil {
			log.Fatalf("k3s namespace creation failed with error: %s\n\n", err)
		}
		log.Printf("k3s namespace %s created!\n", ns)
	}
	log.Println("All k3s namespaces succesfully created!")
}

func main() {
	log.Println("Start setup part 3 of Inception Of Things!\n")
	aptInstall([]string{"docker.io", "kubectl", "wget"})
	installK3d()
	createCluster("swquinc")
	createNamespaces([]string{"dev", "argo"})
	deployArgo()
	log.Println("Setup of third part of Inception Of Things Complete!")
}
