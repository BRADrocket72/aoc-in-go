package main

import (
	"image"
	"testing"
)

func TestCalcArea(t *testing.T) {
	type args struct {
		pointA image.Point
		pointB image.Point
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Zero Area Test (Same Points)",
			args: args{
				pointA: image.Point{X: 0, Y: 0},
				pointB: image.Point{X: 0, Y: 0},
			},
			want: 1.0,
		},
		{
			name: "24 Area Test ",
			args: args{
				pointA: image.Point{X: 2, Y: 5},
				pointB: image.Point{X: 9, Y: 7},
			},
			want: 24.0,
		},
		{
			name: "Zero Area Test (Same Points)",
			args: args{
				pointA: image.Point{X: 7, Y: 1},
				pointB: image.Point{X: 11, Y: 7},
			},
			want: 35.0,
		},
		{
			name: "Zero Area Test (Same Points)",
			args: args{
				pointA: image.Point{X: 2, Y: 5},
				pointB: image.Point{X: 11, Y: 1},
			},
			want: 50.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalcArea(tt.args.pointA, tt.args.pointB); got != tt.want {
				t.Errorf("CalcArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
