package helpers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"moviedl/utils"
)

func Response(w http.ResponseWriter, data utils.M, statusCode int, headers map[string]interface{}) {
	res, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if headers == nil {
		headers = utils.M{"Content-Type": "application/json"}
	}
	for k, v := range headers {
		w.Header().Set(k, fmt.Sprintf("%v", v))
	}
	w.Write([]byte(res))
}
