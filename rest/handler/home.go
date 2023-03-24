package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type URL string
type Documentation struct {
	Url         URL    `json:"url"`
	Method      string `json:"method"`
	Description string `json:"desccription"`
}

func (u URL) MarshalText() ([]byte, error) {
	return []byte(fmt.Sprintf("http://localhost:3000%s", u)), nil
}

func Home(w http.ResponseWriter, r *http.Request) {
	documentation := []Documentation{
		{
			Url:         URL("/"),
			Method:      "GET",
			Description: "Rest API View",
		},
		{
			Url:         URL("/block"),
			Method:      "GET",
			Description: "block API View",
		},
		{
			Url:         URL("/block"),
			Method:      "POST",
			Description: "Post API View",
		},
		{
			Url:         URL("/block/{id}"),
			Method:      "GET",
			Description: "GET Block Retreive View",
		},
	}
	json.NewEncoder(w).Encode(documentation)
}
