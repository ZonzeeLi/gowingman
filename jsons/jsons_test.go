package jsons

import "testing"

func TestJsonEqual(t *testing.T) {
	testCases := []struct {
		a, b string
		want bool
	}{
		{
			a:    `{"name": "Tom", "age": 20, "isStudent": false, "grades": {"math": 90, "english": 85}, "friends": ["Jerry", "Spike"]}`,
			b:    `{"friends": ["Jerry", "Spike"], "grades": {"english": 85, "math": 90}, "isStudent": false, "age": 20, "name": "Tom"}`,
			want: true,
		},
		{
			a:    `{"name": "Tom", "age": 20, "isStudent": false, "grades": {"math": 90, "english": 85}, "friends": ["Jerry", "Spike"]}`,
			b:    `{"friends": ["Jerry", "Spike"], "grades": {"english": 85, "math": 91}, "isStudent": false, "age": 20, "name": "Tom"}`,
			want: false,
		},
		{
			a:    `{"name": "Tom", "age": 20, "isStudent": false, "grades": {"math": 90, "english": 85}, "friends": ["Jerry", "Spike"]}`,
			b:    `{"friends": ["Jerry", "Spike", "Tyke"], "grades": {"english": 85, "math": 90}, "isStudent": false, "age": 20, "name": "Tom"}`,
			want: false,
		},
		{
			a:    `{"name": "Tom", "age": 20, "isStudent": false, "grades": {"math": 90, "english": 85}, "friends": ["Jerry", "Spike"]}`,
			b:    `{"friends": ["Jerry", "Spike"], "grades": {"english": 85, "math": 90}, "isStudent": true, "age": 20, "name": "Tom"}`,
			want: false,
		},
	}

	for _, tc := range testCases {
		if got := JsonEqual(tc.a, tc.b); got != tc.want {
			t.Errorf("JsonEqual(%s, %s) = %v; want %v", tc.a, tc.b, got, tc.want)
		}
	}
}
