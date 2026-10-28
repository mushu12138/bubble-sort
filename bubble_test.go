package main

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		want []int
	}{
		{"普通数组", []int{64, 34, 25, 12, 22, 11, 90}, []int{11, 12, 22, 25, 34, 64, 90}},
		{"已有序", []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{"逆序", []int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{"含重复元素", []int{3, 1, 2, 3, 1}, []int{1, 1, 2, 3, 3}},
		{"含负数", []int{-2, 5, 0, -8, 3}, []int{-8, -2, 0, 3, 5}},
		{"单元素", []int{42}, []int{42}},
		{"空切片", []int{}, []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := make([]int, len(tt.in)) // 拷贝，避免污染用例数据
			copy(got, tt.in)
			got = BubbleSort(got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BubbleSort(%v) = %v, 期望 %v", tt.in, got, tt.want)
			}
		})
	}
}
