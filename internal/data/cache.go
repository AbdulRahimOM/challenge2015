package data

import (
	"sync"
)

type cache struct {
	PersonMutex sync.RWMutex
	MovieMutex  sync.RWMutex
	Persons     map[string]*Person
	Movies      map[string]*Movie
}

var CachedData cache

func init() {
	CachedData.Persons = make(map[string]*Person)
	CachedData.Movies = make(map[string]*Movie)
}

func (c *cache) GetCachedPerson(personURL string) *Person {
	c.PersonMutex.RLock()
	defer c.PersonMutex.RUnlock()
	return c.Persons[personURL]
}

func (c *cache) GetCachedMovie(movieURL string) *Movie {
	c.MovieMutex.RLock()
	defer c.MovieMutex.RUnlock()
	return c.Movies[movieURL]
}

func (c *cache) CachePerson(personURL string, person *Person) {
	if person == nil {
		return
	}
	c.PersonMutex.Lock()
	defer c.PersonMutex.Unlock()
	if _, ok := c.Persons[personURL]; ok {
		c.Persons[personURL] = person
	}
}

func (c *cache) CacheMovie(movieURL string, movie *Movie) {
	if movie == nil {
		return
	}
	c.MovieMutex.Lock()
	defer c.MovieMutex.Unlock()
	if _, ok := c.Movies[movieURL]; ok {
		c.Movies[movieURL] = movie
	}
}
