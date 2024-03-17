package main

import (
	"fmt"

	"github.com/charmbracelet/huh"
)

func main() {
	var project_name string
	var project_template string

	huh.NewInput().
		Title("Project name").
		Value(&project_name).
		Run()

	huh.NewSelect[string]().
		Title("Project template").
		Options(
			huh.NewOption("Python", "python"),
			huh.NewOption("Zig", "zig"),
		).
		Value(&project_template).
		Run()

	fmt.Printf("Project name: %s!\n", project_name)
	fmt.Printf("Project template: %s\n", project_template)

}
