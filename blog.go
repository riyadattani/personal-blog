package blog

import (
	"fmt"
	"github.com/russross/blackfriday/v2"
	"html/template"
	"io/ioutil"
	"log"
)

type Post struct {
	Title   string
	Content template.HTML
}

func NewPost(title string) Post {
	output := loadPost(title)

	return Post{
		Title:   title,
		Content: template.HTML(output),
	}
}

func loadPost(title string) []byte {
	body, err := ioutil.ReadFile(fmt.Sprintf("../web/posts/%s", title))
	if err != nil {
		log.Fatal(fmt.Sprint("Could not read blog file ", err))
	}
	output := blackfriday.Run(body)
	return output
}
