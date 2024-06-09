package streaming_test

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"robocar-webserver/src/packages/appConfig"
	"robocar-webserver/src/packages/streaming"
	"testing"
)

const port = 10000

// TODO: add auto handling and source stream simulation
func TestStreamDistribution(t *testing.T) {
	err := os.Setenv("ROBOCAR_STREAM_URL", "http://192.168.1.71:81")
	appConfig.Init()

	if err != nil {
		t.Errorf("Failed to set environment variable: %v", err)
	}

	streamHandler := streaming.MakeMultiplexedStreamHandler()

	http.HandleFunc("/stream", streamHandler)

	log.Printf("Server started on http://localhost:%d/stream\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
