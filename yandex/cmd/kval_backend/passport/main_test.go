package main

import "testing"

func TestValidatePassportCode(t *testing.T) {
	type args struct {
		code string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "1",
			args:    args{
				code: "1234567890",
			},
			want:    "1234-567890",
			wantErr: false,
		},
		{
			name:    "12",
			args:    args{
				code: "1234 567890",
			},
			want:    "1234-567890",
			wantErr: false,
		},
		{
			name:    "123",
			args:    args{
				code: "(1234) 567890",
			},
			want:    "1234-567890",
			wantErr: false,
		},
		{
			name:    "1234",
			args:    args{
				code: "(1234)567890",
			},
			want:    "1234-567890",
			wantErr: false,
		},
		{
			name:    "12345",
			args:    args{
				code: "1234-567890",
			},
			want:    "1234-567890",
			wantErr: false,
		},
		{
			name:    "123456",
			args:    args{
				code: "PC-1234-567890",
			},
			want:    "1234-567890",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ValidatePassportCode(tt.args.code)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidatePassportCode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ValidatePassportCode() got = %v, want %v", got, tt.want)
			}
		})
	}
}
