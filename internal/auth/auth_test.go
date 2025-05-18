package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name    string
		header  http.Header
		wantKey string
		wantErr bool
	}{
		{
			name:    "valid ApiKey header",
			header:  http.Header{"Authorization": []string{"ApiKey my-secret-key"}},
			wantKey: "my-secret-key",
			wantErr: false,
		},
		{
			name:    "missing Authorization header",
			header:  http.Header{},
			wantKey: "",
			wantErr: true,
		},
		{
			name:    "malformed header - wrong prefix",
			header:  http.Header{"Authorization": []string{"Bearer token"}},
			wantKey: "",
			wantErr: true,
		},
		{
			name:    "malformed header - missing key",
			header:  http.Header{"Authorization": []string{"ApiKey"}},
			wantKey: "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotKey, err := GetAPIKey(tt.header)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAPIKey() error = %v, wantErr %v", err, tt.wantErr)
			}
			if gotKey != tt.wantKey {
				t.Errorf("GetAPIKey() = %v, want %v", gotKey, tt.wantKey)
			}
		})
	}
}
