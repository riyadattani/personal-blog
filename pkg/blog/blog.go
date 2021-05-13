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

	metData, content, err := CreatePost(fileContent)
	if err != nil {
		return Post{}, err
	}

	return Post{
		Title:   metData.Title,
		Content: content,
		Date:    metData.Date,
		Picture: metData.Picture,
		Tags:    metData.Tags,
	}, nil
}

func CreatePost(fileContent []byte) (metaData MetaData, content template.HTML, err error) {
	metaData, err = getMetaData(bytes.NewReader(fileContent))
	if err != nil {
		return MetaData{}, "", err
	}

	content = getContent(fileContent)

	return metaData, content, nil
}

type MetaData struct {
	Title   string
	Date    time.Time
	Picture string
	Tags    []string
}

func getMetaData(r io.Reader) (MetaData, error) {
	metaData := MetaData{}
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	readLine := func() string {
		scanner.Scan()
		return scanner.Text()
	}

	metaData.Title = readLine()
	date, err := stringToDate(readLine())
	if err != nil {
		return MetaData{}, err
	}
	metaData.Date = date
	metaData.Picture = readLine()
	metaData.Tags = strings.Split(readLine(), ",")

	return metaData, nil
}

func getContent(byteArray []byte) template.HTML {
	body := bytes.Split(byteArray, []byte("-----\n"))[1]
	content := blackfriday.Run(body, blackfriday.WithRenderer(renderer()), blackfriday.WithExtensions(blackfriday.CommonExtensions))
	return template.HTML(content)
}

func renderer() *bfchroma.Renderer {
	return bfchroma.NewRenderer(
		bfchroma.WithoutAutodetect(),
		bfchroma.ChromaOptions(
			html.WithLineNumbers(false),
		),
		bfchroma.Extend(
			blackfriday.NewHTMLRenderer(blackfriday.HTMLRendererParameters{
				Flags: blackfriday.CommonHTMLFlags,
			}),
		),
		bfchroma.Style("lovelace"),
	)
}

func stringToDate(stringDate string) (time.Time, error) {
	const shortFormDate = "2006-Jan-02"
	date, err := time.Parse(shortFormDate, stringDate)
	if err != nil {
		return time.Time{}, err
	}
	return date, nil
}
