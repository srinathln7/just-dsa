package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Node represents a node in the N-ary tree
type Node struct {
	Val      int
	Children []*Node
}

// Serialize encodes the N-ary tree into a string
func Serialize(root *Node) string {
	if root == nil {
		return ""
	}

	var sb strings.Builder
	sb.WriteString(strconv.Itoa(root.Val))
	sb.WriteString("(")

	for i, child := range root.Children {
		sb.WriteString(Serialize(child))
		if i != len(root.Children)-1 {
			sb.WriteString(",")
		}
	}

	sb.WriteString(")")
	return sb.String()
}

// Deserialize decodes the string into an N-ary tree
func Deserialize(data string) *Node {
	if len(data) == 0 {
		return nil
	}

	var i int
	val, i := parseValue(data, i)
	root := &Node{Val: val}

	i++ // skip '('
	children := []string{}
	for i < len(data) {
		if data[i] == ',' {
			i++
		} else if data[i] == ')' {
			i++
			break
		} else {
			child, j := parseNode(data, i)
			children = append(children, child)
			i = j + 1
		}
	}

	for _, child := range children {
		root.Children = append(root.Children, Deserialize(child))
	}

	return root
}

func parseValue(data string, start int) (int, int) {
	end := start
	for end < len(data) && data[end] != '(' {
		end++
	}
	val, _ := strconv.Atoi(data[start:end])
	return val, end
}

func parseNode(data string, start int) (string, int) {
	stack := 0
	end := start
	for end < len(data) {
		if data[end] == '(' {
			stack++
		} else if data[end] == ')' {
			stack--
		}
		if stack == 0 {
			break
		}
		end++
	}
	return data[start:end], end
}

func main() {
	root := &Node{
		Val: 1,
		Children: []*Node{
			{Val: 3, Children: []*Node{{Val: 5}, {Val: 6}}},
			{Val: 2},
			{Val: 4},
		},
	}

	serialized := Serialize(root)
	fmt.Println("Serialized tree:", serialized)

	deserialized := Deserialize(serialized)
	fmt.Println("Deserialized tree:", Serialize(deserialized))
}
