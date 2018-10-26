package rc

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInternalNoop(t *testing.T) {
	call := Calls.Get("rc/noop")
	assert.NotNil(t, call)
	in := Params{
		"String": "hello",
		"Int":    42,
	}
	out, err := call.Fn(in)
	require.NoError(t, err)
	require.NotNil(t, out)
	assert.Equal(t, in, out)
}

func TestInternalError(t *testing.T) {
	call := Calls.Get("rc/error")
	assert.NotNil(t, call)
	in := Params{}
	out, err := call.Fn(in)
	require.Error(t, err)
	require.Nil(t, out)
}

func TestInternalList(t *testing.T) {
	call := Calls.Get("rc/list")
	assert.NotNil(t, call)
	in := Params{}
	out, err := call.Fn(in)
	require.NoError(t, err)
	require.NotNil(t, out)
	assert.Equal(t, Params{"commands": Calls.List()}, out)
}

func TestCorePid(t *testing.T) {
	call := Calls.Get("core/pid")
	assert.NotNil(t, call)
	in := Params{}
	out, err := call.Fn(in)
	require.NoError(t, err)
	require.NotNil(t, out)
	pid := out["pid"]
	assert.NotEqual(t, nil, pid)
	_, ok := pid.(int)
	assert.Equal(t, true, ok)
}

func TestCoreMemstats(t *testing.T) {
	call := Calls.Get("core/memstats")
	assert.NotNil(t, call)
	in := Params{}
	out, err := call.Fn(in)
	require.NoError(t, err)
	require.NotNil(t, out)
	sys := out["Sys"]
	assert.NotEqual(t, nil, sys)
	_, ok := sys.(uint64)
	assert.Equal(t, true, ok)
}

func TestCoreGC(t *testing.T) {
	call := Calls.Get("core/gc")
	assert.NotNil(t, call)
	in := Params{}
	out, err := call.Fn(in)
	require.NoError(t, err)
	require.Nil(t, out)
	assert.Equal(t, Params(nil), out)
}