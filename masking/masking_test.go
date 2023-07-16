package masking

import (
	"reflect"
	"testing"
)

func Test_jsonPath_String(t *testing.T) {
	tests := []struct {
		name string
		j    jsonPath
		want string
	}{
		{
			name: "Success",
			j:    []string{"[0]", "[1]", "name", "[2]"},
			want: "[0][1]name[2]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.j.String(); got != tt.want {
				t.Errorf("jsonPath.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	type args struct {
		input *MaskingInput
	}
	tests := []struct {
		name string
		args args
		want *Masking
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewWithFile(t *testing.T) {
	type args struct {
		configFile string
	}
	tests := []struct {
		name string
		args args
		want *Masking
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewWithFile(tt.args.configFile); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewWithFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_split(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
		want jsonPath
	}{
		{
			name: "Success",
			args: args{key: "friends[0].name"},
			want: jsonPath{"friends", "[0]", "name"},
		},
		{
			name: "Success with multi index",
			args: args{key: "[0][0].name"},
			want: jsonPath{"[0]", "[0]", "name"},
		},
		{
			name: "Success with key after index",
			args: args{key: "[0]friends[0]"},
			want: jsonPath{"[0]", "friends", "[0]"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := split(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("split() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMasking_Replace(t *testing.T) {
	type fields struct {
		deniedJsonPathList []jsonPath
		useRegex           bool
		format             bool
	}
	type args struct {
		body []byte
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Masking{
				deniedJsonPathList: tt.fields.deniedJsonPathList,
				useRegex:           tt.fields.useRegex,
				format:             tt.fields.format,
			}
			if got := m.Replace(tt.args.body); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Masking.Replace() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMasking_processData(t *testing.T) {
	type fields struct {
		deniedJsonPathList []jsonPath
		useRegex           bool
		format             bool
	}
	type args struct {
		path jsonPath
		node *interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Masking{
				deniedJsonPathList: tt.fields.deniedJsonPathList,
				useRegex:           tt.fields.useRegex,
				format:             tt.fields.format,
			}
			if got := m.processData(tt.args.path, tt.args.node); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Masking.processData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMasking_denied(t *testing.T) {
	type fields struct {
		deniedJsonPathList []jsonPath
		useRegex           bool
		format             bool
	}
	type args struct {
		path jsonPath
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Masking{
				deniedJsonPathList: tt.fields.deniedJsonPathList,
				useRegex:           tt.fields.useRegex,
				format:             tt.fields.format,
			}
			if got := m.denied(tt.args.path); got != tt.want {
				t.Errorf("Masking.denied() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMasking_regexMatch(t *testing.T) {
	type fields struct {
		deniedJsonPathList []jsonPath
		useRegex           bool
		format             bool
	}
	type args struct {
		path jsonPath
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Masking{
				deniedJsonPathList: tt.fields.deniedJsonPathList,
				useRegex:           tt.fields.useRegex,
				format:             tt.fields.format,
			}
			if got := m.regexMatch(tt.args.path); got != tt.want {
				t.Errorf("Masking.regexMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMasking_match(t *testing.T) {
	type fields struct {
		deniedJsonPathList []jsonPath
		useRegex           bool
		format             bool
	}
	type args struct {
		path jsonPath
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Masking{
				deniedJsonPathList: tt.fields.deniedJsonPathList,
				useRegex:           tt.fields.useRegex,
				format:             tt.fields.format,
			}
			if got := m.match(tt.args.path); got != tt.want {
				t.Errorf("Masking.match() = %v, want %v", got, tt.want)
			}
		})
	}
}
