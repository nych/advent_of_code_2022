package day2

import (
	"io"
	"strings"
	"testing"
)

func TestSolvePartOne(t *testing.T) {
	type args struct {
		reader io.Reader
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
				strings.NewReader(
					"A Y\n" +
						"B X\n" +
						"C Z\n\n",
				),
			},
			want:    "15",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SolvePartOne(tt.args.reader)
			if (err != nil) != tt.wantErr {
				t.Errorf("SolvePartOne() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SolvePartOne() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolvePartTwo(t *testing.T) {
	type args struct {
		reader io.Reader
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
				strings.NewReader(
					"A Y\n" +
						"B X\n" +
						"C Z\n\n",
				),
			},
			want:    "12",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SolvePartTwo(tt.args.reader)
			if (err != nil) != tt.wantErr {
				t.Errorf("SolvePartTwo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SolvePartTwo() = %v, want %v", got, tt.want)
			}
		})
	}
}
