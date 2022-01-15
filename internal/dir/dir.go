package dir

import (
	"fmt"
	"io/fs"
	"os"
	"text/template"
)

//Create a directory
func Create(name string, perm fs.FileMode) error {
	err := os.MkdirAll(name, perm)
	if err != nil {
		return fmt.Errorf("dir_create_mkdirall %w", err)
	}

	return nil
}

//CreateFile new file in a specific directory with template variables
func CreateFile(name, directory, extension string, vars map[string]interface{}) error {
	tmpl, err := template.ParseFiles(fmt.Sprintf("templates/%s.tmpl", name))
	if err != nil {
		return fmt.Errorf("dir_createfile_parsefiles %w", err)
	}

	file, err := os.Create(fmt.Sprintf("%s/%s%s", directory, name, extension))
	if err != nil {
		return fmt.Errorf("dir_createfile_create %w", err)
	}

	defer file.Close()

	err = tmpl.Execute(file, vars)
	if err != nil {
		return fmt.Errorf("dir_createfile_execute %w", err)
	}

	return nil
}
