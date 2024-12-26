package txtest_test

import (
	"net/http"
	"testing"
)

func TestRootRequest(t *testing.T) {
	res, err := http.Get("http://localhost:8080")

	if err != nil {
		t.Fatal(err)
	}

	t.Log(res)
}
