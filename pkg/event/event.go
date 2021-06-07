package event

import (
	"bufio"
	"bytes"
	"html/template"
	"io"
	"personal-blog/pkg/helpers"
	"strings"
	"time"
)

type Event struct {
	Title   string
	Description template.HTML
	Date    time.Time
	Picture string
	Tags    []string
}

type Events []Event

func New(fileContent io.Reader) (Event, error) {
	event, err := getEvent(fileContent)
	if err != nil {
		return Event{}, err
	}

	return event, nil
}

func getEvent(r io.Reader) (Event, error) {
	post := Event{}
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	readLine := func() string {
		scanner.Scan()
		return scanner.Text()
	}

	post.Title = readLine()
	date, err := helpers.StringToDate(readLine())
	if err != nil {
		return Event{}, err
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
	post.Description = helpers.RenderMarkdown(body.Bytes())

	return post, nil
}
