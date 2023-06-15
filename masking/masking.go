package masking

import (
	"encoding/json"
	"fmt"
	"strings"
)

const defaltmaskingValue = "*"

type masking struct {
	deniedKeySet map[string]struct{}
}

func New(deniedKeys []string) *masking {
	keySet := make(map[string]struct{})
	for _, deniedKey := range deniedKeys {
		keySet[deniedKey] = struct{}{}
	}
	return &masking{deniedKeySet: keySet}
}

func (m *masking) Replace(body []byte) []byte {
	var data interface{}
	err := json.Unmarshal(body, &data)
	if err != nil {
		panic(err)
	}

	m.processData("", &data)

	output, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	return output
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
		if m.denied(jsonPath) {
			return defaltmaskingValue
		}
		return nil
	}
	return nil
}

func (m *masking) denied(jsonPath string) bool {
	_, ok := m.deniedKeySet[jsonPath]
	return ok
}
