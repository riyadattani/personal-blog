package blog

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/russross/blackfriday/v2"
	"html/template"
	"io"
	"io/ioutil"
	"strings"
	"time"
)

type Post struct {
	Title   string
	Content template.HTML
	Date    time.Time
	Picture string
	Tags    []string
}

func NewPost(fileName string) (Post, error) {
	title, content, date, picture, tags, err := createPost(fileName)
	if err != nil {
		return Post{}, err
	}

	const shortForm = "2006-Jan-02"
	formattedDate, err := time.Parse(shortForm, date)
	if err != nil {
		return Post{}, err
	}

	return Post{
		Title:   title,
		Content: template.HTML(content),
		Date:    formattedDate,
		Picture: picture,
		Tags:    tags,
	}, nil
}

func createPost(filename string) (title string, body []byte, date string, picture string, tags []string, err error) {
	fileContent, err := ioutil.ReadFile(fmt.Sprintf("../../cmd/web/posts/%s", filename))
	if err != nil {
		return "", nil, "", "", nil, nil
	}

	r := bytes.NewReader(fileContent)

	metaData := GetMetaData(r)
	title = metaData[0]
	date = metaData[1]
	picture = metaData[2]
	tags = strings.Split(metaData[3], ",")

	body = GetContentBody(fileContent)
	content := blackfriday.Run(body)

	return title, content, date, picture, tags, nil
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
