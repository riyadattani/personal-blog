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

func NewPost(title string) Post {
	output, dateString := loadPost(title)

	date, _ := time.Parse(layoutISO, dateString)
	//formattedDate := date.Format(layoutUS)

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
	date, _, err = readLine(r, 3)
	if err != nil {
		log.Fatal(fmt.Sprint("Could not read the date in the post", err))
	}

	output := blackfriday.Run(body)
	return output, date
}

func readLine(r io.Reader, lineNum int) (line string, lastLine int, err error) {
	sc := bufio.NewScanner(r)
	for sc.Scan() {
		lastLine++
		if lastLine == lineNum {
			return sc.Text(), lastLine, sc.Err()
		}
	}
	return line, lastLine, io.EOF
}
