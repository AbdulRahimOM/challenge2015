package data

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func FetchMovieDataFromExternalAPI(movieURL string) (*Movie, error) {
	url := fmt.Sprintf("http://data.moviebuff.com/%s", movieURL)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch data: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-200 response: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	var movie Movie
	err = json.Unmarshal(body, &movie)
	if err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %v", err)
	}

	return &movie, nil
}

func FetchPersonDataFromExternalAPI(personURL string) (*Person, error) {
	url := fmt.Sprintf("http://data.moviebuff.com/%s", personURL)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch data: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-200 response: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	var person Person
	err = json.Unmarshal(body, &person)
	if err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %v", err)
	}

	return &person, nil
}
