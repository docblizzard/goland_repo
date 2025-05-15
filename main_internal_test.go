package main

import "testing"

func Example_main() {
	main()
	// Output:
	// Hello world!
}

func TestGreet(t *testing.T) {
	type testCast struct {
		lang language
		want string
	}

	var tests = map[string]testCast{
		"English": {
			lang: "en",
			want: "Hello World!",
		},
		"French": {
			lang: "fr",
			want: "Bonjour le monde!",
		},
		"Spanish, not supported": {
			lang: "es",
			want: `Unsupported language: "es"`,
		},
		"Empty": {
			lang: "",
			want: `Unsupported language: ""`,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := greet(tc.lang)

			if got != tc.want {
				t.Errorf("expected: %q, got: %q", tc.want, got)
			}
		})
	}

}
