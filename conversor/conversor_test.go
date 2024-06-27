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
		numeroEmRomano := "XL"
		conversor := conv.Conve{}.NewConv()

		//act
		numero := conversor.Conversor(numeroEmRomano)

		//assert
		if numero != 0 {
			t.Errorf("O numero %d não deve ser diferente de 4", numero)
		}
	})
}

/* for i := 0; i < n; i++ {
    value := romanToIntMap[rune(s[i])]
    if i < n-1 && value < romanToIntMap[rune(s[i+1])] {
        total -= value
    } else {
        total += value
    }
}*/
