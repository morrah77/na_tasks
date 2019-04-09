package tokens_test

import (
	"testing"
	"gotest.tools/assert"
	"tasks/tokens"
)

func TestNewTokenChecker(t *testing.T) {
	var (
		checker *tokens.TokensChecker
		wordList []string = []string{`was`, `it`, `a`, `rat`, `i`, `saw`}
		expectedDictionary []string = []string{`was`, `saw`, `rat`, `it`, `i`, `a`}
	)
	checker = tokens.NewTokenChecker(wordList)
	assert.DeepEqual(t, checker.Dictionary(), expectedDictionary)
}

func TestTokensChecker_CheckTrue(t *testing.T) {
	var (
		checker *tokens.TokensChecker
		wordList []string = []string{`was`, `it`, `a`, `rat`, `i`, `saw`}
		phrase string = `wasitaratisaw`
		expectedDictionary []string = []string{`was`, `saw`, `rat`, `it`, `i`, `a`}
		checked bool
		resultExpected bool = true
	)
	checker = tokens.NewTokenChecker(wordList)
	assert.DeepEqual(t, checker.Dictionary(), expectedDictionary)
	checked = checker.Check(phrase)
	assert.Equal(t, checked, resultExpected, `Should check phrase`)
}

func TestTokensChecker_CheckFalse(t *testing.T) {
	var (
		checker *tokens.TokensChecker
		wordList []string = []string{`was`, `it`, `a`, `rat`, `i`, `saw`}
		phrase string = `wasitaratisaaw`
		expectedDictionary []string = []string{`was`, `saw`, `rat`, `it`, `i`, `a`}
		checked bool
		resultExpected bool = false
	)
	checker = tokens.NewTokenChecker(wordList)
	assert.DeepEqual(t, checker.Dictionary(), expectedDictionary)
	checked = checker.Check(phrase)
	assert.Equal(t, checked, resultExpected, `Should not check phrase`)
}

func TestTokensChecker_CheckTrue2(t *testing.T) {
	var (
		checker *tokens.TokensChecker
		wordList []string = []string{`a`, `what`, `an`, `nice`, `day`}
		phrase string = `whataniceday`
		expectedDictionary []string = []string{`what`, `nice`, `day`, `an`, `a`}
		checked bool
		resultExpected bool = true
	)
	checker = tokens.NewTokenChecker(wordList)
	assert.DeepEqual(t, checker.Dictionary(), expectedDictionary)
	checked = checker.Check(phrase)
	assert.Equal(t, checked, resultExpected, `Should check phrase`)
}

func TestTokensChecker_CheckFalse2(t *testing.T) {
	var (
		checker *tokens.TokensChecker
		wordList []string = []string{`a`, `what`, `an`, `nice`, `day`}
		phrase string = `dawhaty`
		expectedDictionary []string = []string{`what`, `nice`, `day`, `an`, `a`}
		checked bool
		resultExpected bool = false
	)
	checker = tokens.NewTokenChecker(wordList)
	assert.DeepEqual(t, checker.Dictionary(), expectedDictionary)
	checked = checker.Check(phrase)
	assert.Equal(t, checked, resultExpected, `Should not check phrase`)
}

func TestTokensChecker_CheckTrue5(t *testing.T) {
	var (
		checker *tokens.TokensChecker
		wordList []string = []string{`a`, `ab`, `bc`}
		phrase string = `aabcabcaaabbcababc`
		expectedDictionary []string = []string{`bc`, `ab`, `a`}
		checked bool
		resultExpected bool = true
	)
	checker = tokens.NewTokenChecker(wordList)
	assert.DeepEqual(t, checker.Dictionary(), expectedDictionary)
	checked = checker.Check(phrase)
	assert.Equal(t, checked, resultExpected, `Should check phrase`)
}
