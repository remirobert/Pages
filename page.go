package pages

import (
	"io"
	"io/ioutil"
	"text/template"

	"github.com/russross/blackfriday"
)

type Page struct {
	Title  string
	Author string
	HTML   string
	File   string
}

const configFile = "config.json"

func NewPage(configPath string) (*Page, error) {
	p := &Page{File: configPath}
	return p, p.load()
}

func (p *Page) Render(layout string, w io.Writer) error {
	t, err := template.ParseFiles(layout)
	if err != nil {
		return err
	}
	return t.Execute(w, p)
}

func (p *Page) load() error {
	c, err := LoadConfig(configFile)
	if err != nil {
		return err
	}
	data, err := ioutil.ReadFile(p.File)
	if err != nil {
		return err
	}
	p.HTML = string(blackfriday.MarkdownCommon(data))
	p.Title = c.Name()
	p.Author = c.Author()
	return nil
	return nil
}
