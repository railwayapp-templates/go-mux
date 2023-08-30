package responder

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// jsonPretty marshals 'v' to a byte buffer, automatically escaping HTML and setting the
// Content-Type as application/json, then writing the resulting buffer to 'w'
func jsonPretty(w http.ResponseWriter, v any, code int, indent string) error {
	buf := &bytes.Buffer{}

	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(true)
	enc.SetIndent("", indent)

	if err := enc.Encode(&v); err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	w.Write(buf.Bytes())

	return nil
}

// responds with formatted json
func JSONPretty(w http.ResponseWriter, v any, code int) error {
	return jsonPretty(w, v, code, "  ") // two spaces
}

// responds with unformatted json
func JSON(w http.ResponseWriter, v any, code int) error {
	return jsonPretty(w, v, code, "") // no spaces
}

// responds with plain text
func PlainText(w http.ResponseWriter, s string, code int) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(code)
	w.Write([]byte(s))
}
