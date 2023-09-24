package main

import "strings"

func isValidRequest(req Request) (bool, string) {
	if len(req.Name) == 0 {
		return false, "Name cannot be empty"
	}

	if len(req.Email) == 0 {
		return false, "Email cannot be empty"
	}

	if !strings.Contains(req.Email, "@") {
		return false, "Email must have an @ character"
	}

	return true, ""
}
