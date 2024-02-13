package auth

//write a test for the GetAPIKey function
//use the following test cases:
//	- test that the function returns an error if no Authorization header is included
//	- test that the function returns an error if the Authorization header is malformed

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name    string
		headers http.Header
		want    string
		wantErr error
	}{
		{
			name:    "no authorization header included",
			headers: http.Header{},
			want:    "",
			wantErr: ErrNoAuthHeaderIncluded,
		},
		{
			name:    "malformed authorization header",
			headers: http.Header{"Authorization": []string{"Bearer"}},
			want:    "",
			wantErr: errors.New("malformed authorization header"),
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAPIKey(tt.headers)
			if err.Error() != tt.wantErr.Error() {
				t.Errorf("GetAPIKey() error = %v, wantErr = %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetAPIKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
