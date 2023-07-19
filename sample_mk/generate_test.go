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
				NewBasicItem("start"),
				NewHeader(1, "header"),
				NewListItem("item 1"),
				NewListItem("item 2"),
				NewBasicItem("end"),
			},
			want: "start\n\n# header\n\n* item 1\n\n* item 2\n\nend\n",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := GenerateText(tc.nodes)
			assert.Equal(t, tc.want, got)
		})
	}
}
