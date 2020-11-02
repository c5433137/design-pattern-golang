package builderpattern

import (
	"reflect"
	"testing"
)

func TestBuilder(t *testing.T) {
	tests := []struct {
		name string
		want *BuildStu
	}{
		// TODO: Add test cases.
		{
			want: &BuildStu{
				msg: &Message{
					Header: &Header{
						DestAddr: "127.0.0.1",
						DestPort: 8080,
						SrcAddr:  "0.0.0.0",
						SrcPort:  80,
						Items:    map[string]string{"1": "2"},
					},
					Body: &Body{
						Items: []string{"czy"},
					},
				},
			},
		},
	}

	b := Builder()
	b.WithBodyItem("czy").
		WithDestAddr("127.0.0.1").WithDestPort(8080).
		WithHeaderItem("1", "2").
		WithSrcAddr("0.0.0.0").WithSrcPort(80)
	m := b.Build()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !reflect.DeepEqual(m, tt.want.msg) {
				t.Errorf("Builder() = %v, want %v", b, tt.want)
			}
		})
	}

}
