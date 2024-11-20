package conversor_test

import (
	conv "convrsorromano"
	"testing"
)

var conversor = conv.Conve{}.NewConv()

func TestConversorDeRomanoParaNatural(t *testing.T) {
	t.Run("deve converter I de romano para result natural", func(t *testing.T) {
		// arrange
		resultEmRomano := "I"

		//act
		result := conversor.Conversor(resultEmRomano)

		//assert
		if result != 1 {
			t.Errorf("O result %d não deve ser diferente de 1", result)
			t.Fail()
		}
	})

	t.Run("deve converter XXI de romano para natural", func(t *testing.T) {
		// arrange
		resultEmRomano := "XXI"
		expected := 21
		//act
		result := conversor.Conversor(resultEmRomano)

		//assert
		if result != expected {
			t.Errorf("O result %d não deve ser diferente do esperado %d", result, expected)
			t.Fail()
		}
	})

	t.Run("deve converter DLV de romano para natural", func(t *testing.T) {
		// arrange
		resultEmRomano := "DLV"
		expected := 555
		//act
		result := conversor.Conversor(resultEmRomano)

		//assert
		if result != expected {
			t.Errorf("O result %d não deve ser diferente do esperado %d", result, expected)
			t.Fail()
		}
	})

	t.Run("deve converter DX de romano para natural", func(t *testing.T) {
		// arrange
		resultEmRomano := "DX"
		expected := 510
		//act
		result := conversor.Conversor(resultEmRomano)

		//assert
		if result != expected {
			t.Errorf("O result %d não deve ser diferente do esperado %d", result, expected)
			t.Fail()
		}
	})

	t.Run("deve converter MD de romano para natural", func(t *testing.T) {
		// arrange
		resultEmRomano := "MD"
		expected := 1500
		//act
		result := conversor.Conversor(resultEmRomano)

		//assert
		if result != expected {
			t.Errorf("O result %d não deve ser diferente do esperado %d", result, expected)
			t.Fail()
		}
	})

}
