package main

import (
	"fmt"
	"net/http"
	"testing"
)

func TestNoSuf(t *testing.T) {
	var myH myHandler

	h := NoSurf(&myH)

	switch v := h.(type) {
	case http.Handler:
		//Do nothing
	default:
		t.Error(fmt.Sprintf("Type is not http.handler, type is %T", v))
	}
}

func TestSessionLoad(t *testing.T) {
	var myH myHandler

	h := NoSurf(&myH)

	switch v := h.(type) {
	case http.Handler:
		//Do nothing
	default:
		t.Error(fmt.Sprintf("Type is not http.handler, type is %T", v))
	}
}
