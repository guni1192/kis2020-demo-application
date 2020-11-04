package handler

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTopPageHandler(t *testing.T) {
	expectReqBody := "Hello world"
	reqBody := bytes.NewBufferString(expectReqBody)
	req := httptest.NewRequest(http.MethodGet, "/", reqBody)

	r := httptest.NewRecorder()

	TopPage(r, req)

	if r.Code != http.StatusOK {
		t.Errorf("expect(%d), acctual(%d)", http.StatusOK, r.Code)
	}

	if r.Body.String() != expectReqBody {
		t.Errorf("expect(%s), acctual(%s)", expectReqBody, r.Body.String())
	}

}
