package auth

import (
	"errors"
	"net/http"
    "testing"
)

func TestAPIKey(t *testing.T){
	tests := []struct {
			name		string
			headers		http.Header
			wantKey		string
			wantErr		bool
			wantErrIs	error
	}{
		{
				name:		"no auth header",
				headers:	http.Header{},
				wantErr:	true,
				wantErrIs:	ErrNoAuthHeaderIncluded,
		},
		{
				name:		"wrong scheme",
				headers:	http.Header{"Authorization": []string{"Bearer abc123"}},
				wantErr:	true,
		},
		{
				name:		"missing key",
				headers: 	http.Header{"Authorization": []string{"ApiKey"}},
				wantErr:	true,

		},
		{
				name:    	"valid api key",
				headers: 	http.Header{"Authorization": []string{"ApiKey abc123"}},
				wantKey: 	"abc123",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T){
				gotKey, err := GetAPIKey(tc.headers)

				if tc.wantErr {
					if err == nil {
							t.Fatalf("expected an error, got nil")
					}
					if tc.wantErrIs != nil && !errors.Is(err, tc.wantErrIs) {
							t.Fatalf("expected error %v, got %v", tc.wantErrIs, err)
					}
					return
				}

				if err !=nil {
						t.Fatalf("unexpected error: %v", err)
				}
				if gotKey != tc.wantKey{
					t.Errorf("got key %q, want %q", gotKey, tc.wantKey)
				}
		})
	}
}