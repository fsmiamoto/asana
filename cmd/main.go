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

	workspaces, err := client.GetWorkspaces()
	if err != nil {
		panic(err)
	}

	if len(workspaces) == 0 {
		fmt.Println("No workspaces found")
		os.Exit(0)
	}

	first := workspaces[1]

	w, err := client.GetWorkspace(first.GID)
	if err != nil {
		panic(err)
	}

	projects, err := client.GetWorkspaceProjects(w.GID)
	if err != nil {
		panic(err)
	}

	tasks, err := client.GetTasksFromProject(projects[0].GID)
	if err != nil {
		panic(err)
	}

	fmt.Println(tasks)
}
