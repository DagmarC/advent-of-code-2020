package passp

import (
	"fmt"
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

// ValidatePassword validates password : password has to fulfill passwordPolicy criteria:
// Given letter must occur X-times, where X is in between the boundaries.
func (pd *Definition) ValidatePassword() bool {
	countLetter := strings.Count(pd.password, pd.policy.letter)
	if countLetter >= pd.policy.lowerBoundary && countLetter <= pd.policy.upperBoundary {
		fmt.Println("Valid: ", pd.password, *pd.policy)
		return true
	}
	return false
}
