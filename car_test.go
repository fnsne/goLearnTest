package car

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	c, err := New("", 100)
	assert.NotNil(t, err)
	assert.Error(t, err)
	assert.Nil(t, c)

	c, err = New("foo", 100)
	assert.Nil(t, err)
	assert.NoError(t, err)
	assert.NotNil(t, c)
	assert.Equal(t, "foo", c.Name)
}

func TestSetName(t *testing.T) {
	c, err := New("bar", 200)
	assert.Nil(t, err)
	assert.NoError(t, err)
	assert.NotNil(t, c)
	assert.Equal(t, "bar", c.Name)

	c.SetName("abc")
	assert.Equal(t, "abc", c.Name)
}
