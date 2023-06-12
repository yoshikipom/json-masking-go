package masking

import (
	"encoding/json"
	"fmt"
	"strings"
)

const defaltMaskingValue = "*"

type Masking struct {
	allowedFieldSet map[string]struct{}
}

func New() *Masking {
	fieldSet := make(map[string]struct{})
	fieldSet["email"] = struct{}{}
	return &Masking{allowedFieldSet: fieldSet}
}

func (m *Masking) Replace(body []byte) []byte {
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

func (m *Masking) processData(jsonPath string, node *interface{}) interface{} {
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
		if !m.allowd(jsonPath) {
			return defaltMaskingValue
		}
		return nil
	}
	return nil
}

func (m *Masking) allowd(jsonPath string) bool {
	_, ok := m.allowedFieldSet[jsonPath]
	return ok
}
