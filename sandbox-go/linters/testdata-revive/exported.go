// Package golint comment
// file-header-check
package testdata

// Test case for the configuration option tp replace the word "stutters" by "repetitive" failure messages

//  GolintRepetitive is a dummy function
func TestdataRepetitive() {} // MATCH /func name will be used as golint.GolintRepetitive by other packages, and that is repetitive; consider calling this Repetitive/

// Test cases for enabling checks of exported methods of private types in exported rule
type private struct{}

// MATCH /comment on exported method private.Method should be of the form "Method ..."/
func (p *private) Method() {
}
