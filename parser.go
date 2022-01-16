package parse

import (
	"log"
	"time"
)

type Author struct {
	Id   int
	Name string
}

type Topic struct {
	Id    int
	Topic string
}

type Language struct {
	Id       int
	Language string
}

type Book struct {
	Id          int
	AuthorId    int
	TopicId     int
	LanguageId  int
	Title       string
	ISBN        string
	ReleaseDate time.Time
}

type Quote struct {
	Id     int
	BookId int
	Page   int
	Quote  string
}

// this interface should be implemented by the component that makes the REST api
// calls
type Entity interface {
	Commit() error
}

type ParseResult struct {
	authors   []Entity
	topics    []Entity
	languages []Entity
	books     []Entity
	quotes    []Entity
}

// Commit all ParseResults using the interface method 'Commit'
func (result ParseResult) Commit() {
	for _, author := range result.authors {
		if err := author.Commit(); err != nil {
			log.Fatal(err)
		}
	}
	for _, topic := range result.topics {
		if err := topic.Commit(); err != nil {
			log.Fatal(err)
		}
	}
	for _, language := range result.languages {
		if err := language.Commit(); err != nil {
			log.Fatal(err)
		}
	}
	for _, book := range result.books {
		if err := book.Commit(); err != nil {
			log.Fatal(err)
		}
	}
	for _, quote := range result.quotes {
		if err := quote.Commit(); err != nil {
			log.Fatal(err)
		}
	}
}

type Parser interface {
	Parse() ParseResult
}
