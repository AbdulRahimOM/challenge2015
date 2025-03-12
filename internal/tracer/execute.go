package tracer

import (
	"context"
	"fmt"
	"sync"
	"test/internal/data"
)

func FindSeperation(p1URL string, targetPerson string) (int, error) {

	if p1URL == targetPerson {
		return 0, nil
	}

	//check if target person exists to avoid infinite search
	exists, err := data.CheckPersonExistence(targetPerson)
	if err != nil {
		return 0, fmt.Errorf("error checking target person existence: %v", err)
	}
	if !exists {
		return 0, fmt.Errorf("target person not found")
	}

	var (
		personURLQueue = []string{p1URL}
		visitedPersons = make(map[string]bool)
		visitedMovies  = make(map[string]bool)
	)

	for seperation := 2; len(personURLQueue) > 0; seperation++ {
		found, newPersonURLQueue := findTargetOrNextPersonList(personURLQueue, targetPerson, visitedPersons, visitedMovies)
		if found {
			return seperation, nil
		}
		personURLQueue = newPersonURLQueue
	}

	return -1, fmt.Errorf("seperation not found")

}

func findTargetOrNextPersonList(personURLQueue []string, targetPerson string, visitedPersons map[string]bool, visitedMovies map[string]bool) (bool, []string) {
	var (
		personChan   = make(chan *data.Person, 10)
		movieChan    = make(chan *data.Movie, 10)
		movieUrlChan = make(chan string, 100)
		ctx          = context.Background()
	)

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	go startFetchingPersons(personURLQueue, personChan, ctx)
	for _, personURL := range personURLQueue {
		visitedPersons[personURL] = true
	}

	//receive person data and send movie urls
	go func() {
		defer close(movieUrlChan)
		for {
			select {
			case personData, ok := <-personChan:
				if !ok {
					return
				}
				for _, movieRole := range personData.MovieRoles {
					if _, isVisited := visitedMovies[movieRole.URL]; isVisited {
						continue
					} else {
						visitedMovies[movieRole.URL] = true

						select {
						case movieUrlChan <- movieRole.URL:
						case <-ctx.Done():
							return
						}
					}
				}
			case <-ctx.Done():
				return
			}
		}
	}()

	//call startMovieFetching to fetch movie data and send person urls of cast and crew
	go startMovieFetching(movieUrlChan, movieChan, ctx)

	newPersonURLQueue := []string{}

	//add unvisited persons to new queue
	for movieData := range movieChan {
		for _, cast := range movieData.Cast {
			if cast.URL == targetPerson {
				return true, nil
			}
			if _, isVisited := visitedPersons[cast.URL]; isVisited {
				continue
			}
			newPersonURLQueue = append(newPersonURLQueue, cast.URL)
		}

		for _, crew := range movieData.Crew {
			if _, isVisited := visitedPersons[crew.URL]; isVisited {
				continue
			}
			if crew.URL == targetPerson {
				return true, nil
			}
			newPersonURLQueue = append(newPersonURLQueue, crew.URL)
		}
	}

	//reset personURLQueue to newPersonURLQueue
	return false, newPersonURLQueue
}

func startFetchingPersons(personURLs []string, personChan chan *data.Person, ctx context.Context) {

	wg := sync.WaitGroup{}
	defer close(personChan)
	defer wg.Wait()
	for _, personURL := range personURLs {
		if person := data.CachedData.GetCachedPerson(personURL); person != nil {
			select {
			case personChan <- person:
				continue
			case <-ctx.Done():
				return
			}
		}

		wg.Add(1)
		select {
		case <-ctx.Done():
			return
		default:
			externalFetcher.PersonRequestChan <- personRequest{
				personURL: personURL,
				replyChan: personChan,
				ctx:       ctx,
				wg:        &wg,
			}
		}
	}
}

// initiateMovieFetcher
func startMovieFetching(movieURLsChan chan string, movieChan chan *data.Movie, ctx context.Context) {

	wg := sync.WaitGroup{}
	defer close(movieChan)
	defer wg.Wait()

	for {
		select {
		case movieURL, ok := <-movieURLsChan:
			if !ok {
				return
			}
			if movie := data.CachedData.GetCachedMovie(movieURL); movie != nil {
				select {
				case movieChan <- movie:
					continue
				case <-ctx.Done():
					return
				}
			}
			wg.Add(1)
			select {
			case <-ctx.Done():
				return
			default:
				externalFetcher.MovieRequestChan <- movieRequest{
					movieURL:  movieURL,
					replyChan: movieChan,
					ctx:       ctx,
					wg:        &wg,
				}
			}
		case <-ctx.Done():
			return
		}
	}
}
