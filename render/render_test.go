package render

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

var pageData = []struct {
	name          string
	renderer      string
	template      string
	errorExpected bool
	errorMessage  string
}{
	{"go_page", "go", "home", false, "error rending go template"},
	{"go_page_notemplate", "go", "no-file", true, "no error rending non existent go template"},
	{"jet_page", "jet", "home", false, "error rending jet template"},
	{"jet_page_notemplate", "jet", "no-file", true, "no error rending non existent jet template"},
	{"invalid_renderer_engine", "fish", "hone", true, "no error using non existent template engine"},
}

func TestRender_Page(t *testing.T) {
	for _, e := range pageData {
		r, err := http.NewRequest("GET", "/some-url", nil)
		if err != nil {
			t.Error(err)
		}

		w := httptest.NewRecorder()

		testRenderer.Renderer = e.renderer
		testRenderer.RootPath = "./testdata"

		err = testRenderer.Page(w, r, e.template, nil, nil)
		if e.errorExpected {
			if err == nil {
				t.Errorf("%s: %s", e.name, e.errorMessage)
			}
		} else {
			if err != nil {
				t.Errorf("%s: %s: %s", e.name, e.errorMessage, err.Error())
			}
		}
	}
}

func TestRender_GoPage(t *testing.T) {
	w := httptest.NewRecorder()
	r, err := http.NewRequest("GET", "/ur/", nil)
	if err != nil {
		t.Error(err)
	}

	testRenderer.Renderer = "go"
	testRenderer.RootPath = "./testdata"

	err = testRenderer.GoPage(w, r, "home", nil)
	if err != nil {
		t.Error("Error rendering page", err)
	}

	err = testRenderer.GoPage(w, r, "no-file", nil)
	if err == nil {
		t.Error("Error rendering non existent page", err)
	}

	testRenderer.Renderer = "fish"
	err = testRenderer.Page(w, r, "home", nil, nil)
	if err == nil {
		t.Error("no error returned rendering with invalid engine", err)
	}

}

func TestRender_JetPage(t *testing.T) {
	w := httptest.NewRecorder()
	r, err := http.NewRequest("GET", "/ur/", nil)
	if err != nil {
		t.Error(err)
	}
	testRenderer.Renderer = "jet"
	testRenderer.RootPath = "./testdata"

	err = testRenderer.JetPage(w, r, "home", nil, nil)
	if err != nil {
		t.Error("Error rendering jet page", err)
	}

	err = testRenderer.JetPage(w, r, "no-file", nil, nil)
	if err == nil {
		t.Error("Error rendering non-existent jet page", err)
	}

}
