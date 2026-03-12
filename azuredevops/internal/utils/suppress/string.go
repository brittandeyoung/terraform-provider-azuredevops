package suppress

import (
	"encoding/json"
	"reflect"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// CaseDifference reports whether old and new, interpreted as UTF-8 strings,
// are equal under Unicode case-folding.
func CaseDifference(_, old, new string, _ *schema.ResourceData) bool {
	return strings.EqualFold(old, new)
}

// WhitespaceJsonDifference reports whether old and new, interpreted as json strings,
// are equal excluding "whitespace-only" drift.
func WhitespaceJsonDifference(k, old, new string, d *schema.ResourceData) bool {
	var oldObj, newObj interface{}

	if old == "" || new == "" {
		return old == new
	}

	err := json.Unmarshal([]byte(old), &oldObj)
	if err != nil {
		return false
	}
	err = json.Unmarshal([]byte(new), &newObj)
	if err != nil {
		return false
	}

	return reflect.DeepEqual(oldObj, newObj)
}
