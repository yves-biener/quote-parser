package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	parse "parse/parser"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	rootUrl     = "http://localhost:8000/api"
	authorUrl   = rootUrl + "/authors"
	topicUrl    = rootUrl + "/topics"
	languageUrl = rootUrl + "/languages"
	bookUrl     = rootUrl + "/books"
	quotesUrl   = rootUrl + "/quotes"
)

func postQuote(quote parse.Quote) {
	data := url.Values{}
	data.Add("BookId", fmt.Sprintf("%d", quote.BookId))
	data.Add("Page", fmt.Sprintf("%d", quote.Page))
	data.Add("Quote", quote.Quote)
	response, err := http.Post(
		quotesUrl,
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
	quote.Id = result.Id
}

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

func userGetBook() (book parse.Book) {
	response, err := http.Get(bookUrl)
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
		author := userGetAuthor()
		topic := userGetTopic()
		language := userGetLanguage()
		data := url.Values{}
		data.Add("Title", book.Title)
		data.Add("ISBN", book.ISBN.String)
		data.Add("AuthorId", fmt.Sprintf("%d", author.Id))
		data.Add("TopicId", fmt.Sprintf("%d", topic.Id))
		data.Add("LanguageId", fmt.Sprintf("%d", language.Id))
		if len(userInput) > 0 {
			data.Add("ReleaseDate", userInput)
		}
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

// TODO: move to other file for specific parsing
func parseKindleCSV(file *os.File, book parse.Book) (quotes []parse.Quote) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		text = text[1 : len(text)-2]
		columns := strings.Split(text, "\",\"")
		if len(columns) != 4 {
			continue
		}
		typ := columns[0]
		if !strings.HasPrefix(typ, "Markierung") {
			continue
		}
		pageSplit := strings.Split(columns[1], " ")
		page, err := strconv.Atoi(pageSplit[1])
		if err != nil {
			log.Fatal(err)
		}
		quoteText := columns[3]
		quote := parse.Quote{
			BookId: book.Id,
			Page:   page,
			Quote:  quoteText,
		}
		quotes = append(quotes, quote)
	}
	return
}

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		log.Fatal("Usage: quote-parser -<type> <pathToCsvWithQuotes>")
	}
	book := userGetBook()
	var parsingFunction func(*os.File, parse.Book) []parse.Quote
	switch args[0] {
	case "-kindle":
		parsingFunction = parseKindleCSV
	default:
		log.Fatalf("Unknown type: %s", args[0])
	}
	file, err := os.Open(args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	quotes := parsingFunction(file, book)
	var wg sync.WaitGroup
	for _, quote := range quotes {
		go func(q parse.Quote, w *sync.WaitGroup) {
			w.Add(1)
			defer w.Done()
			postQuote(q)
		}(quote, &wg)
	}
	for _, quote := range quotes {
		fmt.Println(quote)
	}
	wg.Wait()
}
