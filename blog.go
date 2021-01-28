package blog

import (
	"fmt"
	"github.com/russross/blackfriday/v2"
	"html/template"
	"io/ioutil"
	"log"
)

type Blog struct {
	Title   string
	Content template.HTML
}

func NewBlog(title string) Blog {
	output := loadBlog(title)

	return Blog{
		Title:   title,
		Content: template.HTML(output),
	}
}

func loadBlog(title string) []byte {
	body, err := ioutil.ReadFile(fmt.Sprintf("../web/blogs/%s", title))
	if err != nil {
		log.Fatal(fmt.Sprint("Could not read blog file ", err))
	}
	output := blackfriday.Run(body)
	return output
}
