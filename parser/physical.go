package parse

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

const PhysicalType = "-physical"

func ParsePhysical(file *os.File, book Book) (quotes []Quote) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		text = text[1 : len(text)-2]
		columns := strings.Split(text, "\",\"")
		if len(columns) != 2 {
			continue
		}
		quoteText := columns[0]
		page, err := strconv.Atoi(columns[1])
		if err != nil {
			log.Fatal(err)
		}
		quote := Quote{
			BookId: book.Id,
			Page:   page,
			Quote:  quoteText,
		}
		quotes = append(quotes, quote)
	}
	return
}
