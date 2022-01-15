package communication

import (
	"bufio"
	"fmt"
	"os"
)

type templateVariables struct {
	BaseDirectory string
	Project       string
	Port          string
	Author        string
	Module        string
}

func GetTemplateVariables() (templateVariables, error) {
	var tv templateVariables

	fmt.Println("Hello there")

	fmt.Println("which directory you want to use as base")
	var baseDirectory string
	_, err := fmt.Scanln(&baseDirectory)
	if err != nil {
		return templateVariables{}, fmt.Errorf("communication_gettemplatevariables_base_directory %w", err)
	}
	tv.BaseDirectory = baseDirectory

	fmt.Println("type the project name directory")
	var project string
	_, err = fmt.Scanln(&project)
	if err != nil {
		return templateVariables{}, fmt.Errorf("communication_gettemplatevariables_project %w", err)
	}
	tv.Project = project

	fmt.Println("what port you want to use")
	var port string
	_, err = fmt.Scanln(&port)
	if err != nil {
		return templateVariables{}, fmt.Errorf("communication_gettemplatevariables_port %w", err)
	}
	tv.Port = port

	fmt.Println("author name?")
	reader := bufio.NewReader(os.Stdin)
	author, err := reader.ReadString('\n')
	if err != nil {
		return templateVariables{}, fmt.Errorf("communication_gettemplatevariables_authorname %w", err)
	}
	tv.Author = author

	fmt.Println("whats the module name")
	var module string
	_, err = fmt.Scanln(&module)
	if err != nil {
		return templateVariables{}, fmt.Errorf("communication_gettemplatevariables_module %w", err)
	}
	tv.Module = module

	return tv, nil
}
