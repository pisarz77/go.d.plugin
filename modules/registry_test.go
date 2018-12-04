package modules

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRegister(t *testing.T) {
	modName := "modName"
	registry := make(Registry)

	// OK case
	assert.NotPanics(
		t,
		func() {
			registry.Register(modName, Creator{})
		})

	_, exist := registry[modName]

	require.True(t, exist)

	// Panic case
	assert.Panics(
		t,
		func() {
			registry.Register(modName, Creator{})
		})

}
