package echoext

import (
	"testing"
)

// test deep validate
func TestValidator_validate(t *testing.T) {
	type Inner struct {
		InnerName string `validate:"required"`
	}

	tests := []struct {
		name    string
		obj     interface{}
		wantErr bool
	}{
		{
			name: "simple required 1",
			obj: &struct {
				Name string `validate:"required"`
			}{Name: "name"},
			wantErr: false,
		},

		{
			name: "simple required 2",
			obj: &struct {
				Name string `validate:"required"`
			}{},
			wantErr: true,
		},

		{
			name: "Inner required 1",
			obj: &struct {
				Name string `validate:"required"`
				Inner
			}{"name", Inner{InnerName: "Inner"}},
			wantErr: false,
		},

		{
			name: "Inner required 2",
			obj: &struct {
				Name string `validate:"required"`
				Inner
			}{"", Inner{InnerName: "Inner"}},
			wantErr: true,
		},

		{
			name: "Inner required 3",
			obj: &struct {
				Inner
			}{},
			wantErr: true,
		},

		{
			name: "Inner required 4",
			obj: &struct {
				Inner Inner
			}{},
			wantErr: true,
		},

		{
			name: "Inner required 5",
			obj: &struct {
				Inner
			}{Inner{InnerName: "ok"}},
			wantErr: false,
		},

		{
			name: "Inner required 6",
			obj: &struct {
				Inner Inner
			}{Inner: Inner{InnerName: "ok"}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		v := &Validator{}
		if err := v.validate(tt.obj); (err != nil) != tt.wantErr {
			t.Errorf("%q. Validator.validate() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}
