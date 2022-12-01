package day1

import (
	"io"
	"strings"
	"testing"
)

func TestSolvePartOne(t *testing.T) {
	type args struct {
		r io.Reader
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
					"1000\n" +
						"2000\n" +
						"3000\n\n" +
						"4000\n\n" +
						"5000\n" +
						"6000\n\n" +
						"7000\n" +
						"8000\n" +
						"9000\n\n" +
						"10000\n\n",
				),
			},
			want:    "24000",
			wantErr: false,
		},
		{
			name: "one elf only",
			args: args{
				strings.NewReader("1000\n"),
			},
			want:    "1000",
			wantErr: false,
		},
		{
			name: "multiple empty lines",
			args: args{
				strings.NewReader("\n\n\n\n\n"),
			},
			want:    "0",
			wantErr: false,
		},
		{
			name: "no input",
			args: args{
				strings.NewReader(""),
			},
			want:    "0",
			wantErr: false,
		},
		{
			name:    "reader is nil",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SolvePartOne(tt.args.r)
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
					"1000\n" +
						"2000\n" +
						"3000\n\n" +
						"4000\n\n" +
						"5000\n" +
						"6000\n\n" +
						"7000\n" +
						"8000\n" +
						"9000\n\n" +
						"10000\n\n",
				),
			},
			want:    "45000",
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
