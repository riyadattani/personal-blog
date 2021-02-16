package blog

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/russross/blackfriday/v2"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"strings"
	"time"
)

type Post struct {
	Title   string
	Content template.HTML
	Date    time.Time
}

const (
	layoutISO = "2006-01-15"
)

func NewPost(title string) Post {
	output, dateString := loadPost(title)

	date, _ := time.Parse(layoutISO, dateString)

	return Post{
		Title:   title,
		Content: template.HTML(output),
		Date:    date,
	}
}

func loadPost(title string) (postBody []byte, date string) {
	body, err := ioutil.ReadFile(fmt.Sprintf("../web/posts/%s", title))
	if err != nil {
		log.Fatal(fmt.Sprint("Could not read blog file ", err))
	}

	r := bytes.NewReader(body)
	date, _, err = readLine(r, 2)
	if err != nil {
		log.Fatal(fmt.Sprint("Could not read the date in the post", err))
	}


	output := blackfriday.Run(body)
	return output, date
}

func readLine(r io.Reader, lineNum int) (line string, lastLine int, err error) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		lastLine++
		if lastLine == lineNum {
			return scanner.Text(), lastLine, scanner.Err()
		}
	}
	return line, lastLine, io.EOF
}

func getMetaData(r io.Reader) string {
	metaData := make([]string, 0)
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		//fmt.Printf("%s\n", line)

		// Break if we hit line break.
		if line == "-----" {
			break
		}

		metaData = append(metaData, line)
	}

	return strings.Join(metaData, "\n")
}
