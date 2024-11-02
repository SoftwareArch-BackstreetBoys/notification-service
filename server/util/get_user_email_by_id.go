package util

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/SoftwareArch-BackstreetBoys/notification-service/configs"
)

func GetUserEmailById(userId string) string {
	return "jaijai211075@gmail.com"

	// Create a new GET request
	req, err := http.NewRequest("GET", configs.UserServiceURI()+"/users/"+userId, nil)
	if err != nil {
		fmt.Println(err)
	}

	// Create an HTTP client
	client := &http.Client{}

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(body))

	return ""
}
