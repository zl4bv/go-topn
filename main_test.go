package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestTopN(t *testing.T) {
	testcases := []struct {
		sizeN int
		lines []string
		want  []int64
	}{
		{
			1,
			[]string{"1"},
			[]int64{1},
		},
		{
			10,
			[]string{"1"},
			[]int64{1},
		},
		{
			5,
			[]string{"1", "20", "2", "19", "3", "18", "4", "17", "5", "16", "6", "15", "7", "14", "8", "13", "9", "12", "10", "11"},
			[]int64{20, 19, 18, 17, 16},
		},
		{
			10,
			[]string{"1", "20", "2", "19", "3", "18", "4", "17", "5", "16", "6", "15", "7", "14", "8", "13", "9", "12", "10", "11"},
			[]int64{20, 19, 18, 17, 16, 15, 14, 13, 12, 11},
		},
		{
			10,
			[]string{"1", "1", "1", "2", "2", "3"},
			[]int64{3, 2, 2, 1, 1, 1},
		},
	}

	for _, tc := range testcases {
		lines := strings.Join(tc.lines, "\n")
		reader := strings.NewReader(lines)

		got, err := TopN(tc.sizeN, reader)
		if err != nil {
			t.Error(err)
		}

		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("TopN(%v, %d) = %v; want %v", tc.lines, tc.sizeN, got, tc.want)
		}
	}
}
