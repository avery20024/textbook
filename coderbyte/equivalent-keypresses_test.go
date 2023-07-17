package coderbyte

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/slices"
)

/*
Write a function that takes an array containing two strings where each string represents keypresses separated by commas. For this problem, a keypress can be either a printable character or a backspace (represented by -B). Your function should determine if the two strings of keypresses are equivalent.

You can produce a printable string from such a string of keypresses by having backspaces erase one preceding character. Consider two strings of keypresses equivalent if they produce the same printable string. For example:

checkEquivalentKeypresses(["a,b,c,d", "a,b,c,c,-B,d"]) // true
checkEquivalentKeypresses(["-B,-B,-B,c,c", "c,c"]) // true
checkEquivalentKeypresses(["", "a,-B,-B,a,-B,a,b,c,c,c,d"]) // false
*/
func Test(t *testing.T) {
	var arr1 = []string{"a", "b", "c", "d"}
	var arr2 = []string{"a", "b", "c", "d", "-B", "d", "-B", "d"}
	var arr3 = []string{"f", "e", "d", "h"}
	var arr4 = []string{"f", "e", "d", "h", "-B", "h", "-B", "h"}

	assert.True(t, EqualByPop(arr1, arr2), "EqualByPop")
	assert.True(t, EqualByGOTO(arr3, arr4), "EqualByGOTO")
}

const DELETE = "-B"

func EqualByPop(arrA, arrB []string) bool {
	var ret = []string{}

	for _, v := range arrB {
		if v == DELETE {
			ret = ret[0 : len(ret)-1]
			continue
		}

		ret = append(ret, v)
	}

	return slices.Equal(arrA, ret)
}

func EqualByGOTO(arrA, arrB []string) bool {
	var ret = arrB

CHECK_DELETE_SIGN:
	if slices.Contains(ret, DELETE) {
		idx := slices.Index(ret, DELETE)
		ret = append(ret[0:idx-1], ret[idx+1:len(ret)]...)
		goto CHECK_DELETE_SIGN
	}

	return slices.Equal(arrA, ret)
}
