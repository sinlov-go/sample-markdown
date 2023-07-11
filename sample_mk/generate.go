package sample_mk

import "strings"

// GenerateText
// return single string which represents all markdown nodes
func GenerateText(nodes []Node) string {
	lines := make([]string, len(nodes))
	for i, node := range nodes {
		lines[i] = node.String()
	}

	result := strings.Join(lines, "\n\n")
	// Ensure there is newline at end of file
	result += "\n"

	return result
}
