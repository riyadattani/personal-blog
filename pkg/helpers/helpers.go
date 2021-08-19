package helpers

import (
	"github.com/Depado/bfchroma"
	"github.com/alecthomas/chroma/formatters/html"
	"github.com/russross/blackfriday/v2"
	"html/template"
	"time"
)

func RenderMarkdown(body []byte) template.HTML {
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
		bfchroma.Style("solarized-dark"),
	)
}

func StringToDate(stringDate string) (time.Time, error) {
	const shortFormDate = "2006-Jan-02"
	date, err := time.Parse(shortFormDate, stringDate)
	if err != nil {
		return time.Time{}, err
	}
	return date, nil
}
