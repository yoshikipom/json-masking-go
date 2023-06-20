package masking

import (
	"bytes"
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
)

const defaltmaskingValue = "*"

type masking struct {
	deniedKeySet map[string]struct{}
	useRegex     bool
	format       bool
}

func New(deniedKeys []string, useRejex bool, format bool) *masking {
	keySet := make(map[string]struct{})
	for _, deniedKey := range deniedKeys {
		keySet[deniedKey] = struct{}{}
	}
	return &masking{
		deniedKeySet: keySet,
		useRegex:     useRejex,
		format:       format,
	}
}

func (m *masking) Replace(body []byte) []byte {
	var data interface{}
	err := json.Unmarshal(body, &data)
	if err != nil {
		panic(err)
	}

	m.processData("", &data)

	b, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	fmt.Println("format")
	fmt.Println(m.format)

	if m.format {
		var output bytes.Buffer
		err = json.Indent(&output, b, "", "  ")
		return output.Bytes()
	} else {
		return b
	}
}

func (m *masking) processData(jsonPath string, node *interface{}) interface{} {
	switch n := (*node).(type) {
	case map[string]interface{}:
		for k, v := range n {
			newPath := fmt.Sprintf("%s.%s", jsonPath, k)
			result := m.processData(newPath, &v)
			if result != nil {
				n[k] = result
			}
		}
	case []interface{}:
		for i, v := range n {
			newPath := fmt.Sprintf("%s[%d]", jsonPath, i)
			result := m.processData(newPath, &v)
			if result != nil {
				n[i] = result
			}
		}
	default:
		jsonPath = strings.TrimPrefix(jsonPath, ".")
		fmt.Println(jsonPath)
		if m.denied(jsonPath) {
			return defaltmaskingValue
		}
		return nil
	}
	return nil
}

func (m *masking) denied(jsonPath string) bool {
	if m.useRegex {
		for k, _ := range m.deniedKeySet {
			r := regexp.MustCompile(k)
			if r.MatchString(jsonPath) {
				return true
			}
		}
		return false
	} else {
		_, ok := m.deniedKeySet[jsonPath]
		return ok
	}
}
