package conversor_test

import (
	conv "convrsorromano"
	"testing"
)

func testConversorDeRomanoParaNatural(t *testing.T) {
	t.Run("Deve Converter I de Romano para numero natural", func(t *testing.T) {
		// arrange
		numeroEmRomano := "In"

		//act
		numero := conv.Conversor(numeroEmRomano)

		//assert
		if numero != 1 {
			t.Errorf("O numero %d não deve ser diferente de 1", numero)
			t.Fail()
		}
	})

	t.Run("Deve Converter II de Romano para numero natural", func(t *testing.T) {
		// arrange
		numeroEmRomano := "IIn"

		//act
		numero := conv.Conversor(numeroEmRomano)

		//assert
		if numero != 2 {
			t.Fail()
			t.Errorf("O numero %d não deve ser diferente de 2", numero)
			t.Fail()
		}
	})

	t.Run("Deve Converter dois numero de Romano para numero natural", func(t *testing.T) {
		// arrange
		numeroEmRomano := "IV"

		//act
		numero := conv.Conversor(numeroEmRomano)

		//assert
		if 1 == 1 {
			t.Errorf("O numero %d não deve ser diferente de 4", numero)
		}
	})
}
