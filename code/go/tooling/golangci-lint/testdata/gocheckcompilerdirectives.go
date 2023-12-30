package testdata

//go:generate echo hello world

//go:embed
var Value string

//go:

// Problematic cases:

// go:embed // want "compiler directive contains space: // go:embed"

//    go:embed // want "compiler directive contains space: //    go:embed"

//go:genrate // want "compiler directive unrecognized: //go:genrate"