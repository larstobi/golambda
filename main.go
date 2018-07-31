package main

// https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-use-lambda-authorizer.html

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	// "os"
	// "strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var (
	ErrorBackend = errors.New("Something went wrong")
)

type Request struct {
	ID int `json:"id"`
}

type DeployArtifactResponse struct {
	Artifacts []Artifact `json:"results"`
}
// {"groupId":"com.myorganization.app","artifactId":"myapp","version":"1.0","packaging":"jar"}✔ ~/git/java-hello-world-maven [master|✚ 1]

type Artifact struct {
	GroupId       string `json:"groupId"`
	ArtifactId    string `json:"artifactId"`
	Version       string `json:"version"`
	Packaging     string `json:"packaging"`
}

func Handler(request Request) (events.APIGatewayProxyResponse, error) {
	log.Print("*** HOOOOOO")
	url := fmt.Sprintf("https://api.themoviedb.org/3/discover/movie?api_key=")

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return events.APIGatewayProxyResponse{}, ErrorBackend
	}

	if request.ID > 0 {
		q := req.URL.Query()
		// q.Add("with_genres", strconv.Itoa(request.ID))
		req.URL.RawQuery = q.Encode()
	}

	resp, err := client.Do(req)
	if err != nil {
		return events.APIGatewayProxyResponse{}, ErrorBackend
	}
	defer resp.Body.Close()

	var data DeployArtifactResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return events.APIGatewayProxyResponse{}, ErrorBackend
	}

	log.Print("*** HEEEEEY")
	body, err := json.Marshal(data)
	return events.APIGatewayProxyResponse{
		Body:       string(body),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
