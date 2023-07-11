package example

import (
	"github.com/sebdah/goldie/v2"
	"github.com/sinlov-go/sample-markdown/sample_mk"
	"testing"
)

func TestGenerateText(t *testing.T) {
	// mock GenerateText
	tests := []struct {
		name    string
		c       []sample_mk.Node
		wantErr error
	}{
		{
			name: "sample", // TODO: testData/TestGenerateText/sample.golden
			c: []sample_mk.Node{
				sample_mk.NewHeader(2, "v1.0.0 (2020-01-18)"),
				sample_mk.NewHeader(3, "Features"),
				sample_mk.NewListItem("feat: new feature"),
				sample_mk.NewHeader(3, "Bug Fixes"),
				sample_mk.NewListItem("fix: new fix"),
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			g := goldie.New(t,
				goldie.WithDiffEngine(goldie.ClassicDiff),
			)

			// do GenerateText
			gotResult := sample_mk.GenerateText(tc.c)

			// verify GenerateText
			g.Assert(t, t.Name(), []byte(gotResult))
		})
	}
}
