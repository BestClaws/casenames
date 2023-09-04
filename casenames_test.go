package casenames

import (
	"reflect"
	"testing"
)

func Test_CaseNames(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"first",
			args{
				"helloWorld foo_bar_split This camelCase_  _BSTUIcsIs Here Multiple   Spaces",
			},
			[]string{
				"hello", "World", "foo", "bar", "split", "This", "camel", "Case", "BSTU", "Ics", "Is", "Here", "Multiple", "Spaces",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CaseNames(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("caseNames() = %v, want %v", got, tt.want)
			}
		})
	}

}
