// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

import (
	"testing"

	"github.com/placeholderplaceholderplaceholder/opentf/internal/legacy/opentf"
)

// TestResourceDataRaw creates a ResourceData from a raw configuration map.
func TestResourceDataRaw(
	t *testing.T, schema map[string]*Schema, raw map[string]interface{}) *ResourceData {
	t.Helper()

	c := opentf.NewResourceConfigRaw(raw)

	sm := schemaMap(schema)
	diff, err := sm.Diff(nil, c, nil, nil, true)
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	result, err := sm.Data(nil, diff)
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	return result
}
