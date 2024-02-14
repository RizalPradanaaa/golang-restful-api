package test

import (
	"golang-restful-api/simple"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFileConnection(t *testing.T) {
	connection, close := simple.InitializedConnection("Coba")
	assert.NotNil(t, connection)
	close()
}
