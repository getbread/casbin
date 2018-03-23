package stringadapter

import (
	"errors"
	"strings"

	"github.com/casbin/casbin/model"
	"github.com/casbin/casbin/persist"
)

// StringPolicyAdapter is a const policy, useful for when you
// actually have an unchanging policy known at build-time, or
// if your dynamic policy is always rebuilt at runtime, using
// the casbin API, instead of loaded via policy persistence.
type StringPolicyAdapter string

// LoadPolicy loads all policy rules from the storage.
func (a StringPolicyAdapter) LoadPolicy(model model.Model) error {
	for _, line := range strings.Split(string(a), "\n") {
		persist.LoadPolicyLine(line, model)
	}
	return nil
}

func (a StringPolicyAdapter) SavePolicy(model model.Model) error {
	return errors.New("not implemented")
}

func (a StringPolicyAdapter) AddPolicy(sec string, ptype string, rule []string) error {
	return errors.New("not implemented")
}

func (a StringPolicyAdapter) RemovePolicy(sec string, ptype string, rule []string) error {
	return errors.New("not implemented")
}

func (a StringPolicyAdapter) RemoveFilteredPolicy(sec string, ptype string, fieldIndex int, fieldValues ...string) error {
	return errors.New("not implemented")
}
