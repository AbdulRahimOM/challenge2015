package tracer

import (
	"context"
	"sync"
	"test/internal/config"
	"test/internal/data"

	"github.com/gofiber/fiber/v2/log"
)

var (
	fetchPersonWorkersCount = config.PersonDataFetchWorkersCount
	fetchMovieWorkersCount  = config.MovieDataFetchWorkersCount
)

type personRequest struct {
	personURL string
	replyChan chan *data.Person
	ctx       context.Context
	wg        *sync.WaitGroup
}

type movieRequest struct {
	movieURL  string
	replyChan chan *data.Movie
	ctx       context.Context
	wg        *sync.WaitGroup
}

var externalFetcher = struct {
	PersonRequestChan chan personRequest
	MovieRequestChan  chan movieRequest
}{
	PersonRequestChan: make(chan personRequest, 3*fetchPersonWorkersCount),
	MovieRequestChan:  make(chan movieRequest, 3*fetchMovieWorkersCount),
}

// initiate workers
func init() {
	for range fetchPersonWorkersCount {
		go fetchPersonWorker(externalFetcher.PersonRequestChan)
	}
	for range fetchMovieWorkersCount {
		go fetchMovieWorker(externalFetcher.MovieRequestChan)
	}
}

// fetchPersonWorkers
func fetchPersonWorker(personRequestChan chan personRequest) {
	for req := range personRequestChan {
		fetchAndCachePerson(req.personURL, req.wg, req.replyChan, req.ctx)
	}
}

func fetchAndCachePerson(personURL string, wg *sync.WaitGroup, replyChan chan *data.Person, ctx context.Context) {
	defer wg.Done()

	// Again recheck if the person is cached (while the URL was in queue)
	if person := data.CachedData.GetCachedPerson(personURL); person != nil {
		select {
		case replyChan <- person:
			return
		case <-ctx.Done():
			return
		}
	}
	person, err := data.FetchPersonDataFromExternalAPI(personURL)
	if err != nil {
		log.Debugf("failed to fetch person data: %v\n", err)
		return
	}

	data.CachedData.CachePerson(personURL, person)

	select {
	case replyChan <- person:
		return
	case <-ctx.Done():
		return
	}
}

// fetchMovieWorkers
func fetchMovieWorker(requestChan chan movieRequest) {
	for req := range requestChan {
		fetchAndCacheMovie(req.movieURL, req.wg, req.replyChan, req.ctx)
	}
}

func fetchAndCacheMovie(movieURL string, wg *sync.WaitGroup, replyChan chan *data.Movie, ctx context.Context) {
	defer wg.Done()

		//Again recheck if the movie is cached (while the URL was in queue)
		if movie := data.CachedData.GetCachedMovie(movieURL); movie != nil {
			select {
			case replyChan <- movie:
				return
			case <-ctx.Done():
				return
			}
		}

		movie, err := data.FetchMovieDataFromExternalAPI(movieURL)
		if err != nil {
			log.Debugf("failed to fetch movie data: %v\n", err)
			return
		}

		data.CachedData.CacheMovie(movieURL, movie)

		select {
		case replyChan <- movie:
			return
		case <-ctx.Done():
			return
		}
	}
