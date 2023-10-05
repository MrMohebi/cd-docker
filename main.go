package main

import (
	"cd-docker/common"
	"cd-docker/configs"
	"cd-docker/router"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

// nodemon --exec go run main.go --signal SIGTERM

func main() {
	configs.Setup()

	//if it's a terminal call then never web server runs
	if terminalCall() {
		os.Exit(0)
	}

	server := gin.Default()

	router.Routs(server)

	err := server.Run("0.0.0.0:8005")
	common.IsErr(err, true, "Err in starting server")
}

func terminalCall() bool {
	var isUsed = true

	var name string
	flag.StringVar(&name, "name", "", "[required] name of service -- will be used to determine service by API call")
	flag.StringVar(&name, "n", "", "[required] name of service -- will be used to determine service by API call")

	var token string
	flag.StringVar(&token, "token", "", "[required] token for API authentication")
	flag.StringVar(&token, "t", "", "[required] token for API authentication")

	var dockerService string
	flag.StringVar(&dockerService, "service-name", "", "[required] docker service name ")
	flag.StringVar(&dockerService, "s", "", "[required] docker service name")

	var deleteServiceName string
	flag.StringVar(&deleteServiceName, "delete", "", "[required] service name to be deleted")
	flag.StringVar(&deleteServiceName, "d", "", "[required] service name to be deleted")

	var help bool
	flag.BoolVar(&help, "help", false, "[required] print this usage")
	flag.BoolVar(&help, "h", false, "[required] print this usage")

	flag.Usage = func() {
		const help = `Usage of cd-docker:
	-n, --name               [required] name of service -- will be used to determine service by API call
	-t, --token              [required] token for API authentication 
	-s, --service-name       [required] docker service name
	-d, --delete             service name to be deleted
	-h, --help               print this usage
						`
		fmt.Println(help)
	}

	flag.Parse()

	if isFlagPassed("help") || isFlagPassed("h") {
		flag.Usage()
		os.Exit(0)
	} else if len(deleteServiceName) != 0 {
		err := deleteService(deleteServiceName)
		common.IsErr(err, true)
		os.Exit(0)
	} else if len(name) > 0 && len(dockerService) > 0 && len(token) > 0 {
		err := createService(name, dockerService, token)
		common.IsErr(err, true)
		os.Exit(0)
	} else {
		isUsed = false
	}

	println("Please specify required parameters")
	flag.Usage()
	return isUsed
}

func deleteService(name string) error {
	if !configs.IniData.HasSection(name) {
		fmt.Printf("service with given name [%q] dosen't exists\n", name)
		os.Exit(0)
	}
	configs.IniData.DeleteSection(name)
	err := configs.IniSave()
	fmt.Printf("service with given name [%q] has been deleted\n", name)
	return err
}

func createService(name string, serviceName string, token string) error {
	if configs.IniData.HasSection(name) {
		fmt.Printf("service with given name [%q] already exists\n", name)
		os.Exit(0)
	}

	configs.IniData.Section(name).Key("serviceName").SetValue(serviceName)
	configs.IniData.Section(name).Key("token").SetValue(token)

	err := configs.IniSave()
	println("service created and saved to ./config.ini")
	return err
}

func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}
