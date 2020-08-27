package client

import (
	"github.com/yigitsadic/gocandela/models"
	"github.com/yigitsadic/gocandela/shared"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"time"
)

type HTTPClient struct {
	Response string
	Re       *regexp.Regexp
}

func NewHTTPClient() *HTTPClient {
	return &HTTPClient{Re: regexp.MustCompile(shared.LineMatcher)}
}

func (h *HTTPClient) ParseLines() ([]models.Earthquake, error) {
	sm := h.Re.FindAllStringSubmatch(h.Response, -1)

	var earthquakes []models.Earthquake

	for _, element := range sm {
		earthquakes = append(earthquakes, *models.NewEarthquake(element))
	}

	return earthquakes, nil
}

func (h *HTTPClient) Fetch() {
	c := http.Client{
		Timeout: 10 * time.Second,
	}

	res, err := c.Get(shared.BaseURL)
	if err != nil {
		log.Fatalf("Unable to fetch from Kandilli")
	}

	defer res.Body.Close()

	read, _ := ioutil.ReadAll(res.Body)
	h.Response = string(read)
}
