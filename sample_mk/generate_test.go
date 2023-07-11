package sample_mk

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerate(t *testing.T) {
	tests := []struct {
		name  string
		nodes []Node
		want  string
	}{
		{
			name: "normal",
			nodes: []Node{
				BasicItem{
					text: "start",
				},
				Header{
					level: 1,
					text:  "header",
				},
				ListItem{
					text: "item 1",
				},
				ListItem{
					text: "item 2",
				},
				BasicItem{
					text: "end",
				},
			},
			want: "start\n\n# header\n\n- item 1\n\n- item 2\n\nend\n",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := GenerateText(tc.nodes)
			assert.Equal(t, tc.want, got)
		})
	}
}
