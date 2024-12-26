package txtest_test

import (
	"net/http"
	"testing"
)

func TestInsert(t *testing.T) {
	res, err := http.Get("http://localhost:8080/insert")
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Logf("expected status code %d, got %d", http.StatusOK, res.StatusCode)
	}

	t.Log(res.Body)
}
