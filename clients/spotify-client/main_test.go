package main

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/magiconair/properties/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"testing"
)

func TestHandlerWithDockerImageBasedMockApi(t *testing.T) {

	t.Run("Should correctly load albums", func(t *testing.T) {
		ctx := context.Background()
		err := runSpotifyMock(ctx)

		response, err := handler(events.APIGatewayProxyRequest{
			HTTPMethod: "GET",
		})
		if err != nil {
			t.Fatal("Everything should be ok")
		}
		assert.Equal(t, response.StatusCode, 200)
	})

	t.Run("Should fail due to throttling", func(t *testing.T) {
		ctx := context.Background()
		err := runSpotifyMock(ctx)

		response, err := handler(events.APIGatewayProxyRequest{
			HTTPMethod: "GET",
			Headers: map[string]string{
				"Scenario": "Throttle",
			},
		})
		if err != nil {
			t.Fatal("Everything should be ok")
		}
		assert.Equal(t, response.StatusCode, 429)
	})
}

func TestHandlerWithDockerfileBasedMockApi(t *testing.T) {

	t.Run("Should correctly load albums", func(t *testing.T) {
		ctx := context.Background()
		err := runSpotifyMockFromDefinition(ctx)

		response, err := handler(events.APIGatewayProxyRequest{
			HTTPMethod: "GET",
		})
		if err != nil {
			t.Fatal("Everything should be ok")
		}
		assert.Equal(t, response.StatusCode, 200)
	})

	t.Run("Should fail due to throttling", func(t *testing.T) {
		ctx := context.Background()
		err := runSpotifyMockFromDefinition(ctx)

		response, err := handler(events.APIGatewayProxyRequest{
			HTTPMethod: "GET",
			Headers: map[string]string{
				"Scenario": "Throttle",
			},
		})
		if err != nil {
			t.Fatal("Everything should be ok")
		}
		assert.Equal(t, response.StatusCode, 429)
	})
}

func runSpotifyMock(ctx context.Context) error {
	req := testcontainers.ContainerRequest{
		Image:        "spotify-mock",
		ExposedPorts: []string{"3000/tcp"},
		WaitingFor:   wait.ForListeningPort("3000/tcp"),
	}
	spotifyMock, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})

	endpoint, err := spotifyMock.Endpoint(ctx, "")
	if err != nil {
		logs, err := spotifyMock.Logs(ctx)
		if err == nil {
			buf := new(bytes.Buffer)
			buf.ReadFrom(logs)
			fmt.Printf(buf.String())
		}
		panic(err)
	}
	DefaultHTTPGetAddress = fmt.Sprintf("http://%s", endpoint)
	return err
}

func runSpotifyMockFromDefinition(ctx context.Context) error {
	req := testcontainers.ContainerRequest{
		FromDockerfile: testcontainers.FromDockerfile{
			Context:    "../../definition",
			Dockerfile: "Dockerfile",
		},
		ExposedPorts: []string{"3000/tcp"},
		WaitingFor:   wait.ForListeningPort("3000/tcp"),
	}
	spotifyMock, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})

	endpoint, err := spotifyMock.Endpoint(ctx, "")
	if err != nil {
		logs, err := spotifyMock.Logs(ctx)
		if err == nil {
			buf := new(bytes.Buffer)
			buf.ReadFrom(logs)
			fmt.Printf(buf.String())
		}
		panic(err)
	}
	DefaultHTTPGetAddress = fmt.Sprintf("http://%s", endpoint)
	return err
}
