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
		fmt.Println(*step.Name)
		method := step.HTTP.Method
		url := step.HTTP.URL
		parsedUrl := testly.ReplacePlaceholders(url, envs)
		queryParams := step.HTTP.Params

		if queryParams != nil {
			for key, value := range queryParams {
				parsedUrl += fmt.Sprintf("%s=%s&", key, value)
			}
		}

		switch method {
		case "GET":
			getRequest, err := testly.GetRequest(parsedUrl, step.HTTP.Headers, step.HTTP.Auth)
			if err != nil {
				panic(err)
			}

			requestBody := getRequest.Body
			fmt.Println("GET", parsedUrl, requestBody)
		case "POST":
			postRequest, err := testly.PostRequest(parsedUrl, step.HTTP.Headers, step.HTTP.Auth, step.HTTP.Body)
			if err != nil {
				fmt.Println(err)
				panic(err)
			}
			fmt.Println("POST", parsedUrl, postRequest)
		case "PUT":
			putRequest, err := testly.PutRequest(parsedUrl, step.HTTP.Headers, step.HTTP.Auth, step.HTTP.Body)
			if err != nil {
				fmt.Println(err)
				panic(err)
			}
			fmt.Println("PUT", parsedUrl, putRequest)
		case "DELETE":
			deleteRequest, err := testly.DeleteRequest(parsedUrl, step.HTTP.Headers, step.HTTP.Auth)
			if err != nil {
				fmt.Println(err)
				panic(err)
			}
			fmt.Println("DELETE", parsedUrl, deleteRequest)
		default:
			fmt.Println("Method not supported")
		}

	}
}
