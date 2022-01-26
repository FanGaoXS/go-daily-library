package number

import (
	"math"
	"reflect"
	"testing"
)

func TestMax(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{a: 1, b: 5}, want: 5},
		{args: args{a: 0, b: 5}, want: 5},
		{args: args{a: 5, b: 5}, want: 5},
		{args: args{a: math.MaxInt, b: math.MinInt}, want: math.MaxInt},
		{args: args{a: math.MaxInt64, b: math.MinInt64}, want: math.MaxInt64},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Max(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Max() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMin(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{a: 1, b: 5}, want: 1},
		{args: args{a: 0, b: 5}, want: 0},
		{args: args{a: 5, b: 5}, want: 5},
		{args: args{a: math.MaxInt, b: math.MinInt}, want: math.MinInt},
		{args: args{a: math.MaxInt64, b: math.MinInt64}, want: math.MinInt64},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Min(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Min() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxFromMultiple(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{nums: []int{1, 7, 5, 8, 9}}, want: 9},
		{args: args{nums: []int{1, 7, math.MaxInt, 8, 9}}, want: math.MaxInt},
		{args: args{nums: []int{1, 7, math.MinInt, 8, 9}}, want: 9},
		{args: args{nums: []int{1, 7, math.MaxInt64, 8, 9}}, want: math.MaxInt64},
		{args: args{nums: []int{1, 7, math.MinInt64, 8, 9}}, want: 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxFromMultiple(tt.args.nums...); got != tt.want {
				t.Errorf("MaxFromMultiple() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinFromMultiple(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{nums: []int{1, 7, 5, 8, 9}}, want: 1},
		{args: args{nums: []int{1, 7, math.MaxInt, 8, 9}}, want: 1},
		{args: args{nums: []int{1, 7, math.MinInt, 8, 9}}, want: math.MinInt},
		{args: args{nums: []int{1, 7, math.MaxInt64, 8, 9}}, want: 1},
		{args: args{nums: []int{1, 7, math.MinInt64, 8, 9}}, want: math.MinInt64},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinFromMultiple(tt.args.nums...); got != tt.want {
				t.Errorf("MinFromMultiple() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{args: args{nums: []int{1, 7, 5, 8, 9}}, want: []int{1, 5, 7, 8, 9}},
		{args: args{nums: []int{1, 7, math.MaxInt, 8, 9}}, want: []int{1, 7, 8, 9, math.MaxInt}},
		{args: args{nums: []int{1, 7, math.MinInt, 8, 9}}, want: []int{math.MinInt, 1, 7, 8, 9}},
		{args: args{nums: []int{1, 7, math.MaxInt64, 8, 9}}, want: []int{1, 7, 8, 9, math.MaxInt64}},
		{args: args{nums: []int{1, 7, math.MinInt64, 8, 9}}, want: []int{math.MinInt64, 1, 7, 8, 9}},
		{args: args{nums: []int{1}}, want: []int{1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sort(tt.args.nums...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bubbleSort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{args: args{nums: []int{1, 7, 5, 8, 9}}, want: []int{1, 5, 7, 8, 9}},
		{args: args{nums: []int{1, 7, math.MaxInt, 8, 9}}, want: []int{1, 7, 8, 9, math.MaxInt}},
		{args: args{nums: []int{1, 7, math.MinInt, 8, 9}}, want: []int{math.MinInt, 1, 7, 8, 9}},
		{args: args{nums: []int{1, 7, math.MaxInt64, 8, 9}}, want: []int{1, 7, 8, 9, math.MaxInt64}},
		{args: args{nums: []int{1, 7, math.MinInt64, 8, 9}}, want: []int{math.MinInt64, 1, 7, 8, 9}},
		{args: args{nums: []int{1, 3, 5, 7, 9}}, want: []int{1, 3, 5, 7, 9}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bubbleSort(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bubbleSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
