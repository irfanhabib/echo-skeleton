package main

import (
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

var echoSkeleton *EchoSkeleton

func TestMain(m *testing.M) {
	echoSkeleton = &EchoSkeleton{}
	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestHelloWorld(t *testing.T) {

	ec := echo.New()
	req := httptest.NewRequest(echo.GET, "/", strings.NewReader(""))
	rec := httptest.NewRecorder()
	context := ec.NewContext(req, rec)

	assert.NoError(t, echoSkeleton.helloWorldHandler(context))
	assert.Equal(t, "Hello World", rec.Body.String())
	assert.NotEqual(t, "Not Hellow World", rec.Body.String())

}
