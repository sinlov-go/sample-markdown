package sample_mk

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name  string
		lines []string
		want  []Node
	}{
		{
			name: "level 1",
			lines: []string{
				"# abc",
				"- xyz",
			},
			want: []Node{
				NewHeader(1, "abc"),
				NewListItem("xyz"),
			},
		},
		{
			name: "level 3 with alternative",
			lines: []string{
				"### xyz",
				"* abc",
			},
			want: []Node{
				NewHeader(3, "xyz"),
				NewListItem("abc"),
			},
		},
		{
			name:  "nil",
			lines: []string{},
			want:  nil,
		},
		{
			name: "line empty",
			lines: []string{
				"# abc",
				"",
				"- xyz",
				"",
			},
			want: []Node{
				NewHeader(1, "abc"),
				NewListItem("xyz"),
			},
		},
		{
			name: "line full type",
			lines: []string{
				"# abc",
				"",
				"info",
				"",
				"## Features",
				"",
				"- xyz",
				"- foo",
				"",
				"## build",
				"",
				"- one",
				"- two",
				"",
			},
			want: []Node{
				NewHeader(1, "abc"),
				NewBasicItem("info"),
				NewHeader(2, "Features"),
				NewListItem("xyz"),
				NewListItem("foo"),
				NewHeader(2, "build"),
				NewListItem("one"),
				NewListItem("two"),
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := Parse(tc.lines)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestParseHeader(t *testing.T) {
	tests := []struct {
		name string
		line string
		want Header
	}{
		{
			name: "level 1",
			line: "# abc",
			want: NewHeader(1, "abc"),
		},
		{
			name: "level 3",
			line: "### xyz",
			want: NewHeader(3, "xyz"),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := parseHeader(tc.line)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestParseListItem(t *testing.T) {
	tests := []struct {
		name string
		line string
		want ListItem
	}{
		{
			name: "normal",
			line: "- abc",
			want: ListItem{
				nodeType: NodeTypeListItem,
				text:     "abc",
			},
		},
		{
			name: "alternative",
			line: "* xyz",
			want: ListItem{
				nodeType: NodeTypeListItem,
				text:     "xyz",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := parseListItem(tc.line)
			assert.Equal(t, tc.want, got)
		})
	}
}
