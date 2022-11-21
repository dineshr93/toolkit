package toolkit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTools_Randomstring(t *testing.T) {
	var testtools *Tools
	assert.Equal(t, 10, len(testtools.Randomstring(10)), "Length size should be 10")

}
