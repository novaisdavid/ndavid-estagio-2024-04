package formando_test

import (
	//formando "Firma/formando"
	"testing"
)

func VerificaVazio(t *testing.T, esperado, actual string) {
	if esperado != actual {
		t.Logf("%s != %s", esperado, actual)
		t.Fail()
	}
}
