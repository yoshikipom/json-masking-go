package masking

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"github.com/yoshikipom/json-masking-go/config"
)

const defaltmaskingValue = "*"

type Masking struct {
	deniedJsonPathList []jsonPath
	useRegex           bool
	format             bool
}

type MaskingInput struct {
	DeniedKeyList []string
	UseRegex      bool
	Format        bool
}

type jsonPath []string

func (j jsonPath) String() string {
	arr := []string{}
	for i, str := range j {
		if i == 0 || !strings.HasPrefix(str, "[") {
			arr = append(arr, str)
		} else {
			arr[len(arr)-1] += str
		}
	}
	return strings.Join(arr, ".")
}

func New(input *MaskingInput) *Masking {
	deniedJsonPathList := []jsonPath{}
	for _, k := range input.DeniedKeyList {
		deniedJsonPath := split(k)
		deniedJsonPathList = append(deniedJsonPathList, jsonPath(deniedJsonPath))
	}

	return &Masking{
		deniedJsonPathList: deniedJsonPathList,
		useRegex:           input.UseRegex,
		format:             input.Format,
	}
}

func NewWithFile(configFile string) *Masking {
	err := config.Initialize(configFile)
	if err != nil {
		panic(err)
	}
	c := config.GetConfig()

	deniedJsonPathList := []jsonPath{}
	for _, k := range c.DeniedKeyList {
		deniedJsonPath := split(k)
		deniedJsonPathList = append(deniedJsonPathList, deniedJsonPath)
	}

	return &Masking{
		deniedJsonPathList: deniedJsonPathList,
		useRegex:           c.UseRegex,
		format:             c.Format,
	}
}

func split(key string) jsonPath {
	re := regexp.MustCompile(`(\[.*?\])|([^.\[\]]+)`)
	matches := re.FindAllString(key, -1)

	var result []string
	result = append(result, matches...)

	return result
}

func (m *Masking) Replace(body []byte) []byte {
	var data interface{}
	err := json.Unmarshal(body, &data)
	if err != nil {
		panic(err)
	}

	m.processData(jsonPath{}, &data)

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

func (m *Masking) processData(path jsonPath, node *interface{}) interface{} {
	switch n := (*node).(type) {
	case map[string]interface{}:
		for k, v := range n {
			newPath := append(path, k)
			result := m.processData(newPath, &v)
			if result != nil {
				n[k] = result
			}
		}
	case []interface{}:
		for i, v := range n {
			newPath := append(path, fmt.Sprintf("[%d]", i))
			result := m.processData(newPath, &v)
			if result != nil {
				n[i] = result
			}
		}
	default:
		// fmt.Println(jsonPath)
		if m.denied(path) {
			return defaltmaskingValue
		}
		return nil
	}
	return nil
}

func (m *Masking) denied(path jsonPath) bool {
	if m.useRegex {
		return m.regexMatch(path)
	} else {
		return m.match(path)
	}
}

func (m *Masking) regexMatch(path jsonPath) bool {
	for _, deniedPath := range m.deniedJsonPathList {
		r := regexp.MustCompile(deniedPath.String())
		if r.MatchString(path.String()) {
			return true
		}
	}
	return false
}

func (m *Masking) match(path jsonPath) bool {
	for _, deniedPath := range m.deniedJsonPathList {
		l := len(deniedPath)
		if len(path) < l {
			continue
		}
		if reflect.DeepEqual(deniedPath, path[:l]) {
			return true
		}
	}
	return false
}
