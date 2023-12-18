package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/raghavsood/spn2"
)

func main() {
	accessKey := os.Getenv("SPN2_ACCESS_KEY")
	secretKey := os.Getenv("SPN2_SECRET_KEY")

	if accessKey == "" || secretKey == "" {
		fmt.Println("Access key and secret key must be set as environment variables.")
		os.Exit(1)
	}

	client := spn2.NewClient(accessKey, secretKey)

	args := os.Args
	if len(args) < 2 {
		fmt.Println("Usage: spn2 <command> [arguments]")
		os.Exit(1)
	}

	command := args[1]
	switch command {
	case "save":
		if len(args) != 3 {
			fmt.Println("Usage: spn2 save <url>")
			os.Exit(1)
		}
		url := args[2]
		response, err := client.SubmitURL(url)
		handleResponse(response, err)

	case "status":
		if len(args) != 3 {
			fmt.Println("Usage: spn2 status <job_id>")
			os.Exit(1)
		}
		jobID := args[2]
		response, err := client.GetStatus(jobID)
		handleResponse(response, err)

	case "health":
		response, err := client.GetSystemStatus()
		handleResponse(response, err)

	case "user":
		response, err := client.GetUserStatus()
		handleResponse(response, err)

	default:
		fmt.Println("Invalid command. Available commands: save, status, health, user")
		os.Exit(1)
	}
}

func handleResponse(response interface{}, err error) {
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}

	jsonData, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		fmt.Printf("Error marshaling JSON: %s\n", err)
		os.Exit(1)
	}

	fmt.Println(string(jsonData))
}
