package sample_mk

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNodeType(t *testing.T) {
	// mock NodeType
	tests := []struct {
		name    string
		node    Node
		wantRes int
	}{
		{
			name:    "NewBasicItem",
			node:    NewBasicItem("foo"),
			wantRes: NodeTypeBasic,
		},
		{
			name:    "NewHeader",
			node:    NewHeader(2, "foo"),
			wantRes: NodeTypeHeader,
		},
		{
			name:    "NewListItem",
			node:    NewListItem("bar"),
			wantRes: NodeTypeListItem,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			// do NodeType

			// verify NodeType
			assert.Equal(t, tc.wantRes, tc.node.Type())
		})
	}
}

func TestBasicItem(t *testing.T) {
	// mock BasicItem
	tests := []struct {
		name    string
		wantRes BasicItem
		c       string
		wantErr error
	}{
		{
			name:    "sample",
			c:       "abc",
			wantRes: NewBasicItem("abc"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			// do BasicItem
			gotResult := NewBasicItem(tc.c)
			// verify BasicItem
			assert.Equal(t, tc.wantRes, gotResult)
		})
	}
}

func TestHeaderString(t *testing.T) {
	tests := []struct {
		name      string
		header    Header
		want      string
		wantLevel int
	}{
		{
			name:      "level 1",
			header:    NewHeader(1, "abc"),
			want:      "# abc",
			wantLevel: 1,
		},
		{
			name:      "level 3",
			header:    NewHeader(3, "xyz"),
			want:      "### xyz",
			wantLevel: 3,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.header.String()
			assert.Equal(t, tc.want, got)
			assert.Equal(t, tc.wantLevel, tc.header.Level())
		})
	}
}

func TestListItemString(t *testing.T) {
	tests := []struct {
		name     string
		listItem ListItem
		want     string
	}{
		{
			name: "normal",
			listItem: ListItem{
				text: "abc",
			},
			want: "- abc",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.listItem.String()
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestEqual(t *testing.T) {
	tests := []struct {
		name string
		n1   Node
		n2   Node
		want bool
	}{
		{
			name: "header",
			n1:   NewHeader(0, "CHANGELOG"),
			n2:   NewHeader(0, "Changelog"),
			want: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := Equal(tc.n1, tc.n2)
			assert.Equal(t, tc.want, got)
		})
	}
}
