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

func NewPost(fileName string) Post {
	title, content, date := createPost(fileName)
	formattedDate, _ := time.Parse(layoutISO, date)

	return Post{
		Title:   title,
		Content: template.HTML(content),
		Date:    formattedDate,
	}
}

func createPost(filename string) (title string, body []byte, date string) {
	fileContent, err := ioutil.ReadFile(fmt.Sprintf("../../cmd/web/posts/%s", filename))
	if err != nil {
		log.Fatal(fmt.Sprint("Could not read markdown file, error: ", err))
	}

	r := bytes.NewReader(fileContent)

	metaData := GetMetaData(r)
	title = metaData[0]
	date = metaData[1]

	body = GetContentBody(fileContent)
	content := blackfriday.Run(body)

	return title, content, date
}

func GetMetaData(r io.Reader) []string {
	metaData := make([]string, 0)
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "-----" {
			break
		}
		metaData = append(metaData, line)
	}

	return metaData
}

func GetContentBody(byteArray []byte) []byte {
	content := bytes.Split(byteArray, []byte("-----\n"))[1]
	return content
}
