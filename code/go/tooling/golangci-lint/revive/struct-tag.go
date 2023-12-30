// file-header-check
package testdata

type decodeAndValidateRequest struct {
	// BEAWRE : the flag of URLParam should match the const string URLParam
	URLParam   string `json:"-" path:"url_param" validate:"numeric"`
	Text       string `json:"text" validate:"max=10"`
	DefaultInt int    `json:"defaultInt" default:"10.0"` // MATCH /field's type and default
}
