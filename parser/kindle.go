package parse

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

const KindleType = "-kindle"

// This function parses the modified version of the email received csv file
// where it does not have the heading columns and only contains the marked
// quotes
func ParseKindle(file *os.File, book Book) (quotes []Quote) {
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
		quote := Quote{
			BookId: book.Id,
			Page:   page,
			Quote:  quoteText,
		}
		quotes = append(quotes, quote)
	}
	return
}
