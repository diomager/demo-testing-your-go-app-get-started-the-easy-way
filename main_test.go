package main

import "testing"

func TestAvg(t *testing.T) {
	type args struct {
		nos []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				nos: []int{2, 2},
			},
			want: 2,
		},
		{
			args: args{
				nos: []int{-1, 0, 1},
			},
			want: 0,
		},
		{
			args: args{
				nos: []int{-4, -4, -4, -4},
			},
			want: -4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Avg(tt.args.nos...); got != tt.want {
				t.Errorf("Avg() = %v, want %v", got, tt.want)
			}
		})
	}
}
