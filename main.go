package main

import (
	"fmt"
	"os"
	"os/exec"
)


func usageExplane() error {
    fmt.Println("- Required min 2 argumets after dcshort")
    fmt.Println("Usage:  dcshort <run|build> <image-name> [<container-name>] [<port>]")
    fmt.Println("build example:  dcshort build image:namev1 # docker build . -t image:namev1 -f Dockerfile")
    fmt.Println("run example:  dcshort run image:namev1 test-con 9090:9090 # docker run --rm -d --name test-con -p 9090 image:namev1")
    fmt.Println("command:")
    fmt.Println("run:   docker run --rm -d --name <container-name> -p <portforward> <image-name>")
    fmt.Println("build:  docker build . -t <image-name> -f Dockerfile")
	return nil
}



func Run(portArgs, containerName, imageName string ) {
    var cmd *exec.Cmd

    if portArgs != "" && containerName != "" {
        cmd = exec.Command("docker", "run", "--rm", "-d", "--name", containerName, "-p", portArgs, imageName)
    } else if containerName != "" {
        cmd = exec.Command("docker", "run", "--rm", "-d", "--name", containerName, imageName)
    }


    err := cmd.Run()
    if err != nil {
        fmt.Printf("Error executing command: %s\n", err)
        return
    }
}

func Build(imageName *string) {
    cmd := exec.Command("docker", "build", ".", "-t", *imageName, "-f", "Dockerfile")

    err := cmd.Run()
    if err != nil {
        fmt.Printf("Error executing command: %s\n", err)
        return
    }
}


func main() {
    if len(os.Args) == 1 {
        usageExplane()
        return
    }

    command := os.Args[1]

	if command == "help" {
		usageExplane()
		return
	}

	if len(os.Args) < 2 {
        usageExplane()
        return
    }

    imageName := os.Args[2] 
	var containerName string
    var portArgs string               

    if len(os.Args) > 4 {
        portArgs = os.Args[4]
    }

	if len(os.Args) > 3 {
        containerName = os.Args[3]
    }


    switch command {
    case "run":
        Run(portArgs, containerName, imageName)
    case "build":
        Build(&imageName)
    default:
        fmt.Println("Unknown command")
        usageExplane()
        return
    }
}
