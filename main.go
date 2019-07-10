package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

var verbose *bool
var caseSensitive *bool

type variableList struct {
	path string
	vars []string
}

func (l variableList) contains(name string) bool {
	for _, v := range l.vars {
		if *caseSensitive {
			if v == name {
				return true
			}
		} else {
			if strings.ToUpper(v) == strings.ToUpper(name) {
				return true
			}
		}
	}

	return false
}

func main() {
	verbose = flag.Bool("v", false, "Print more verbose output")
	caseSensitive = flag.Bool("cs", false, "Check should be case sensitive")

	flag.Parse()

	if len(flag.Args()) < 2 {
		flag.Usage()
		os.Exit(0)
	}

	lists := []variableList{}
	for _, arg := range flag.Args() {
		list, err := parseFile(arg)
		if err != nil {
			exit(err.Error())
		}

		lists = append(lists, *list)
	}

	for _, listA := range lists {
		for _, listB := range lists {
			if listA.path == listB.path {
				continue
			}

			compare(listA, listB)
		}
	}
}

func compare(list1 variableList, list2 variableList) {

	p1 := list1.path
	if !*verbose {
		p1 = path.Base(p1)
	}

	p2 := list2.path
	if !*verbose {
		p2 = path.Base(p2)
	}

	for _, v := range list1.vars {
		if !list2.contains(v) {
			fmt.Printf("%v from %v missing in %v\n", v, p1, p2)
		}
	}
}

func parseFile(path string) (*variableList, error) {
	contentBytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	if string(contentBytes[0]) == "{" {
		vars, err := parseJSONFile(contentBytes)
		if err != nil {
			return nil, err
		}

		return &variableList{
			path: path,
			vars: vars,
		}, nil
	}

	return &variableList{
		path: path,
		vars: parseEnvFile(contentBytes),
	}, nil
}

func parseJSONFile(content []byte) ([]string, error) {
	var target map[string]interface{}
	if err := json.Unmarshal(content, &target); err != nil {
		return nil, err
	}

	names := []string{}
	for k := range target {
		names = append(names, k)
	}

	return names, nil
}

func parseEnvFile(content []byte) []string {
	lines := strings.Split(string(content), "\n")
	names := []string{}

	for _, line := range lines {
		line = strings.Trim(line, " ")

		if len(line) == 0 {
			continue
		}

		if line[0:1] == "#" {
			continue
		}

		parts := strings.Split(line, "=")
		names = append(names, parts[0])
	}

	return names
}

func exit(message string) {
	fmt.Println(message)
	os.Exit(1)
}
