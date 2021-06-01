package main

import (
	"fmt"
	"io/ioutil"
)

// basic page structure
type Page struct {
	Title string
	Body  []byte
}

// saving a page to disk as file appengin .txt
// This is a method named save that takes a receiver p, a pointer to Page
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

// loadPage constructs the file name from the title parameter, reads
// the files contents into new variable body and returns a pointer
// the a Page literal constructed with proper title and body values
// will return also an error if file not exists or cannot be read
func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	// catch the error
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func main() {

	p1 := &Page{Title: "TestPage", Body: []byte("This is a sample page")}
	p1.save()
	p2, _ := loadPage("TestPage")
	fmt.Println(string(p2.Body))
}
