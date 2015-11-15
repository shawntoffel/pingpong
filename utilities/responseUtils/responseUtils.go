package responseUtils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Response(w http.ResponseWriter, content []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	fmt.Fprintf(w, "%s", content)
}

func JsonResponse(w http.ResponseWriter, response interface{}) {
	jsonContent, _ := json.Marshal(&response)

	Response(w, jsonContent)
}
