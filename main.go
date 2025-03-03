package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	testly "github.com/rukshn/testly/testly_api"
)

func main() {
	var testFile string
	var jsonTests testly.Test

	flag.StringVar(&testFile, "file", "test.json", "File to read")

	fileContent, err := os.ReadFile(testFile)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	json.Unmarshal(fileContent, &jsonTests)
	steps := jsonTests.Steps
	envs := jsonTests.Env
	for _, step := range steps {
		method := step.HTTP.Method
		url := step.HTTP.URL
		parsedUrl := testly.ReplacePlaceholders(url, envs)

		switch method {
		case "GET":
			getRequest, err := testly.GetRequest(parsedUrl, step.HTTP.Headers, step.HTTP.Auth)
			if err != nil {
				fmt.Println(err)
				panic(err)
			}
			fmt.Println("GET", parsedUrl, getRequest)
		case "POST":
			fmt.Println("POST", parsedUrl)
		case "PUT":
			fmt.Println("PUT", parsedUrl)
		case "DELETE":
			fmt.Println("DELETE", parsedUrl)
		default:
			fmt.Println("Method not supported")
		}

	}
}
