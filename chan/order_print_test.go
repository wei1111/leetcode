package _chan

import "testing"

func TestOrderPrint(t *testing.T) {
	type args struct {
		n int
		m int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "TestOrderPrint",
			args: args{
				10, 100,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			OrderPrint(tt.args.n, tt.args.m)
		})
	}
}
