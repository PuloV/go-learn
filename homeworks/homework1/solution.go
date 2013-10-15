package main

import (
	"regexp"
)

// Directory Separator
// "/"		for Linux/OSX/FreeBSD/etc.
// "\\"		should work for Windows, not tested.
const DS = "/"

// Prepends or/and appends missing Directory Separators
func normilizePath(path string) string {
	if len(path) == 0 || path[0] != DS[0] {
		path = string(DS[0]) + path
	}

	if path[len(path)-1] != DS[0] {
		path = path + string(DS[0])
	}

	return path
}

// Parses the path
func parsePath(path string) string {
	path = normilizePath(path)

	// Remove single dots as they interrupt the next
	// replacement.
	regex, err := regexp.Compile(DS + `\.` + DS)
	if err != nil {
		return "Zero regex won't compile, FIX IT!"
	}
	for regex.MatchString(path) == true {
		path = regex.ReplaceAllLiteralString(path, DS)
	}

	// Removing /somedir/../ pairs.
	// It will parse multiline paths as well.
	// p.s. some black magic involved.
	regex, err = regexp.Compile(`(?sU)` + DS + `\.*?[^` + DS + `\.]+\.*[^` + DS + `]*` + DS + `\.{2}` + DS)
	if err != nil {
		return "First regex won't compile, FIX IT!"
	}
	for regex.MatchString(path) == true {
		path = regex.ReplaceAllLiteralString(path, DS)
	}

	// Cleaning up remaining dots
	regex, err = regexp.Compile(DS + `\.{1,2}` + DS)
	if err != nil {
		return "Second regex won't compile, FIX IT!"
	}
	for regex.MatchString(path) == true {
		path = regex.ReplaceAllLiteralString(path, DS)
	}

	return path
}
