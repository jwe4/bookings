package forms

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)
	isValid := form.Valid()
	if !isValid {
		t.Error("got invalid, but expected valid")
	}
}

func TestForm_Required(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	form.Required("a", "b", "c")
	if form.Valid() {
		t.Error("form shows valid when required fields missing")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")
	postedData.Add("b", "b")
	postedData.Add("c", "c")
	r, _ = http.NewRequest("POST", "/whatever", nil)
	r.PostForm = postedData

	form = New(r.PostForm)
	form.Required("a", "b", "c")
	if !form.Valid() {
		t.Error("shows does not have required fields when it does")
	}
}

// TestForm_Has tests the Has method
func TestForm_Has(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)
	has := form.Has("whatever")
	if has {
		t.Error("form shows has field when it does not")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")
	form = New(postedData)
	has = form.Has("a")
	if !has {
		t.Error("shows from does not have field when it should")
	}

}

func TestForm_MinLength(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	form.MinLength("x", 10)
	if form.Valid() {
		t.Error("form shows min length for non-existent field")
	}

	isError := form.Errors.Get("x")

	if isError == "" {
		t.Error("should have an error, but did not get one")
	}

	postedValues := url.Values{}
	postedValues.Add("some_field", "some value")
	form = New(postedValues)
	form.MinLength("some_field", 100)
	if form.Valid() {
		t.Error("shows minlength of 100 met when data is shorter")
	}
	form = New(postedValues)
	form.MinLength("some_field", 1)
	if !form.Valid() {
		t.Error("shows minlength of 1 is not met when it is")
	}

	isError = form.Errors.Get("x")

	if isError != "" {
		t.Error("should not have an error, but did  get one")
	}
}

func TestForm_IsEmail(t *testing.T) {
	postedValues := url.Values{}
	form := New(postedValues)
	form.IsEmail("x")
	if form.Valid() {
		t.Error("form shows valid email for non-existent field")
	}
	postedValues = url.Values{}
	postedValues.Add("email", "abc")
	form = New(postedValues)
	form.IsEmail("abc")
	if form.Valid() {
		t.Error("form shows valid email for invalid email field")
	}
	postedValues.Add("abc", "joe@abc.com")
	form = New(postedValues)
	form.IsEmail("abc")
	if !form.Valid() {
		t.Error("form shows invalid email for valid email field")
	}
}

func TestNew(t *testing.T) {
	data := url.Values{}
	data.Add("key", "value")

	form := New(data)

	if form == nil {
		t.Fatal("New returned nil")
	}

	if form.Get("key") != "value" {
		t.Errorf("Expected form data 'key' to be 'value', got '%s'", form.Get("key"))
	}

	if form.Errors == nil {
		t.Error("Expected form.Errors to be initialized, but it was nil")
	}

	if len(form.Errors) != 0 {
		t.Errorf("Expected form.Errors to be empty, but got length %d", len(form.Errors))
	}
}
