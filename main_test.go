package main

import "testing"

type evenOddCount func(int) bool

func TestSanity(t *testing.T) {

	t.Run("Text solution", func(t *testing.T) {
		CheckCorrectness(t, textSolution)
	})
	t.Run("Text solution - shortcut", func(t *testing.T) {
		CheckCorrectness(t, textSolutionShortcut)
	})
	t.Run("Numeric solution", func(t *testing.T) {
		CheckCorrectness(t, numericSolution)
	})

}

func CheckCorrectness(t *testing.T, implementation evenOddCount) {
	inputs := []int{0, 1, 123, 42, 11111, 01, 100}
	correctResults := []bool{true, false, false, true, false, false, true}

	for i := 0; i < len(inputs); i++ {
		res := implementation(inputs[i])
		if res != correctResults[i] {
			t.Fail()
		}
	}
}

func BenchmarkEvenOddCount(b *testing.B) {
	var input int = 123

	b.Run("Text solution", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			textSolution(input)
		}
	})
	b.Run("Text solution - shortcut", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			textSolutionShortcut(input)
		}
	})
	b.Run("Numeric solution", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			numericSolution(input)
		}
	})
}
