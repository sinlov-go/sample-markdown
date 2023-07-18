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
				Header{
					level: 1,
					text:  "abc",
				},
				ListItem{
					text: "xyz",
				},
			},
		},
		{
			name: "level 3 with alternative",
			lines: []string{
				"### xyz",
				"* abc",
			},
			want: []Node{
				Header{
					level: 3,
					text:  "xyz",
				},
				ListItem{
					text: "abc",
				},
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
				Header{
					level: 1,
					text:  "abc",
				},
				ListItem{
					text: "xyz",
				},
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
				Header{
					level: 1,
					text:  "abc",
				},
				BasicItem{
					text: "info",
				},
				Header{
					level: 2,
					text:  "Features",
				},
				ListItem{
					text: "xyz",
				},
				ListItem{
					text: "foo",
				},
				Header{
					level: 2,
					text:  "build",
				},
				ListItem{
					text: "one",
				},
				ListItem{
					text: "two",
				},
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
			want: Header{
				level: 1,
				text:  "abc",
			},
		},
		{
			name: "level 3",
			line: "### xyz",
			want: Header{
				level: 3,
				text:  "xyz",
			},
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
				text: "abc",
			},
		},
		{
			name: "alternative",
			line: "* xyz",
			want: ListItem{
				text: "xyz",
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
