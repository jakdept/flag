package flagTrap

import (
	"testing"

	"github.com/alioygur/is"
	"github.com/stretchr/testify/assert"
)

func TestSetEmpty(t *testing.T) {
	t.Parallel()
	test := &StringTrap{
		Default: "pony",
	}

	assert.Nil(t, test.Set(""), "should have not gotten an error setting the test")
	assert.True(t, test.IsSet(), "I set the flag, it should show up set")
	assert.Equal(t, "pony", test.String(), "should have gotten the string")
}

func TestSetStringSame(t *testing.T) {
	t.Parallel()
	test := &StringTrap{
		Default: "pony",
	}

	assert.Nil(t, test.Set("pony"), "should have not gotten an error setting the test")
	assert.True(t, test.IsSet(), "I set the flag, it should show up set")
	assert.Equal(t, "pony", test.String(), "should have gotten the string")
}

func TestSetStringOverride(t *testing.T) {
	t.Parallel()
	test := &StringTrap{
		Default: "pony",
	}

	assert.Nil(t, test.Set("horsey"), "should have not gotten an error setting the test")
	assert.True(t, test.IsSet(), "I set the flag, it should show up set")
	assert.Equal(t, "horsey", test.String(), "should have gotten the string")
}

func TestNoSet(t *testing.T) {
	t.Parallel()
	test := &StringTrap{
		Default: "pony",
	}

	assert.False(t, test.IsSet(), "I set the flag, it should show up set")
	assert.Equal(t, "pony", test.String(), "should have gotten the string")
}
func TestDoubleSet(t *testing.T) {
	t.Parallel()
	test := &StringTrap{
		Default: "pony",
	}

	assert.Nil(t, test.Set("horsey"), "should have not gotten an error setting the test")
	assert.NotNil(t, test.Set("donkey"), "should have gotten an error setting the test")
	assert.True(t, test.IsSet(), "I set the flag, it should show up set")
	assert.Equal(t, "horsey", test.String(), "should have gotten the string")
}

func TestBadIsEmail(t *testing.T) {
	t.Parallel()
	test := &StringTrap{
		Validator: is.Email,
		Default:   "pony@thisisnotadomain.com",
	}

	assert.NotNil(t, test.Set("horsey"), "should have gotten an error setting the test")
	assert.False(t, test.IsSet(), "I should have failed setting the flag")
	assert.Equal(t, "pony@thisisnotadomain.com", test.String(), "should have gotten the string")
}

func TestIsAlpha(t *testing.T) {
	t.Parallel()
	test := &StringTrap{
		Validator: is.Alpha,
		Default:   "pony",
	}

	assert.Nil(t, test.Set("horsey"), "should have not gotten an error setting the test")
	assert.True(t, test.IsSet(), "I set the flag, it should show up set")
	assert.Equal(t, "horsey", test.String(), "should have gotten the string")
}
