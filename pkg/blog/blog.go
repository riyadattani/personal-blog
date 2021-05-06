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

	metData, body, err := CreatePost(fileContent)
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
	formattedDate, err := time.Parse(shortForm, metData.Date)
	if err != nil {
		return Post{}, err
	}

	return Post{
		Title:   metData.Title,
		Content: template.HTML(content),
		Date:    formattedDate,
		Picture: metData.Picture,
		Tags:    metData.Tags,
	}, nil
}

func CreatePost(fileContent []byte) (metaData MetaData, body []byte, err error) {
	r := bytes.NewReader(fileContent)

	metaData = getMetaData(r)
	body = getContentBody(fileContent)

	return metaData, body,nil
}

//TODO: return a structure rather than a string
// do not use a for loop. You only need a for loop when reading the body because you dont know how many lines that will be

type MetaData struct {
	Title   string
	Date    string
	Picture string
	Tags    []string
}

func getMetaData(r io.Reader) MetaData {
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

	return MetaData{
		Title:   metaData[0],
		Date:    metaData[1],
		Picture: metaData[2],
		Tags:    strings.Split(metaData[3], ","),
	}
}

func getContentBody(byteArray []byte) []byte {
	content := bytes.Split(byteArray, []byte("-----\n"))[1]
	return content
}
