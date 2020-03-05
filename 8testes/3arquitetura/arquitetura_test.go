package arquitetura_test

import (
	"runtime"
	"testing"
)

func TestDependente(t *testing.T) {
	if runtime.GOARCH == "amd64" {
		t.Skip("Não Funciona em arquitetura amd64")
	}
	t.Fail()
}
