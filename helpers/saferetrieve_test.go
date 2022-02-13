package helpers

import (
	"reflect"
	"testing"

	"byvko.dev/repo/am-stats-dataprep-api/stats/presets"
	"byvko.dev/repo/am-stats-dataprep-api/stats/types"
)

func TestSafeRetrieve(t *testing.T) {
	type testCase struct {
		value interface{}
		path  string
		want  interface{}
	}

	testCases := []testCase{
		{
			value: presets.DefaultOptions,
			path:  "RatingBattles.Type",
			want:  types.OverviewTypeRating,
		},
		{
			value: map[string]interface{}{
				"a": map[string]interface{}{
					"b": map[string]interface{}{
						"c": "d",
					},
				},
			},
			path: "a.b.c",
			want: "d",
		},
		{
			value: map[string]interface{}{
				"a": map[string]interface{}{
					"b": map[string]interface{}{
						"c": "d",
					},
				},
			},
			path: "a.b.c.e",
			want: nil,
		},
		{
			value: map[string]interface{}{
				"a": map[string]interface{}{
					"b": map[string]interface{}{
						"c": "d",
					},
				},
			},
			path: "a.b.c.e.f",
			want: nil,
		},
	}

	for _, tc := range testCases {
		got := SafeRetrieve(tc.value, tc.path)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("SafeRetrieve(%v, %v) = %v,\n want %v", tc.value, tc.path, got, tc.want)
		}
	}
}
