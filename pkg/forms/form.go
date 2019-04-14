package forms

import (
	"fmt"
	"net/url"
	"strings"
	"unicode/utf8"
)

//Form contains the data and errors for a form submission
type Form struct {
	url.Values
	Errors errors
}

//New initializes a new custom form struct
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

//Required validates that the given form fields are not blank
func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This field cannot be blank")
		}
	}
}

//MaxLength verifies that a form field no longer than the max length
func (f *Form) MaxLength(field string, d int) {
	value := f.Get(field)
	if value == "" {
		return
	}
	if utf8.RuneCountInString(value) > d {
		f.Errors.Add(field, fmt.Sprintf("This field is too long (maximum is %d characters)", d))
	}
}

//PermittedValues makes sure that only the specified fields are permitted
func (f *Form) PermittedValues(field string, opts ...string) {
	value := f.Get(field)
	if value == "" {
		return
	}

	for _, opt := range opts {
		if value == opt {
			return
		}
	}
	f.Errors.Add(field, "This field is invalid")
}

//Valid checks if the given form data is valid
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}
