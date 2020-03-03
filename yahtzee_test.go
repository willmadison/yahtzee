package yahtzee_test

import (
	"fmt"
	"testing"

	"github.com/willmadison/yahtzee"
	"github.com/stretchr/testify/assert"
)

func TestBasicUpperScoring(t *testing.T) {
	cases := []struct {
		given    []int
		expected int
	}{
		{
			[]int{2, 3, 5, 5, 6},
			10,
		},
		{
			[]int{1, 1, 1, 1, 3},
			4,
		},
		{
			[]int{1, 1, 1, 3, 3},
			6,
		},
		{
			[]int{1, 2, 3, 4, 5},
			5,
		},
		{
			[]int{6, 6, 6, 6, 6},
			30,
		},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("yahtzee.Score(%v)", tc.given), func(t *testing.T) {
			actual := yahtzee.Score(tc.given)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
