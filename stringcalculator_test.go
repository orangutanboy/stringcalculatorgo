package stringcalculator

import "testing"
import "github.com/stretchr/testify/assert"

func TestMain(m *testing.M) {
	m.Run()
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add("//;\n1;2")
	}
}

func Test_EmptyStringReturns0(t *testing.T) {
	testCases := map[string]int{
		"": 0,
	}

	addAndCheckOutput(t, testCases)
}

func Test_SingleNumberReturnsThatNummber(t *testing.T) {
	testCases := map[string]int{
		"1": 1,
		"2": 2,
	}

	addAndCheckOutput(t, testCases)
}

func Test_CommaSeparatedNumbersReturnsSum(t *testing.T) {
	testCases := map[string]int{
		"1,2":     3,
		"1,2,1,5": 9,
	}

	addAndCheckOutput(t, testCases)
}

func Test_CommaOrNewlineSeparatedNumbersReturnsSum(t *testing.T) {
	testCases := map[string]int{
		"1\n2":      3,
		"1\n2,1\n5": 9,
	}

	addAndCheckOutput(t, testCases)
}

func Test_CustomDelimiterSeparatedNumbersReturnsSum(t *testing.T) {
	testCases := map[string]int{
		"//;\n1;2":   3,
		"//;\n1;2;3": 6,
	}
	addAndCheckOutput(t, testCases)
}

func Test_NegativeNumbersThrow(t *testing.T) {
	testCases := map[string]string{
		"-1":    "negatives not allowed -1",
		"-1,-2": "negatives not allowed -1-2",
	}

	for input, expectedErrorMessage := range testCases {
		thenErrorIsThrown(t, input, expectedErrorMessage)
	}
}

func Test_NumbersAbove1000Ignored(t *testing.T) {
	testCases := map[string]int{
		"2,1001":        2,
		"2,1000":        1002,
		"2,1002,4,1006": 6,
	}

	addAndCheckOutput(t, testCases)
}

func thenErrorIsThrown(t *testing.T, input string, expectedErrorMessage string) {
	_, e := Add(input)
	assert.Equal(t, e.Error(), expectedErrorMessage)
}

func addAndCheckOutput(t *testing.T, testCases map[string]int) {
	for input, expectedValue := range testCases {
		result, e := Add(input)
		assert.Equal(t, expectedValue, result)
		assert.Nil(t, e)
	}
}
