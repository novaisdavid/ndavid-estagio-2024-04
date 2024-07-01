package conversor_test

import (
	conv "convrsorromano"
	"testing"
)

func TestConversorDeRomanoParaNatural(t *testing.T) {
	t.Run("Deve Converter I de Romano para numero natural", func(t *testing.T) {
		// arrange
		numeroEmRomano := "I"
		conversor := conv.Conve{}.NewConv()

		//act
		numero := conversor.Conversor(numeroEmRomano)

		//assert
		if numero != 1 {
			t.Errorf("O numero %d não deve ser diferente de 1", numero)
			t.Fail()
		}
	})

	t.Run("Deve Converter II de Romano para numero natural", func(t *testing.T) {
		// arrange
		numeroEmRomano := "II"
		conversor := conv.Conve{}.NewConv()

		//act
		numero := conversor.Conversor(numeroEmRomano)

		//assert
		if numero != 2 {
			t.Errorf("O numero %d não deve ser diferente de 2", numero)
			t.Fail()
		}
	})

	t.Run("Deve Converter dois numero de Romano para numero natural", func(t *testing.T) {
		// arrange
		numeroEmRomano := "IV"
		conversor := conv.Conve{}.NewConv()

		//act
		numero := conversor.Conversor(numeroEmRomano)

		//assert
		if numero == 0 {
			t.Errorf("O numero %d não deve ser diferente de 4", numero)
		}
	})

	t.Run("Deve Converter XIX numero de Romano para numero natural", func(t *testing.T) {
		// arrange
		numeroEmRomano := "XIX"
		conversor := conv.Conve{}.NewConv()

		//act
		numero := conversor.Conversor(numeroEmRomano)

		//assert
		if numero != 19 {
			t.Errorf("O numero %d não deve ser diferente de 4", numero)
		}
	})

	t.Run("Deve Converter DX numero de Romano para numero natural", func(t *testing.T) {
		// arrange
		numeroEmRomano := "DX"
		conversor := conv.Conve{}.NewConv()

		//act
		numero := conversor.Conversor(numeroEmRomano)

		//assert
		if numero != 510 {
			t.Errorf("O numero %d não deve ser diferente de 4", numero)
		}
	})

	t.Run("Deve Converter MD numero de Romano para numero natural", func(t *testing.T) {
		// arrange
		numeroEmRomano := "MD"
		conversor := conv.Conve{}.NewConv()

		//act
		numero := conversor.Conversor(numeroEmRomano)

		//assert
		if numero != 1500 {
			t.Errorf("O numero %d não deve ser diferente de 1500", numero)
		}
	})


}
