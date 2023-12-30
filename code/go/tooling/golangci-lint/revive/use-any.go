// file-header-check
package testdata

type afoo = interface{} // MATCH /since GO 1.18 'interface{}' can be replaced by 'any'/
