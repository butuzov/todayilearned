package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
)

type Dashboard struct {
	Books   json.RawMessage `json:"books,omitempty"`
	Authors json.RawMessage `json:"authors,omitempty"`

	Req json.RawMessage `json:"http_request,omitempty"`
}

type DashboardService interface {
	Get(data []string) (*Dashboard, error)
}

type defaultDashboardService struct {
	urlBooksService    string
	urlAuthorsServices string
	// service fields
	client *http.Client
	log    *log.Logger
}

func NewDashboardService(
	urlBooks string,
	urlAuthors string,
	logger *log.Logger,
) *defaultDashboardService {
	return &defaultDashboardService{
		client:             http.DefaultClient,
		urlBooksService:    urlBooks,
		urlAuthorsServices: urlAuthors,
		log:                logger,
	}
}

func (s defaultDashboardService) Get(data []string) (*Dashboard, error) {
	d := &Dashboard{}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(2)

	go func(ctx context.Context, wg *sync.WaitGroup) {
		defer wg.Done()

		books, err := s.readBooks(ctx)
		if err != nil {
			s.log.Errorf("books error: %v", err)

			return
		}
		d.Books = books
	}(ctx, &wg)

	go func(ctx context.Context, wg *sync.WaitGroup) {
		defer wg.Done()

		authors, err := s.readAuthors(ctx)
		if err != nil {
			s.log.Errorf("authors error: %v", err)
			return
		}
		d.Authors = authors
	}(ctx, &wg)

	wg.Wait()

	tmp, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	d.Req = json.RawMessage(tmp)

	return d, nil
}

func (s defaultDashboardService) readAuthors(ctx context.Context) (authors json.RawMessage, err error) {
	url := s.urlAuthorsServices + "authors/"
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("authors: %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("authors: %w", err)
	}

	defer resp.Body.Close()

	fmt.Println("authors", url)

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("authors: %w", err)
	}

	s.log.Infof("authors: content(%s)", string(data))

	if err := json.Unmarshal(data, &authors); err != nil {
		return nil, fmt.Errorf("authors: %w", err)
	}

	return authors, nil
}

func (s defaultDashboardService) readBooks(ctx context.Context) (books json.RawMessage, err error) {
	url := s.urlBooksService + "books"
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("books: %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("books: %w", err)
	}
	defer resp.Body.Close()

	s.log.Infof("books: url(%s)", url)

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("books: %w", err)
	}

	s.log.Infof("books: content(%s)", string(data))

	if err := json.Unmarshal(data, &books); err != nil {
		return nil, fmt.Errorf("books: %w", err)
	}

	return books, nil
}
