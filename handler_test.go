package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestServe(t *testing.T) {
	server := httptest.NewServer(SetupHandler())
	defer server.Close()

	testPostMessage(t, server.URL+"/drawer1/messages/new", "john", "testmsg1")
	testPostMessage(t, server.URL+"/drawer2/messages/new", "smith", "testmsg2")
	testPostMessage(t, server.URL+"/messages/new", "admin", "broadcast")

	jsonMessages := testGetMessages(t, server.URL+"/drawer1/messages")
	if !strings.Contains(jsonMessages, "testmsg1") {
		t.Errorf("Drower1 must has %s, but does not.", "testmsg1")
	}
	if !strings.Contains(jsonMessages, "broadcast") {
		t.Errorf("Drower1 must has %s, but does not.", "broadcast")
	}
}

func testPostMessage(t *testing.T, requestURL, from, body string) {
	postValues := url.Values{
		"from": {from},
		"body": {body},
	}
	res, err := http.PostForm(requestURL, postValues)
	if err != nil {
		t.Fatalf("Error occured: %s", err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Logf("URL[%s]", requestURL)
		t.Logf("Parameter:from[%s], body[%s]", from, body)
		t.Errorf("res.StatusCode => %d, want %d", res.StatusCode, http.StatusOK)
	}
}

func testGetMessages(t *testing.T, requestURL string) string {
	res, err := http.Get(requestURL)
	if err != nil {
		t.Fatalf("Error occured: %s", err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Logf("URL[%s]", requestURL)
		t.Errorf("res.StatusCode => %d, want %d", res.StatusCode, http.StatusOK)
	}

	resMsg, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("Responce read error occured: %s", err)
	}

	return string(resMsg)
}
