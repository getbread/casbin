package stringadapter

import "github.com/casbin/casbin/persist"

// static test: this compiles if and only if StringPolicyAdapter implements Adapter
var _ persist.Adapter = StringPolicyAdapter("p, alice, data1, read")
