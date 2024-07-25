package models

type Author struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	// Books     []*Book `json:"books"`
}

type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

var Books []Book

func CreateBookInstance(id string, isbn string, title string, author *Author) *Book {
	return &Book{
		ID:     id,
		Isbn:   isbn,
		Title:  title,
		Author: author,
	}
}

func CreateAuthorInstance(firstName string, lastName string) *Author {
	return &Author{
		FirstName: firstName,
		LastName:  lastName,
	}
}
