package pulid_prefix

import (
	"context"
	"fmt"

	"delivrio.io/shared-utils/pulid"
)

type LabelTable struct {
	L string
	T string
}

// prefixMap maps PULID prefixes to table names.
var prefixMap = map[pulid.ID]LabelTable{
	"LD": {T: "local_device", L: "local_device"},
	"AT": {T: "print_job", L: "print_job"},
	"AS": {T: "remote_connection", L: "remote_connection"},
	"LE": {T: "log_error", L: "log_error"},
	"UC": {T: "unique_computer", L: "unique_computer"},
}

// IDToType maps a pulid.ID to the underlying table.
func IDToType(ctx context.Context, id pulid.ID) (string, error) {
	if len(id) < 2 {
		return "", fmt.Errorf("IDToType: id too short")
	}
	prefix := id[:2]
	if val, ok := prefixMap[prefix]; ok {
		return val.T, nil
	}

	return "", fmt.Errorf("IDToType: could not map prefix '%s' to a type", prefix)
}

func TypeToPrefix(label string) string {

	for p, t := range prefixMap {
		if t.L == label {
			return string(p)
		}
	}

	panic(fmt.Sprintf("label not found: %s", label))
}
