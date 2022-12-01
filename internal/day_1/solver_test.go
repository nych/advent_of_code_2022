package day1

import (
	"bufio"
	"strings"
	"testing"
)

func TestSolve(t *testing.T) {
	type args struct {
		s *bufio.Scanner
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "example input",
			args: args{
				bufio.NewScanner(strings.NewReader(
					"1000" +
						"2000" +
						"3000\n" +
						"4000\n" +
						"5000" +
						"6000\n" +
						"7000" +
						"8000" +
						"9000\n" +
						"10000",
				)),
			},
			want:    "24000",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Solve(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("Solve() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
