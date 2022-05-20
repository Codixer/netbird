package system

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_LocalVersion(t *testing.T) {
	got := GetInfo(nil)
	want := "development"
	assert.Equal(t, want, got.WiretrusteeVersion)
}
