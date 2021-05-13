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

// TODO: this should take a reader
func NewPost(fileName string) (Post, error) {
	fileContent, err := ioutil.ReadFile(fmt.Sprintf("../../cmd/web/posts/%s", fileName))
	if err != nil {
		return Post{}, err
	}

	post, err := CreatePost(bytes.NewReader(fileContent))
	if err != nil {
		return Post{}, err
	}

	return post, nil
}

func CreatePost(fileContent io.Reader) (Post, error) {
	post, err := getPost(fileContent)
	if err != nil {
		return Post{}, err
	}

	return post, nil
}

func getPost(r io.Reader) (Post, error) {
	post := Post{}
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	readLine := func() string {
		scanner.Scan()
		return scanner.Text()
	}

	post.Title = readLine()
	date, err := stringToDate(readLine())
	if err != nil {
		return Post{}, err
	}
	post.Date = date
	post.Picture = readLine()
	post.Tags = strings.Split(readLine(), ",")
	readLine()

	body := bytes.Buffer{}

	for scanner.Scan() {
		body.Write(scanner.Bytes())
	}
	post.Content = renderMarkdown(body.Bytes())

	return post, nil
}

func renderMarkdown(body []byte) template.HTML {
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
