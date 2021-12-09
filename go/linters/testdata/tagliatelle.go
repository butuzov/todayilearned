package testdata

// json and camel case
type Foo struct {
	ID     string `json:"ID"`     // must be "id"
	UserID string `json:"UserID"` // must be "userId"
	Name   string `json:"name"`
	Value  string `json:"val,omitempty"` // must be "value"
}
