package main

import "strings"

func simplifyPath(path string) string {

	// Key Idea: Use a stack. Split the string using a the delimiter root "/"

	root := "/"
	cmds := strings.Split(path, root)
	idx := 0
	for _, cmd := range cmds {
		if len(cmd) > 0 {
			cmds[idx] = cmd
			idx++
		}
	}
	cmds = cmds[:idx]

	var stack []string
	for _, cmd := range cmds {
		n := len(stack)
		switch cmd {
		case ".":
			continue
		case "..":
			if n > 0 {
				stack = stack[:n-1]
			}
		default:
			stack = append(stack, cmd)
		}
	}

	if len(stack) == 0 {
		return root
	}

	result := root + strings.Join(stack, "/")
	return result
}
