//go:build integration
// +build integration

package goapp_test

import (
	"testing"

	goapp "github.com/bongnv/go-app-integration-testing"
	"github.com/stretchr/testify/require"
)

func Test_LoadDataFromDB(t *testing.T) {
	name, err := goapp.LoadDataFromDB()
	require.NoError(t, err)
	require.Equal(t, "One", name)
}
