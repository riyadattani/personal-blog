package blog

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/Depado/bfchroma"
	"github.com/alecthomas/chroma/formatters/html"
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
	fileContent, err := ioutil.ReadFile(fmt.Sprintf("../../cmd/web/posts/%s", fileName))
	if err != nil {
		return Post{}, err
	}

	title, body, date, picture, tags, err := CreatePost(fileContent)
	if err != nil {
		return Post{}, err
	}

	renderer := bfchroma.NewRenderer(
		bfchroma.WithoutAutodetect(),
		bfchroma.ChromaOptions(
			html.WithLineNumbers(true),
		),
		bfchroma.Extend(
			blackfriday.NewHTMLRenderer(blackfriday.HTMLRendererParameters{
				Flags: blackfriday.CommonHTMLFlags,
			}),
		),
		bfchroma.Style("lovelace"),
	)

	content := blackfriday.Run(body, blackfriday.WithRenderer(renderer), blackfriday.WithExtensions(blackfriday.CommonExtensions))

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

func CreatePost(fileContent []byte) (title string, body []byte, date string, picture string, tags []string, err error) {
	r := bytes.NewReader(fileContent)

	metaData := getMetaData(r)
	title = metaData[0]
	date = metaData[1]
	picture = metaData[2]
	tags = strings.Split(metaData[3], ",")

	body = getContentBody(fileContent)

	return title, body, date, picture, tags, nil
}

func getMetaData(r io.Reader) []string {
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

func getContentBody(byteArray []byte) []byte {
	content := bytes.Split(byteArray, []byte("-----\n"))[1]
	return content
}
