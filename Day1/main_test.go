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

func TestGetListsFromLines(t *testing.T) {
	type args struct {
		text []string
	}
	tests := []struct {
		name  string
		args  args
		want  []int
		want1 []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := GetListsFromLines(tt.args.text)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetListsFromLines() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("GetListsFromLines() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestCreateFrequencyMap(t *testing.T) {
	type args struct {
		rightList []int
	}
	tests := []struct {
		name string
		args args
		want map[int]int
	}{
		{
			name: "Normal case",
			args: args{
				rightList: []int{1, 2, 2, 3, 3, 3},
			},
			want: map[int]int{1: 1, 2: 2, 3: 3},
		},
		{
			name: "Empty list",
			args: args{
				rightList: []int{},
			},
			want: map[int]int{},
		},
		{
			name: "All same values",
			args: args{
				rightList: []int{5, 5, 5, 5},
			},
			want: map[int]int{5: 4},
		},
		{
			name: "Single element",
			args: args{
				rightList: []int{7},
			},
			want: map[int]int{7: 1},
		},
		{
			name: "Negative values",
			args: args{
				rightList: []int{-1, -2, -2, -3, -3, -3},
			},
			want: map[int]int{-1: 1, -2: 2, -3: 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateFrequencyMap(tt.args.rightList); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateFrequencyMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculateSimilarity(t *testing.T) {
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
			name: "Given case",
			args: args{
				leftList:  []int{1, 2, 3, 3, 3, 4}, // Sorted
				rightList: []int{3, 3, 3, 4, 5, 9}, // Sorted
			},
			want: 31,
		},
		{
			name: "No matching elements",
			args: args{
				leftList:  []int{1, 1, 1},
				rightList: []int{2, 2, 2},
			},
			want: 0, // No matches between the lists
		},
		{
			name: "All matching elements",
			args: args{
				leftList:  []int{1, 2, 3},
				rightList: []int{1, 2, 3},
			},
			want: 6,
		},
		{
			name: "Duplicate elements",
			args: args{
				leftList:  []int{3, 3, 4},
				rightList: []int{3, 3, 3, 4},
			},
			want: 22,
		},
		{
			name: "Empty left list",
			args: args{
				leftList:  []int{},
				rightList: []int{1, 2, 3},
			},
			want: 0, // Empty left list results in no similarity
		},
		{
			name: "Empty right list",
			args: args{
				leftList:  []int{1, 2, 3},
				rightList: []int{},
			},
			want: 0, // Empty right list results in no similarity
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateSimilarity(tt.args.leftList, tt.args.rightList); got != tt.want {
				t.Errorf("CalculateSimilarity() = %v, want %v", got, tt.want)
			}
		})
	}
}
