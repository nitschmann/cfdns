package main

import (
	"fmt"
	"log"
	"os/user"

	"github.com/nitschmann/cfdns/pkg/checkip"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	// Current User
	fmt.Println("Hi " + user.HomeDir + " (id: " + user.Uid + ")")

	ckeckipClient := checkip.New()
	ip, err := ckeckipClient.GetPublicIpV4()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(ip)
}
