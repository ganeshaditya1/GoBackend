package middleware

import "testing"

func TestOriginAllowed_ExactAndWildcard(t *testing.T) {
	tests := []struct {
		name    string
		origin  string
		allowed []string
		want    bool
	}{
		{
			name:    "allow all wildcard",
			origin:  "https://api.example.com",
			allowed: []string{"*"},
			want:    true,
		},
		{
			name:    "exact https origin match",
			origin:  "https://example.com",
			allowed: []string{"https://example.com"},
			want:    true,
		},
		{
			name:    "exact http does not match https",
			origin:  "http://example.com",
			allowed: []string{"https://example.com"},
			want:    false,
		},
		{
			name:    "exact with port",
			origin:  "https://example.com:8443",
			allowed: []string{"https://example.com:8443"},
			want:    true,
		},
		{
			name:    "scheme+host wildcard allowed (https://*.example.com)",
			origin:  "https://api.example.com",
			allowed: []string{"https://*.example.com"},
			want:    true,
		},
		{
			name:    "scheme mismatch for scheme-wildcard",
			origin:  "http://api.example.com",
			allowed: []string{"https://*.example.com"},
			want:    false,
		},
		{
			name:    "host-only wildcard matches host",
			origin:  "https://sub.example.com",
			allowed: []string{"*.example.com"},
			want:    true,
		},
		{
			name:    "host-only exact host match (no scheme)",
			origin:  "http://api.example.com",
			allowed: []string{"api.example.com"},
			want:    true,
		},
		{
			name:    "host-only entry should not match different host",
			origin:  "https://other.com",
			allowed: []string{"api.example.com"},
			want:    false,
		},
		{
			name:    "invalid origin returns false",
			origin:  "://bad-origin",
			allowed: []string{"*"},
			want:    false,
		},
		{
			name:    "empty origin returns false",
			origin:  "",
			allowed: []string{"*"},
			want:    false,
		},
		{
			name:    "allowed entry without scheme matches host when origin has port",
			origin:  "https://api.example.com:8080",
			allowed: []string{"api.example.com:8080"},
			want:    true,
		},
		{
			name:    "allowed host without port should match origin without port",
			origin:  "https://example.com",
			allowed: []string{"example.com"},
			want:    true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := originAllowed(tc.origin, tc.allowed)
			if got != tc.want {
				t.Fatalf("originAllowed(%q, %v) = %v; want %v", tc.origin, tc.allowed, got, tc.want)
			}
		})
	}
}
