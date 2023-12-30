// file-header-check
package testdata

func fooFlagParam(a bool, b int) { // MATCH /parameter 'a' seems to be a control flag, avoid control coupling/
	if a {
	}
}
