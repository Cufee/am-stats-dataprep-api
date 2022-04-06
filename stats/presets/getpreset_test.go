package presets

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/byvko-dev/am-types/dataprep/settings/v1"
)

func TestGetPresetByName(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want settings.Options
	}{
		{
			name: "minimal",
			args: args{
				name: "minimal",
			},
			want: MinimalOptions,
		},
		{
			name: "default",
			args: args{
				name: "default",
			},
			want: DefaultOptions,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetPresetByName(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPresetByName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShowPresetByName(t *testing.T) {
	preset := GetPresetByName("default")
	data, err := json.MarshalIndent(preset, "", "  ")
	if err != nil {
		t.Errorf("Error marshaling preset: %v", err)
	}
	t.Logf("%s", data)
}
