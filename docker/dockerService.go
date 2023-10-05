package docker

import (
	"cd-docker/common"
	"fmt"
	"os/exec"
	"strings"
)

func UpdateDockerService(dockerServiceName string, image string) error {
	fmt.Printf("pulling latest image [%s]...\n", image)
	pullImageOut, err := exec.Command("docker", "pull", image).Output()
	common.IsErr(err, false)
	fmt.Println(string(pullImageOut))

	fmt.Printf("update service [%s] with latest image...\n", dockerServiceName)
	updateOut, err := exec.Command(
		"docker",
		"service",
		"update",
		"--force",
		"--with-registry-auth",
		dockerServiceName,
		"--image",
		image,
	).Output()

	fmt.Printf("service [%s] updated with latest image!\n", dockerServiceName)

	fmt.Println(string(updateOut))

	return err
}

func UpdateDockerServiceLatest(dockerServiceName string) error {
	image, err := GetServiceImageNameWithoutTag(dockerServiceName)
	common.IsErr(err, false)

	fmt.Printf("pulling latest image [%s]...\n", image)
	pullImageOut, err := exec.Command("docker", "pull", image+":latest").Output()
	common.IsErr(err, false)
	fmt.Println(string(pullImageOut))

	fmt.Printf("update service [%s] with latest image...\n", dockerServiceName)
	updateOut, err := exec.Command(
		"docker",
		"service",
		"update",
		"--force",
		"--with-registry-auth",
		dockerServiceName,
		"--image",
		image+":latest",
	).Output()

	fmt.Printf("service [%s] updated with latest image!\n", dockerServiceName)

	fmt.Println(string(updateOut))

	return err
}

func GetServiceImageNameWithoutTag(serviceName string) (string, error) {
	cmd := exec.Command("docker", "service", "inspect", "--format={{ index .Spec.TaskTemplate.ContainerSpec.Image }}", serviceName)
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	imageName := string(output)
	imageName = imageName[:len(imageName)-1] // remove newline character at the end
	imageName = strings.Split(imageName, "@")[0]
	imageName = strings.Split(imageName, ":")[0]

	return imageName, nil
}
