package main


import (
    "fmt"
    "os"
    "os/exec"
)


func usageExplane() {
		fmt.Println("Usage: docker-short <run|build> <image-name> [<container-name>] [<port>]")
		fmt.Println("example: docker-short build image:namev1 # docker build . -t image:namev1 -f Dockerfile")
		fmt.Println("example: docker-short run image:namev1 test-con 9090:9090 # docker run --rm -d --name test-con -p 9090 image:namev1")
		fmt.Println("command:")
		fmt.Println("run:  docker run --rm -d --name <container-name> -p <portforward> <image-name>")
		fmt.Println("build: docker build . -t <image-name> -f Dockerfile")
	return 
}


func main() {
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
        containerName := os.Args[3]
    }

    var cmd *exec.Cmd

    switch command {
    case "run":
        if portArgs != "" && containerName != "" {
            cmd = exec.Command("docker", "run", "--rm", "-d", "--name", containerName, "-p", portArgs, imageName)
        } else if containerName != "" {
            cmd = exec.Command("docker", "run", "--rm", "-d", "--name", containerName, imageName)
        }
    case "build":
        cmd = exec.Command("docker", "build", ".", "-t", imageName, "-f", "Dockerfile")
    default:
        fmt.Println("Unknown command")
        return
    }

    err := cmd.Run()
    if err != nil {
        fmt.Printf("Error executing command: %s\n", err)
        return
    }
}
