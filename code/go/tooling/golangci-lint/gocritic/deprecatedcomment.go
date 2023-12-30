package testdata_gocritic

// Example for the rule: deprecatedComment

// BadFormat1 is an example.
/*! the proper format is `Deprecated: <text>` */
// This function is deprecated, use XYZ instead.
func BadFormat1() {}
