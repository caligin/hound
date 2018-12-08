package diffs_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/caligin/hound/diffs"
	"testing"
)

var diffListFixtures = []struct {
	desc     string
	expected []string
	actual   []string
	excess   []string
	missing  []string
}{
	{"empty actual yields expected as missing and empty excess", []string{"a", "c"}, []string{}, []string{}, []string{"a", "c"}},
	{"multiple instances of same element in either expected or actual are deduplicated in the results", []string{"a", "a"}, []string{"c", "c"}, []string{"c"}, []string{"a"}},
	{"empty expected yields actual as excess and empty missing", []string{}, []string{"a", "c"}, []string{"a", "c"}, []string{}},
	{"empty expected and actual yield empty excess and missing", []string{}, []string{}, []string{}, []string{}},
	{"empty and expected containing same elements yield empty excess and missing", []string{"a", "c"}, []string{"a", "c"}, []string{}, []string{}},
	{"empty and expected containing same elements yield empty excess and missing, even in different order", []string{"a", "c"}, []string{"c", "a"}, []string{}, []string{}},
	{"actual subset of expected yields empty excess and set difference as missing", []string{"a", "c"}, []string{"a"}, []string{}, []string{"c"}},
	{"expected subset of actual yields empty missing and set difference as excess", []string{"a"}, []string{"a", "c"}, []string{"c"}, []string{}},
	{"expected and actual disjointed yields expected as missing and actual as excess", []string{"a"}, []string{"b"}, []string{"b"}, []string{"a"}},
	{"expected and actual different but intersecting yields respective set differences", []string{"a", "c"}, []string{"a", "b"}, []string{"b"}, []string{"c"}},
	{"nil expected is treated as empty expected", nil, []string{"c"}, []string{"c"}, []string{}},
	{"nil actual is treated as empty actual", []string{"c"}, nil, []string{}, []string{"c"}},
	{"nil expect and actual are treated as empty", nil, nil, []string{}, []string{}},
}

func TestDiffLists(t *testing.T) {
	for _, fixture := range diffListFixtures {
		excess, missing := diffs.DiffSetSlices(fixture.expected, fixture.actual)
		assert.ElementsMatch(t, fixture.missing, missing, fixture.desc + " (comparing \"missing\")")
		assert.ElementsMatch(t, fixture.excess, excess, fixture.desc + " (comparing \"excess\")")
	}
}
