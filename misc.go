package Skapt

import (
	"fmt"
	"html/template"
	"os"
	"strconv"
	"strings"
)

// getPathVersion searches the current directory
// for the VERSION file if it's included return the
// exact path of it. If the VERSION file is not included
// this will return empty string
func getPathVersion() string {
	// our file name
	name := "VERSION"

	rootPath, err := os.Getwd()

	if err != nil {
		fmt.Println("Can't optain the root path of the project")
		return ""
	}

	// host our root/base path and our fileName(VERSION)
	holder := []string{
		rootPath,
		name,
	}

	path := strings.Join(holder, "/")

	return path
}

// getContentVersion this will extract the content Version
// from the VERSION file and parse it
func getContentVersion() string {
	// local var that stores the content
	cnt := make([]byte, 10)

	file := openFileVersion(filePath)

	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	n, err := file.Read(cnt)

	if err != nil && n == 0 {
		fmt.Println("Can't read VERSION file")
	}

	return string(cnt)
}

func openFileVersion(path string) *os.File {

	file, err := os.OpenFile(path, os.O_RDONLY, 0666)

	if err != nil {
		fmt.Println("Can't open file VERSION")
	}

	return file

}
func getTarget(opt Option, args []string) (string, int) {

	var (
		lenArgs = len(args)
		i       int
		targetS string
		targetI int
		err     error
	)

	// for every argument passed
	for i = 0; i < lenArgs; i++ {
		// if the arg is equal with the name or
		// if the arg is equal with the alas
		if args[i] == opt.alias || args[i] == opt.name {
			if opt.typeFlag == STRING {
				targetS = args[i+1]
			} else {
				if opt.typeFlag == INT {
					targetI, err = atoiWrapper(args[i+1])
				}
			}
			break
		} //if
	}

	//if err had occured
	// handle it
	if err != nil {
		errOnExit(err)
	}

	return targetS, targetI
}

// simple wrapper for strConv
// providing custom error output
func atoiWrapper(value string) (int, error) {
	val, err := strconv.Atoi(value)

	//if error
	if err != nil {
		return val, errTINT
	}

	return val, nil
}

// transfrom string to Flag INT, STRING, BOOL
func stringToFlag(value string) uint8 {
	switch value {
	case "STRING":
		return STRING
	case "BOOL":
		return BOOL
	case "INT":
		return INT
	}

	return UNKNOWN
}

const (
	appFlagHelpTemplate = `

NAME:	{{ .Name }}

USAGE:
	{{ .Usage }}

DESCRIPTION:
	{{ .Description }}

OPTIONS:
{{range $opt := .Options }}
	{{ $opt.Name }}, {{ $opt.Alias }}
		{{ $opt.Description }}
{{ end }}

	--help, -h  print out the help message

AUTHORS :
	{{ range $auth := .Authors }} {{ $auth }} {{ end }}
VERSION:
	{{ .Version }}
`
	appCommandHelpTemplate = `
NAME :	{{ .Name }}

USAGE:
	{{ .Usage }}

DESCRIPTION:
	{{ .Description }}

COMMANDS:
{{ range $cmd := .Commands }}
	{{ $cmd.Name }} - {{ $cmd.Usage }}
		{{ $cmd.Description }}
		{{ range $opt := $cmd.Options }}
		{{ $opt.Name }}, {{ $opt.Alias }}
		{{ $opt.Description }}
		{{ end }}
{{ end }}

	--help, -h  print out the help message

VERSION:
	{{ .Version }}

AUTHORS:
	{{ range $auth := .Authors }} {{ $auth }} {{ end }}
`
)

// Basic simple help generation template
// filling the template with all the info dynamically when
// the App struct is filled
func getHelpTemplate(temp string) *template.Template {
	tmpl, err := template.New("help").Parse(temp)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return tmpl
}

// parsing and echo it to the STDOUT the template
func (a App) echoHelp() {
	tmp := appFlagHelpTemplate

	// command base app
	if a.options == nil {
		tmp = appCommandHelpTemplate
	}

	if tmpl := getHelpTemplate(tmp); tmpl != nil {
		if err := tmpl.Execute(os.Stdout, a); err != nil {
			fmt.Println(err.Error())
		}
	}
}
