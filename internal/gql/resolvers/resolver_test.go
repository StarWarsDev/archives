package resolvers

import "testing"

func Test_isRegexpMatch(t *testing.T) {
	type args struct {
		subject string
		pattern string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "basic string",
			args: args{
				subject: "this is a test",
				pattern: "test",
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "pattern string",
			args: args{
				subject: "this is a test",
				pattern: "(?i)Test",
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "pattern string",
			args: args{
				subject: "this is a check",
				pattern: "(?i)Test",
			},
			want:    false,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := isRegexpMatch(tt.args.subject, tt.args.pattern)
			if (err != nil) != tt.wantErr {
				t.Errorf("isRegexpMatch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("isRegexpMatch() got = %v, want %v", got, tt.want)
			}
		})
	}
}
