package main

import (
	"fmt"
	"os"
	"time"

	"github.com/fsmiamoto/asana"
)

func main() {
	url := os.Getenv("ASANA_API_URL")
	token := os.Getenv("ASANA_API_TOKEN")

	client := asana.NewClient(
		asana.WithBaseURL(url),
		asana.WithTimeout(5*time.Second),
		asana.WithToken(token),
	)

	w, err := client.GetWorkspaces()
	if err != nil {
		panic(err)
	}

	fmt.Println(w)
}
