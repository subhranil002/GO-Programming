package response

import (
	"encoding/json"
	"net/http"
)

// JSON writes a successful JSON response with message, data, and success: true.
func JSON(w http.ResponseWriter, status int, message string, data any) {
	if data == nil {
		data = map[string]any{}
	}

	payload := map[string]any{
		"success": true,
		"message": message,
		"data":    data,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}

// Error writes a JSON error response with message, empty data, and success: false.
func Error(w http.ResponseWriter, status int, message string) {
	payload := map[string]any{
		"success": false,
		"message": message,
		"data":    map[string]any{},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}
