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

	if len(w) == 0 {
		fmt.Println("No workspaces found")
		os.Exit(0)
	}

	first := w[0]

	workspace, err := client.GetWorkspace(first.GID)
	if err != nil {
		panic(err)
	}

	fmt.Println(workspace)
}
