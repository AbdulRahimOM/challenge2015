package data

import (
	"fmt"
	"net/http"
	_ "test/internal/status" //to use GetStatistics function
)

func CheckPersonExistence(personURL string) (bool, error) {
	if CachedData.GetCachedPerson(personURL) != nil {
		return true, nil
	}

	// If not in cache, check if it exists in external API
	url := fmt.Sprintf("http://data.moviebuff.com/%s", personURL)

	resp, err := http.Head(url) // Use HEAD instead of GET to check existence
	if err != nil {
		return false, fmt.Errorf("failed to fetch data: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		return true, nil
	} else {
		return false, nil
	}
}
