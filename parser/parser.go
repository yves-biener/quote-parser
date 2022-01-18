package parse

import (
	"database/sql"
	"log"
	"time"
)

type Result struct {
	Id int
}

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
	ISBN        sql.NullString
	ReleaseDate time.Time
}

type Quote struct {
	Id     int
	BookId int
	Page   int
	Quote  string
}

// this interface should be implemented by the component that makes the REST api
// calls. The Commit should set the Id value of the element!
type Entity interface {
	Commit() error
}

type ParseResult struct {
	entities []Entity
}

// Commit all ParseResults using the interface method 'Commit'
func (result ParseResult) Commit() {
	for _, entity := range result.entities {
		// if a dependend entity fails it will cause a deadlock
		go func(e Entity) {
			if err := e.Commit(); err != nil {
				log.Fatal(err)
			}
		}(entity)
	}
}

type Parser interface {
	Parse() ParseResult
}
