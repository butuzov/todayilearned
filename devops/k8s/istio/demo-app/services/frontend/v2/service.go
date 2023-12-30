package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Book struct {
	Id       string   `json:"id"`
	Title    string   `json:"title"`
	Pages    int      `json:"id,omitempty"`
	AuthorId string   `json:"authorId,omitempty"`
	Genres   []string `json:"genres,omitempty"`
}

type Author struct {
	Id        string `json="id,omitempty"`
	FirstName string `json="firstName,omitempty"`
	LastName  string `json="lastName,omitempty"`
	Born      string `json="born,omitempty"`
}

type Dashboard struct {
	Books   []Book   `json:"books"`
	Authors []Author `json:"authors"`
}

type DashboardService interface {
	Get() (*Dashboard, error)
}

type defaultDashboardService struct {
	clint              http.Client
	urlBooksService    string
	urlAuthorsServices string
}

func NewDashboardService(
	urlBooks string,
	urlAuthors string,
) *defaultDashboardService {
	return &defaultDashboardService{
		urlBooksService:    urlBooks,
		urlAuthorsServices: urlAuthors,
	}
}

func (s defaultDashboardService) Get() (*Dashboard, error) {
	d := &Dashboard{}
	if books, err := s.readBooks(); err != nil {
		return nil, err
	} else {
		d.Books = books
	}

	if authors, err := s.readAuthors(); err != nil {
		return nil, err
	} else {
		d.Authors = authors
	}

	return d, nil
}

func (s defaultDashboardService) readAuthors() (authors []Author, err error) {
	url := s.urlAuthorsServices + "authors"
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("authors: %w", err)
	}
	defer resp.Body.Close()

	fmt.Println("authors", url)

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("authors: %w", err)
	}

	fmt.Println("authors", url, string(data))

	if err := json.Unmarshal(data, &authors); err != nil {
		return nil, fmt.Errorf("authors: %w", err)
	}

	return authors, nil
}

func (s defaultDashboardService) readBooks() (books []Book, err error) {
	url := s.urlBooksService + "books"
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("books: %w", err)
	}
	defer resp.Body.Close()

	fmt.Println("books", url)

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("books: %w", err)
	}

	fmt.Println("books", url, string(data))

	if err := json.Unmarshal(data, &books); err != nil {
		return nil, fmt.Errorf("books: %w", err)
	}

	return books, nil
}
