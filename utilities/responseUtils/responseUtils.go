package responseUtils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Writes a 200 success response containing content
func Response(w http.ResponseWriter, content []byte) {
	w.WriteHeader(200)

	fmt.Fprintf(w, "%s", content)
}

// Writes json response with appropriate headers
func JsonResponse(w http.ResponseWriter, response interface{}) {
	w.Header().Set("Content-Type", "application/json")
	jsonContent, _ := json.Marshal(&response)

	Response(w, jsonContent)
}
