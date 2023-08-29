package ignore

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadRulesNoHelmignore(t *testing.T) {
	r, err := LoadRules("chart_no_helmignore")
	assert.Nil(t, err)
	// default pattern
	assert.Len(t, r.patterns, 1)
}

func TestLoadRulesHelmignore(t *testing.T) {
	r, err := LoadRules("chart_helmignore")
	assert.Nil(t, err)
	assert.Len(t, r.patterns, 3)
}

func TestFilter(t *testing.T) {
	rules, err := Parse(strings.NewReader("/bar/\nREADME.md\n"))
	assert.Nil(t, err)
	files := []string{"Chart.yaml", "bar/xxx", "template/svc.yaml", "baz/bar/biz.txt", "README.md"}
	actual, err := FilterFiles(files, rules)
	assert.Nil(t, err)
	expected := []string{"Chart.yaml", "baz/bar/biz.txt", "template/svc.yaml"}
	assert.ElementsMatch(t, expected, actual)
}
