package main

import (
	"fmt"
	"time"

	"github.com/junkd0g/go-project-generator/internal/communication"
	"github.com/junkd0g/go-project-generator/internal/dir"
)

func main() {

	templateVariables, err := communication.GetTemplateVariables()
	if err != nil {
		panic(err)
	}

	projectFullDirectory := fmt.Sprintf("%s/%s", templateVariables.BaseDirectory, templateVariables.Project)
	err = dir.Create(projectFullDirectory, 0755)
	if err != nil {
		panic(err)
	}

	assets := fmt.Sprintf("%s/%s", projectFullDirectory, "assets")
	err = dir.Create(assets, 0755)
	if err != nil {
		panic(err)
	}

	config := fmt.Sprintf("%s/%s", assets, "config")
	err = dir.Create(config, 0755)
	if err != nil {
		panic(err)
	}

	internal := fmt.Sprintf("%s/%s", projectFullDirectory, "internal")
	err = dir.Create(internal, 0755)
	if err != nil {
		panic(err)
	}

	controller := fmt.Sprintf("%s/%s", internal, "controller")
	err = dir.Create(controller, 0755)
	if err != nil {
		panic(err)
	}

	configInternal := fmt.Sprintf("%s/%s", internal, "config")
	err = dir.Create(configInternal, 0755)
	if err != nil {
		panic(err)
	}

	//--- documentation
	documentation := fmt.Sprintf("%s/%s", projectFullDirectory, "documentation")
	dir.Create(documentation, 0755)

	//create Dockerfile
	vars := make(map[string]interface{})
	dir.CreateFile(
		"Dockerfile",
		projectFullDirectory,
		"",
		vars,
	)

	//create app.go file
	vars = make(map[string]interface{})
	vars["Date"] = time.Now().Format("01/02/2006")
	vars["Author"] = templateVariables.Author
	vars["Port"] = templateVariables.Port
	vars["Module"] = templateVariables.Module
	dir.CreateFile(
		"app",
		projectFullDirectory,
		".go",
		vars,
	)

	//create Makefile file
	dir.CreateFile(
		"Makefile",
		projectFullDirectory,
		"",
		vars,
	)

	//create configs files
	dir.CreateFile(
		"production",
		config,
		".yaml",
		vars,
	)

	dir.CreateFile(
		"staging",
		config,
		".yaml",
		vars,
	)

	dir.CreateFile(
		"dev",
		config,
		".yaml",
		vars,
	)

	dir.CreateFile(
		"docker",
		config,
		".yaml",
		vars,
	)
}
