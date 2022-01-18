package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	parse "parse/parser"
	"strconv"
	"strings"
	"time"
)

const (
	rootUrl     = "http://localhost:8000/api"
	authorUrl   = rootUrl + "/authors"
	topicUrl    = rootUrl + "/topics"
	languageUrl = rootUrl + "/languages"
	bookUrl     = authorUrl + "/%d/books"
	quotesUrl   = rootUrl + "/quotes"
)

func userGetAuthor() (author parse.Author) {
	response, err := http.Get(authorUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	authorsJson, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	var authors []parse.Author
	err = json.Unmarshal(authorsJson, &authors)
	if err != nil {
		log.Fatal(err)
	}
	for _, author := range authors {
		fmt.Printf("Id: %d\tName: %s\n", author.Id, author.Name)
	}
	fmt.Println("Input Id to select an existing author, or name of new author:")
	var userInput string
	_, err = fmt.Scanln(&userInput)
	if err != nil {
		log.Fatal(err)
	}
	authorId, err := strconv.Atoi(userInput)
	if err != nil {
		if len(userInput) == 0 {
			log.Fatal(err)
		}
		author.Name = userInput
		// create author
		data := url.Values{}
		data.Add("Name", userInput)
		response, err = http.Post(
			authorUrl,
			"application/x-www-form-urlencoded",
			strings.NewReader(data.Encode()))
		if err != nil {
			log.Fatal(err)
		}
		defer response.Body.Close()
		resultJson, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		result := parse.Result{}
		err = json.Unmarshal(resultJson, &result)
		if err != nil {
			log.Fatal(err)
		}
		author.Id = result.Id
	} else {
		for _, a := range authors {
			if a.Id == authorId {
				author = a
				return
			}
		}
		log.Fatalf("No author with Id: %d exists", authorId)
	}
	return
}

func userGetTopic() (topic parse.Topic) {
	response, err := http.Get(topicUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	topicsJson, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	var topics []parse.Topic
	err = json.Unmarshal(topicsJson, &topics)
	if err != nil {
		log.Fatal(err)
	}
	for _, topic := range topics {
		fmt.Printf("Id: %d\tTopic: %s\n", topic.Id, topic.Topic)
	}
	fmt.Println("Input Id to select an existing topic, or new topic:")
	var userInput string
	_, err = fmt.Scanln(&userInput)
	if err != nil {
		log.Fatal(err)
	}
	topicId, err := strconv.Atoi(userInput)
	if err != nil {
		if len(userInput) == 0 {
			log.Fatal(err)
		}
		topic.Topic = userInput
		// create topic
		data := url.Values{}
		data.Add("Topic", userInput)
		response, err = http.Post(
			topicUrl,
			"application/x-www-form-urlencoded",
			strings.NewReader(data.Encode()))
		if err != nil {
			log.Fatal(err)
		}
		defer response.Body.Close()
		resultJson, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		result := parse.Result{}
		err = json.Unmarshal(resultJson, &result)
		if err != nil {
			log.Fatal(err)
		}
		topic.Id = result.Id
	} else {
		for _, a := range topics {
			if a.Id == topicId {
				topic = a
				return
			}
		}
		log.Fatalf("No topic with Id: %d exists", topicId)
	}
	return
}

func userGetLanguage() (language parse.Language) {
	response, err := http.Get(languageUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	languagesJson, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	var languages []parse.Language
	err = json.Unmarshal(languagesJson, &languages)
	if err != nil {
		log.Fatal(err)
	}
	for _, language := range languages {
		fmt.Printf("Id: %d\tLanguage: %s\n", language.Id, language.Language)
	}
	fmt.Println("Input Id to select an existing language, or new language:")
	var userInput string
	_, err = fmt.Scanln(&userInput)
	if err != nil {
		log.Fatal(err)
	}
	languageId, err := strconv.Atoi(userInput)
	if err != nil {
		if len(userInput) == 0 {
			log.Fatal(err)
		}
		language.Language = userInput
		// create language
		data := url.Values{}
		data.Add("Language", userInput)
		response, err = http.Post(
			languageUrl,
			"application/x-www-form-urlencoded",
			strings.NewReader(data.Encode()))
		if err != nil {
			log.Fatal(err)
		}
		defer response.Body.Close()
		resultJson, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		result := parse.Result{}
		err = json.Unmarshal(resultJson, &result)
		if err != nil {
			log.Fatal(err)
		}
		language.Id = result.Id
	} else {
		for _, a := range languages {
			if a.Id == languageId {
				language = a
				return
			}
		}
		log.Fatalf("No language with Id: %d exists", languageId)
	}
	return
}

func userGetBook(authorId int) (book parse.Book) {
	response, err := http.Get(fmt.Sprintf(bookUrl, authorId))
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	booksJson, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	var books []parse.Book
	err = json.Unmarshal(booksJson, &books)
	if err != nil {
		log.Fatal(err)
	}
	for _, book := range books {
		fmt.Printf("Id: %d\tTitle: %s\n", book.Id, book.Title)
	}
	fmt.Println("Input Id to select an existing book, or new Title:")
	var userInput string
	_, err = fmt.Scanln(&userInput)
	if err != nil {
		log.Fatal(err)
	}
	bookId, err := strconv.Atoi(userInput)
	if err != nil {
		if len(userInput) == 0 {
			log.Fatal(err)
		}
		book.Title = userInput
		fmt.Printf("ISBN: ")
		n, err := fmt.Scanln(&userInput)
		if err != nil && n > 0 {
			log.Fatal(err)
		}
		book.ISBN.String = userInput
		fmt.Printf("Release date: ")
		n, err = fmt.Scanln(&userInput)
		if err != nil && n > 0 {
			log.Fatal(err)
		}
		if n > 0 {
			book.ReleaseDate, err = time.Parse(time.ANSIC, userInput)
			if err != nil {
				log.Fatal(err)
			}
		}
		// create book
		// TODO: I need the related Author, Topic and Language now
		data := url.Values{}
		data.Add("Title", book.Title)
		data.Add("ISBN", book.ISBN.String)
		if len(userInput) > 0 {
			data.Add("ReleaseDate", userInput)
		}
		// TODO: fix url for post
		response, err = http.Post(
			bookUrl,
			"application/x-www-form-urlencoded",
			strings.NewReader(data.Encode()))
		if err != nil {
			log.Fatal(err)
		}
		defer response.Body.Close()
		resultJson, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		result := parse.Result{}
		err = json.Unmarshal(resultJson, &result)
		if err != nil {
			log.Fatal(err)
		}
		book.Id = result.Id
	} else {
		for _, a := range books {
			if a.Id == bookId {
				book = a
				return
			}
		}
		log.Fatalf("No book with Id: %d exists", bookId)
	}
	return
}

func main() {
	// TODO: change order to make the user make as little inputs as required
	author := userGetAuthor()
	topic := userGetTopic()
	language := userGetLanguage()
	book := userGetBook(author.Id)
	fmt.Println(author, topic, language, book)
}
