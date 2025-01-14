package strconv

import (
	"strconv"
	"testing"
)

func TestAtoi(t *testing.T) {
	tests := []struct {
		name     string // 测试用例的名称
		input    string // 输入的字符串
		expected int    // 预期转换后的整数
		wantErr  bool   // 是否预期发生错误
	}{
		{"Positive number", "123456", 123456, false},
		{"Negative number", "-789012", -789012, false},
		{"Leading zeros", "0042", 42, false},
		{"Empty string", "", 0, true},               // Atoi 对于空字符串应该返回错误
		{"Non-digit characters", "abc123", 0, true}, // 包含非数字字符应该返回错误
		{"Spaces around number", "  42  ", 42, false},
		{"Max int", strconv.Itoa(int(^uint(0) >> 1)), int(^uint(0) >> 1), false}, // MaxInt
		{"Min int", strconv.Itoa(int(uint(0) >> 1)), int(uint(0) >> 1), false},   // MinInt
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := strconv.Atoi(tt.input)

			if (err != nil) != tt.wantErr {
				t.Fatalf("strconv.Atoi(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
			}

			if !tt.wantErr && actual != tt.expected {
				t.Errorf("expected %d, got %d", tt.expected, actual)
			}
		})
	}
}

func TestItoa(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected string
		wantErr  bool
	}{
		{"Positive number", 123456, "123456", false},
		{"Negative number", -789012, "-789012", false},
		{"Leading zeros", 42, "42", false},
		{"Non-digit characters", 0, "0", true}, // 包含非数字字符应该返回错误
		{"Max int", int(^uint(0) >> 1), strconv.Itoa(int(^uint(0) >> 1)), false}, // MaxInt
		{"Min int", int(uint(0) >> 1), strconv.Itoa(int(uint(0) >> 1)), false},   // MinInt
	}

	for _, test := range tests {
	    t.Run(test.name, func(t *testing.T) {
			actual := strconv.Itoa(test.input)

			if actual != test.expected {
			    t.Errorf("expected %s, got %s", test.expected, actual)
			}
		})
	}
}
