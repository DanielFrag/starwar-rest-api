package utils

import (
	"testing"
)

func TestRequestUtils(t *testing.T) {
	requestUtils := RequestUtils{}
	t.Run("ValidRequest", func(t *testing.T) {
		result, resultError := requestUtils.PerformRequest("https://www.google.com", "GET", nil, nil)
		if resultError != nil {
			t.Error("Request error: " + resultError.Error())
		}
		if result == nil || len(result) == 0 {
			t.Error("Body with no length")
		}
	})
	t.Run("InvalidProtocol", func(t *testing.T) {
		_, resultError := requestUtils.PerformRequest("www.google.com", "GET", nil, nil)
		if resultError == nil {
			t.Error("Should return an invalid protocol error")
		}
	})
	t.Run("InvalidUrl", func(t *testing.T) {
		_, resultError := requestUtils.PerformRequest("https://www.a.com", "GET", nil, nil)
		if resultError == nil {
			t.Error("Should return no host")
		}
	})
	t.Run("InvalidResource", func(t *testing.T) {
		_, resultError := requestUtils.PerformRequest("https://www.google.com/nada", "GET", nil, nil)
		if resultError == nil {
			t.Error("Sould return a not found error")
		}
	})
}
