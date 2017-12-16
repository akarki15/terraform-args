package crawler

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/akarki15/terraform-args/models"
)

var (
	dirs = []string{
		"/Users/aashishkarki/Desktop/terraform-provider-aws-master/website/docs/d",
		//"testfiles/d",
		// "testfiles/r",
	}
)

func parseDir(dirpath string) (models.Commands, error) {
	dir, err := os.Open(dirpath)
	if err != nil {
		return nil, err
	}
	files, err := dir.Readdir(0)
	if err != nil {
		return nil, err
	}

	commands := models.Commands{}
	for _, file := range files {
		filename := filepath.Join(dirpath, file.Name())
		doc := parseFile(filename)
		command, err := parseCommand(doc)
		if err != nil {
			return nil, err
		}
		commands = append(commands, command)
	}
	return commands, nil
}

func parseFile(filename string) string {
	file, err := os.Open(filename)
	defer func() { _ = file.Close() }()
	panicIfErr(err)
	content, err := ioutil.ReadAll(file)
	panicIfErr(err)
	return string(content)
}

func panicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}

func parseCommand(content string) (models.Command, error) {
	splits := strings.Split(content, "page_title:")
	if len(splits) <= 1 {
		return models.Command{}, errors.New("page title : splits isn't big enough!")
	}
	secondSplit := strings.Split(splits[1], "\n")
	if len(splits) == 0 {
		return models.Command{}, errors.New("first split should be title!")
	}
	strCommand := string(secondSplit[0])
	strCommand = strCommand[5 : len(strCommand)-1] // more madness
	description := parseCommandDescription(content)
	args := parseDArguments(content)

	return models.Command{
		Name:        strCommand,
		Description: description,
		ArgsStr:     args,
	}, nil
}

func parseCommandDescription(content string) string {
	splits := strings.Split(content, "description: |-")
	if len(splits) <= 1 {
		panic(fmt.Sprintf("description: splits isn't big enough! : %v", content))
	}
	secondSplit := strings.Split(splits[1], "\n")
	if len(secondSplit) == 0 {
		panic("first split should be title!")
	}
	return string(secondSplit[1])
}

func itemizeArguments(arguments string) map[string]string {
	return nil

}

func parseDArguments(content string) string {
	splits := strings.Split(content, "## Argument Reference")
	if len(splits) == 1 {
		return ""
	}
	//fmt.Printf("%+v", splits)
	secondSplits := strings.Split(splits[1], "## Attributes Reference")
	if len(secondSplits) == 0 {
		panic("first split should be title!")
	}
	return secondSplits[0]
}
