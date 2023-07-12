package masking

import (
	"bytes"
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/yoshikipom/json-masking-go/config"
)

const defaltmaskingValue = "*"

type Masking struct {
	deniedKeySet map[string]struct{}
	useRegex     bool
	format       bool
}

type MaskingInput struct {
	DeniedKeyList []string
	UseRegex      bool
	Format        bool
}

func New(input *MaskingInput) *Masking {
	keySet := make(map[string]struct{})
	for _, deniedKey := range input.DeniedKeyList {
		keySet[deniedKey] = struct{}{}
	}
	return &Masking{
		deniedKeySet: keySet,
		useRegex:     input.UseRegex,
		format:       input.Format,
	}

}

func NewWithFile(configFile string) *Masking {
	err := config.Initialize(configFile)
	if err != nil {
		panic(err)
	}
	c := config.GetConfig()

	keySet := make(map[string]struct{})
	for _, deniedKey := range c.DeniedKeyList {
		keySet[deniedKey] = struct{}{}
	}
	return &Masking{
		deniedKeySet: keySet,
		useRegex:     c.UseRegex,
		format:       c.Format,
	}
}

func (m *Masking) Replace(body []byte) []byte {
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

	if m.format {
		var output bytes.Buffer
		err = json.Indent(&output, b, "", "  ")
		if err != nil {
			panic(err)
		}
		return output.Bytes()
	} else {
		return b
	}
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
		// fmt.Println(jsonPath)
		if m.denied(jsonPath) {
			return defaltmaskingValue
		}
		return nil
	}
	return nil
}

func (m *Masking) denied(jsonPath string) bool {
	if m.useRegex {
		for k := range m.deniedKeySet {
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
