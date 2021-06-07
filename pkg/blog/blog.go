package blog

import (
	"bufio"
	"bytes"
	"html/template"
	"io"
	"personal-blog/pkg/helpers"
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

type Posts []Post

func NewPost(fileContent io.Reader) (Post, error) {
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
	date, err := helpers.StringToDate(readLine())
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
		body.WriteString("\n")
	}
	post.Content = helpers.RenderMarkdown(body.Bytes())

	return post, nil
}
