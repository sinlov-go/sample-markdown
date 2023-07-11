package sample_mk

import "strings"

const (
	headerToken = '#'
	listToken   = '-'
)

var listTokens = map[rune]struct{}{
	listToken: {},
	'*':       {},
}

// Node is single markdown syntax representation
// Example: Header, ListItem, BasicItem ...
type Node interface {
	// String return string representation of node
	String() string
}

// BasicItem
// basic item node without any markdown-syntax
type BasicItem struct {
	text string
}

// NewBasicItem
// new basic item node
func NewBasicItem(text string) BasicItem {
	text = strings.TrimSpace(text)

	return BasicItem{
		text: text,
	}
}

func (i BasicItem) String() string {
	return i.text
}

// Header is markdown header node
type Header struct {
	level int
	text  string
}

// NewHeader
// new header node
//
//	level: header level will append headerToken as '#'
func NewHeader(level int, text string) Header {
	text = strings.TrimSpace(text)

	return Header{
		level: level,
		text:  text,
	}
}

// Example: # Your title
func (h Header) String() string {
	var builder strings.Builder

	for i := 0; i < h.level; i++ {
		builder.WriteString(string(headerToken))
	}

	builder.WriteString(" ")
	builder.WriteString(h.text)

	return builder.String()
}

type ListItem struct {
	text string
}

func NewListItem(text string) ListItem {
	text = strings.TrimSpace(text)

	return ListItem{
		text: text,
	}
}

func (i ListItem) String() string {
	return string(listToken) + " " + i.text
}

// Equal
// compare two nodes by strings.EqualFold
func Equal(n1, n2 Node) bool {
	return strings.EqualFold(n1.String(), n2.String())
}
