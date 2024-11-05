package fizzbuzz

import(
		"strconv"
)

type FizzBuzz string

func Now() string{
	return ""
}

func Fizz(n []int) string {

	if n[0] % 3 == 0 {

		return "fizz"
	}

	return ""
}

func Buzz(n []int) string {

	if n[0] % 5 == 0 {

		return "buzz"
	}

	return ""

}

func Fizzbuzz(n []int) string {

	if n[0] % 3 == 0 && n[0] % 5 == 0 {

		return "fizzbuzz"
	}

	return ""
}

func DiferenteDoFizzbuzz(n []int) string {

	if n[0] % 3 != 0 && n[0] % 5 != 0 {

		return strconv.Itoa(n[0])
	}
	return ""
}