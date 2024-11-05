package fizzbuzz_test

import(
	fizzbuz "fizzbuzz"
	"testing"
)

func verifica(t *testing.T, esperado, actual string) {
	if esperado != actual {
		t.Logf("%s != %s", esperado, actual)
		t.Fail()
	}
}

func TestNumeroSimples__NumeroSimples(t *testing.T){
	//arrange
	fizzbuz.Now()

	//act
	n := []int {2}

	//assert
	verifica(t,fizzbuz.DiferenteDoFizzbuzz(n), "2")

}

func TestNumero3__Fizz(t *testing.T){
	//arrange
	fizzbuz.Now()

	//act
	n := []int {3}

	//assert
	verifica(t,fizzbuz.Fizz(n), "fizz")

}

func TestNumero5__Buzz(t *testing.T){
	//arrange
	fizzbuz.Now()

	//act
	n := []int {5}

	//assert
	verifica(t,fizzbuz.Buzz(n), "buzz")

}

func TestNumero6__Fizz(t *testing.T){
	//arrange
	fizzbuz.Now()

	//act
	n := []int {6}

	//assert
	verifica(t,fizzbuz.Fizz(n), "fizz")

}

func TestNumero6__FizzBuzz(t *testing.T){
	//arrange
	fizzbuz.Now()

	//act
	n := []int {15}

	//assert
	verifica(t,fizzbuz.Fizzbuzz(n), "fizzbuzz")

}