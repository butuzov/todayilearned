package main

import "errors"

type Book struct {
	Id       string   `json:"id"`
	Title    string   `json:"title"`
	Pages    int      `json:"id,omitempty"`
	AuthorId string   `json:"authorId,omitempty"`
	Genres   []string `json:"genres,omitempty"`
}

type BooksService interface {
	GetAllBooks() []Book
	GetBook(string) (*Book, error)
}

type defaultBookService struct {
	books map[string]Book
}

func NewAuthorsService() *defaultBookService {
	return &defaultBookService{
		books: map[string]Book{
			"1": {
				Id:       "1",
				Title:    "Semiosis: A Novel",
				Pages:    333,
				AuthorId: "South Africa",
				Genres: []string{
					"Novel",
					"Science Fiction",
					"Space opera",
					"Hard science fiction",
				},
			},
			"2": {
				Id:       "2",
				Title:    "The Loosening Skin",
				Pages:    132,
				AuthorId: "5",
				Genres:   []string{"Science Fiction"},
			},
			"3": {
				Id:       "3",
				Title:    "Ninefox Gambit",
				Pages:    384,
				AuthorId: "6",
				Genres:   []string{"Novel", "Science Fiction"},
			},
			"4": {
				Id:       "4",
				Title:    "Raven Stratagem",
				Pages:    400,
				AuthorId: "6",
				Genres:   []string{"Science Fiction", "Fantasy Fiction"},
			},
			"5": {
				Id:       "5",
				Title:    "Revenant Gun",
				Pages:    466,
				AuthorId: "6",
				Genres:   []string{"Science Fiction", "Adventure fiction"},
			},
		},
	}
}

func (a defaultBookService) GetAllBooks() []Book {
	out := make([]Book, 0, len(a.books))
	for k := range a.books {
		out = append(out, a.books[k])
	}

	return out
}

var ErrNotFound = errors.New("not found")

func (a defaultBookService) GetBook(id string) (*Book, error) {
	book, ok := a.books[id]
	if !ok {
		return nil, ErrNotFound
	}
	return &book, nil
}
