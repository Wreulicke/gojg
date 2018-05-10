package refs

import (
	"bufio"
	"strings"
	"testing"

	"github.com/wreulicke/gojg/parse"
)

func TestRefs(t *testing.T) {
	tests := []struct {
		src string
		len int
	}{
		{`"{{Test}}"`, 1},
		{`{{Test}}`, 1},
		{`{"{{Test}}": "{{Test}}"}`, 1},
		{`"{{Test}}"`, 1},
	}

	for _, v := range tests {
		reader := bufio.NewReader(strings.NewReader(v.src))
		node, err := parse.Parse(reader)

		if err != nil {
			t.Errorf("error occured. src:%s err: %v", v.src, err)
		}
		refs := Refs(node)
		if len := len(refs); len != v.len {
			t.Errorf("unexpected condition. src:%s expected: %d, actual: %d", v.src, v.len, len)
		}
	}
}
