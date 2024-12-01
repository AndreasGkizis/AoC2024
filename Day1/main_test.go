package main

import (
	"os"
	"reflect"
	"testing"
)

func TestCalculateDiffs(t *testing.T) {
	type args struct {
		leftList  []int
		rightList []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example case",
			args: args{
				leftList:  []int{1, 2, 3, 3, 3, 4}, // Sorted
				rightList: []int{3, 3, 3, 4, 5, 9}, // Sorted
			},
			want: 11,
		},
		{
			name: " All same values",
			args: args{
				leftList:  []int{1, 2, 3, 3, 3, 4},
				rightList: []int{1, 2, 3, 3, 3, 4},
			},
			want: 0,
		},
		{
			name: "Example case flipped",
			args: args{
				leftList:  []int{3, 3, 3, 4, 5, 9}, // Sorted
				rightList: []int{1, 2, 3, 3, 3, 4}, // Sorted
			},
			want: 11,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateDiffs(tt.args.leftList, tt.args.rightList); got != tt.want {
				t.Errorf("CalculateDiffs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReadFile(t *testing.T) {
	// create dummy file
	tempFile, err := os.CreateTemp("", "testfile*.txt")

	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tempFile.Name()) // Clean up the file after the test

	testContent := "line1\nline2\nline3\n"
	// Write test content to the temp file
	_, err = tempFile.Write([]byte(testContent))
	if err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	tempFile.Close()

	type args struct {
		filename string
	}
	tests := []struct {
		name      string
		args      args
		want      []string
		expectErr bool
	}{
		{
			name:      "Valid file",
			want:      []string{"line1", "line2", "line3"},
			expectErr: false,
			args: args{
				filename: tempFile.Name(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadFile(tt.args.filename)
			if (err != nil) != tt.expectErr {
				t.Errorf("ReadFile() error = %v, expectErr = %v", err, tt.expectErr)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
