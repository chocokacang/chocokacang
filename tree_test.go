package chocokacang_test

import (
	"testing"

	"github.com/chocokacang/chocokacang"
	"github.com/stretchr/testify/assert"
)

func TestGetParam(t *testing.T) {
	param := chocokacang.Param{
		Key:   "key1",
		Value: "value1",
	}

	params := make(chocokacang.Params, 2)
	params = append(params, param)

	val1, found1 := params.Get("key1")
	assert.Equal(t, "value1", val1, "Value of param with key1 should be value1")
	assert.Equal(t, true, found1, "Found indicator should be true")

	val2, found2 := params.Get("key2")
	assert.Equal(t, "", val2, "Value of param with key2 should be empty")
	assert.Equal(t, false, found2, "Found indicator should be false")
}
