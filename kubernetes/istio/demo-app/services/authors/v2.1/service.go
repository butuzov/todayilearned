package main

import "errors"

type Author struct {
	Id        string `json="id,omitempty"`
	FirstName string `json="firstName,omitempty"`
	LastName  string `json="lastName,omitempty"`
	Born      string `json="born,omitempty"`
	AmazonID  string `json="amazon_id,omitempty"`
}

type AuthorsService interface {
	GetAllAuthors() []Author
	GetAuthor(string) (*Author, error)
}

type authorsService struct {
	authors map[string]Author
}

func NewAuthorsService() *authorsService {
	return &authorsService{
		authors: map[string]Author{
			"1": {
				Id:        "1",
				FirstName: "Loreth Anne",
				LastName:  "White",
				Born:      "South Africa",
				AmazonID:  "B001HMNJRG",
			},
			"2": {
				Id:        "2",
				FirstName: "Lisa",
				LastName:  "Regan",
				Born:      "USA",
			},
			"3": {
				Id:        "3",
				FirstName: "Ty",
				LastName:  "Patterson",
				Born:      "USA",
				AmazonID:  "B00APU24BI",
			},
			"4": {
				Id:        "4",
				FirstName: "Sue",
				LastName:  "Burke",
				Born:      "USA",
				AmazonID:  "B007IG9WGW",
			},
			"5": {
				Id:        "5",
				FirstName: "Aliya",
				LastName:  "Whiteley",
				Born:      "UK",
				AmazonID:  "B0034P9Y28",
			},
			"6": {
				Id:        "6",
				FirstName: "Yoon Ha",
				LastName:  "Lee",
				Born:      "USA",
				AmazonID:  "B00IORCR7U",
			},
		},
	}
}

func (a authorsService) GetAllAuthors() []Author {
	out := make([]Author, 0, len(a.authors))
	for k := range a.authors {
		out = append(out, a.authors[k])
	}

	return out
}

var ErrNotFound = errors.New("not found")

func (a authorsService) GetAuthor(id string) (*Author, error) {
	for k := range a.authors {
		if k == id {
			author := a.authors[k]
			return &author, nil
		}
	}

	return nil, ErrNotFound
}
