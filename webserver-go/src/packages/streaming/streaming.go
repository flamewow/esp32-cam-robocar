package streaming

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"math/rand"
	"mime"
	"mime/multipart"
	"net/http"
	"robocar-webserver/src/packages/appConfig"
	"strings"
)

var _config = appConfig.Load()

func consumeMP(ch chan *bytes.Buffer) error {
	resp, err := http.Get(_config.RobocarStreamUrl)

	if err != nil {
		fmt.Println("Error fetching the URL:", err)
		return err
	}

	defer resp.Body.Close()

	contentType := resp.Header.Get("Content-Type")
	mediatype, p, err := mime.ParseMediaType(contentType)

	if !strings.HasPrefix(mediatype, "multipart/") {
		fmt.Println("Invalid content type:", contentType)
		return err
	}

	boundary := p["boundary"]
	reader := multipart.NewReader(resp.Body, boundary)

	for {
		part, err := reader.NextPart()

		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error reading part:", err)
			return err
		}

		defer part.Close()

		buffer := bytes.NewBuffer(nil)
		if _, err := io.Copy(buffer, part); err != nil {
			return err
		}

		ch <- buffer
	}

	return nil
}

func multiplex(srcCh chan *bytes.Buffer, fanOut map[int]chan *bytes.Buffer) {
	for {
		select {
		case buffer := <-srcCh:
			for _, ch := range fanOut {
				ch <- buffer
			}
		}
	}
}

func distributeStream(fanOut map[int]chan *bytes.Buffer, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "multipart/x-mixed-replace; boundary=frame")

	channelID := rand.Int()
	transmissionChannel := make(chan *bytes.Buffer)
	fanOut[channelID] = transmissionChannel

	log.Printf("Channel %d opened\n", channelID)

	defer func() {
		delete(fanOut, channelID)
		close(transmissionChannel)

		log.Printf("Channel %d closed\n", channelID)
	}()

	done := false

	for done == false {
		select {
		case buffer := <-transmissionChannel:
			func() {
				fmt.Fprintf(w, "--frame\r\nContent-Type: image/jpeg\r\n\r\n%s\r\n", buffer)
				flusher, ok := w.(http.Flusher)
				if ok {
					flusher.Flush()
				}
			}()
		case <-r.Context().Done():
			done = true
		}
	}
}

// MakeMultiplexedStreamHandler creates a handler that streams the multipart response
func MakeMultiplexedStreamHandler() http.HandlerFunc {
	mainChannel := make(chan *bytes.Buffer)
	fanOut := make(map[int]chan *bytes.Buffer)

	go func() {
		err := consumeMP(mainChannel)
		if err != nil {
			fmt.Println("Error consuming multipart:", err)
		}
	}()

	go multiplex(mainChannel, fanOut)

	fmt.Println("Multiplexed stream handler created")
	return func(w http.ResponseWriter, r *http.Request) {
		distributeStream(fanOut, w, r)
	}
}
