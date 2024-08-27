package structs_methods_interfaces

import (
	"math"
	"testing"
)

func TestShapePerimeter(t *testing.T) {
	type args struct {
		shape Shape
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Rectangle a=b=1",
			args: args{Rectangle{1, 1}},
			want: 4,
		},
		{
			name: "Rectangle a=b=2",
			args: args{Rectangle{2, 2}},
			want: 8,
		},
		{
			name: "Rectangle a=2,b=3",
			args: args{Rectangle{2, 3}},
			want: 10,
		},
		{
			name: "Circle r=1",
			args: args{Circle{1}},
			want: 2 * math.Pi * 1,
		},
		{
			name: "Circle r=2",
			args: args{Circle{2}},
			want: 2 * math.Pi * 2,
		},
		{
			name: "Circle r=3.5",
			args: args{Circle{3.5}},
			want: 2 * math.Pi * 3.5,
		},
		{
			name: "Triangle a=b=c=1",
			args: args{Triangle{1, 1, 1}},
			want: 3,
		},
		{
			name: "Triangle a=b=c=2",
			args: args{Triangle{2, 2, 2}},
			want: 6,
		},
		{
			name: "Triangle a=3,b=4,c=5",
			args: args{Triangle{3, 4, 5}},
			want: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.shape.Perimeter(); got != tt.want {
				t.Errorf("Perimeter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShapeArea(t *testing.T) {
	type args struct {
		shape Shape
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Rectangle a=b=1",
			args: args{Rectangle{1, 1}},
			want: 1,
		},
		{
			name: "Rectangle a=b=2",
			args: args{Rectangle{2, 2}},
			want: 4,
		},
		{
			name: "Rectangle a=2,b=3",
			args: args{Rectangle{2, 3}},
			want: 6,
		},
		{
			name: "Circle r=1",
			args: args{Circle{1}},
			want: math.Pi * 1 * 1,
		},
		{
			name: "Circle r=2",
			args: args{Circle{2}},
			want: math.Pi * 2 * 2,
		},
		{
			name: "Circle r=3.5",
			args: args{Circle{3.5}},
			want: math.Pi * 3.5 * 3.5,
		},
		{
			name: "Triangle a=b=c=1",
			args: args{Triangle{1, 1, 1}},
			want: math.Sqrt((3.0 / 2.0) * (3.0/2.0 - 1) * (3.0/2.0 - 1) * (3.0/2.0 - 1)),
		},
		{
			name: "Triangle a=b=c=2",
			args: args{Triangle{2, 2, 2}},
			want: math.Sqrt((6 / 2) * ((6)/2 - 2) * ((6)/2 - 2) * ((6)/2 - 2)),
		},
		{
			name: "Triangle a=3,b=4,c=5",
			args: args{Triangle{3, 4, 5}},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.shape.Area(); math.Abs(got-tt.want) > 0.0001 {
				t.Errorf("Area() = %v, want %v", got, tt.want)
			}
		})
	}
}
