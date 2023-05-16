package sort

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	type args struct {
		input []int64
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		{
			name: "test quick sort1",
			args: args{input: []int64{3, 2, 1, 2, 3,}},
			want: []int64{1, 2, 2,3,3,},
		},
		{
			name: "test quick sort1",
			args: args{input: []int64{3, 2, 1, 1, 1,}},
			want: []int64{1, 1, 1,2,3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := QuickSort(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QuickSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
