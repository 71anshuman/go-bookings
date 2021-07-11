package forms

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {
	r := httptest.NewRequest("POST", "/anything", nil)

	form := New(r.PostForm)

	isValid := form.Valid()

	if !isValid {
		t.Error("got invalid when should have been valid")
	}
}

func TestForm_Required(t *testing.T) {
	r := httptest.NewRequest("POST", "/requiredTest", nil)

	form := New(r.PostForm)

	form.Required("name", "date", "dob")
	if form.Valid() {
		t.Error("this isn't valid as req fields are missing")
	}

	postData := url.Values{}
	postData.Add("name", "anshuman")
	postData.Add("email", "71anshuman@hotmail.com")
	postData.Add("a", "b")

	r, _ = http.NewRequest("POST", "/whatever", nil)
	r.PostForm = postData
	form = New(r.PostForm)
	form.Required("name", "email", "a")
	if !form.Valid() {
		t.Error("it has req fields and show it doesn't")
	}
}

func TestForm_Has(t *testing.T) {
	r, _ := http.NewRequest("POST", "/has", nil)
	form := New(r.Form)
	if form.Has("a") {
		t.Error("form shows has field, when it does not")
	}

	postData := url.Values{}
	postData.Add("a", "a")
	form = New(postData)

	if !form.Has("a") {
		t.Error("form has field and it shows it does not")
	}

}

func TestForm_MinLength(t *testing.T) {
	form := New(url.Values{})
	form.MinLength("lastname", 3)
	if form.Valid() {
		t.Error("the expected length not match, still pass")
	}

	isError := form.Errors.Get("lastname")

	if isError == "" {
		t.Error("Expected error but didn't get that")
	}

	postData := url.Values{}
	postData.Add("first_name", "Anshuman")

	form = New(postData)
	form.MinLength("first_name", 3)
	if !form.Valid() {
		t.Error("min length is expected, still failed")
	}

	isError = form.Errors.Get("first_name")
	if isError != "" {
		t.Error("Not Expected error but get that")
	}
}

func TestForm_IsEmail(t *testing.T) {
	postData := url.Values{}
	postData.Add("email", "71anshuman@gmail.com")

	form := New(postData)
	form.IsEmail("email")
	if !form.Valid() {
		t.Error("email is correct, still gets error")
	}

	postData.Set("email", "71annshuman")
	form = New(postData)
	form.IsEmail("email")
	if form.Valid() {
		t.Error("email isn't correct, still gets pass")
	}
}