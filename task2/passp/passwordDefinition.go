package passp

import (
	"strings"
)

type Definition struct {
	password string
	policy   *passwordPolicy
}

type passwordPolicy struct {
	letter        string
	lowerBoundary int
	upperBoundary int
}

func CreatePasswordDefinition(password string, pp *passwordPolicy) *Definition {
	return &Definition{
		password: password,
		policy:   pp,
	}
}

func CreatePasswordPolicy(letter string, from, to int) *passwordPolicy {
	return &passwordPolicy{
		letter:        letter,
		lowerBoundary: from,
		upperBoundary: to,
	}
}

// ValidatePasswordFirst validates password : password has to fulfill passwordPolicy criteria:
// Given letter must occur X-times, where X is in between the boundaries.
func (pd *Definition) ValidatePasswordFirst() bool {
	countLetter := strings.Count(pd.password, string(pd.policy.letter))
	if countLetter >= pd.policy.lowerBoundary && countLetter <= pd.policy.upperBoundary {
		//fmt.Println("Valid: ", pd.password, *pd.policy)
		return true
	}
	return false
}

// ValidatePasswordSecond Each policy actually describes two positions in the password,
// where 1 means the first character, 2 means the second character, and so on. (Be careful; Toboggan Corporate Policies
// have no concept of "index zero"!) Exactly one of these positions must contain the given letter.
// Other occurrences of the letter are irrelevant for the purposes of policy enforcement.
// 1-3 a: abcde is valid: position 1 contains a and position 3 does not.
// 1-3 b: cdefg is invalid: neither position 1 nor position 3 contains b.
func (pd *Definition) ValidatePasswordSecond() bool {

	// Both can't be equal at the same time.
	if string(pd.password[pd.policy.lowerBoundary-1]) == pd.policy.letter &&
		string(pd.password[pd.policy.upperBoundary-1]) == pd.policy.letter {
		return false
	}
	// One on ether side can be equal at the same time.
	if string(pd.password[pd.policy.lowerBoundary-1]) == pd.policy.letter ||
		string(pd.password[pd.policy.upperBoundary-1]) == pd.policy.letter {
		return true
	}
	// Otherwise, false.
	return false
}
