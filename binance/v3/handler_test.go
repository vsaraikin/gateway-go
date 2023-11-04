package v3

import (
	"encoding/json"
	"gateaway/binance/models"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetExchangeInfo(t *testing.T) {
	// Create a mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Define the expected JSON response from the server
		response := models.ExchangeInfo{
			// Populate with expected fields...
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	// Create a new instance of BinanceClient with the mock server URL
	client := NewBinanceClient("api-key", "secret-key")
	client.BaseURL = server.URL // override the BaseURL with our mock server URL

	// Call the GetExchangeInfo method
	_, err := client.GetExchangeInfo()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Here you would check the contents of info to make sure it's what you expect.
	// Since the structure of models.ExchangeInfo isn't known, you'll need to fill this in.
	// For example:
	// if info.SomeField != "expected value" {
	//     t.Errorf("Expected SomeField to be 'expected value', got '%v'", info.SomeField)
	// }

	// Note: Replace the above if statement with actual checks on the fields of your models.ExchangeInfo struct.
}
