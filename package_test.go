package ceph

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestImports(t *testing.T) {
	if !assert.Equal(t, 1, 1) {
		t.Error("Something is wrong.")
	}
}
