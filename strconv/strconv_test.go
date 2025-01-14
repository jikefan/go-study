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

func TestParseInt(t *testing.T) {
	// parse int
	intStr := "123"
	intValue, err := strconv.ParseInt(intStr, 10, 32)
	if err != nil {
		t.Error("parse int have a error, ", err)
	}

	if intValue != 123 {
		t.Error("expected 123, got", intValue)
	}
}

func TestParseBool(t *testing.T) {
	boolStr := "false"

	boolValue, err := strconv.ParseBool(boolStr)

	if err != nil {
		t.Error("have a error, ", err)
	}

	if boolValue != false {
		t.Errorf("expected %v, got %v", false, boolValue)
	}
}

func TestFormat(t *testing.T) {
	// 格式化整数
	intValue := 123
	intStr := strconv.FormatInt(int64(intValue), 10)
	t.Errorf("Formatted int string: %s\n", intStr)

	t.Error(strconv.FormatInt(int64(intValue), 36))

	t.Error(strconv.FormatInt(int64(intValue), 2))

	t.Error(strconv.FormatInt(int64(intValue), 16))

	// 格式化布尔值
	boolValue := true
	boolStr := strconv.FormatBool(boolValue)
	t.Errorf("Formatted bool string: %s\n", boolStr)

	// 格式化浮点数
	floatValue := 3.14
	floatStr := strconv.FormatFloat(floatValue, 'f', -1, 64)
	t.Errorf("Formatted float string: %s\n", floatStr)
}

func TestAppend(t *testing.T) {
	// 追加整数到字节数组
    num1 := 123
    byteSlice := []byte("Number: ")
    byteSlice = strconv.AppendInt(byteSlice, int64(num1), 10)
    t.Errorf("Appended int: %s\n", byteSlice)

    // 追加布尔值到字节数组
    boolVal := true
    byteSlice = []byte("Bool: ")
    byteSlice = strconv.AppendBool(byteSlice, boolVal)
    t.Errorf("Appended bool: %s\n", byteSlice)

    // 追加浮点数到字节数组
    floatVal := 3.14
    byteSlice = []byte("Float: ")
    byteSlice = strconv.AppendFloat(byteSlice, floatVal, 'f', -1, 64)
    t.Errorf("Appended float: %s\n", byteSlice)

	byteSlice = strconv.AppendQuote(byteSlice, "[]")
	t.Errorf("%s\n", byteSlice)
}

func TestQuote(t *testing.T) {
	str := `路多辛的, "所思所想"!`

	quoted := strconv.Quote(str)
	t.Error("Quoted: ", quoted)

	unquoted, err := strconv.Unquote(quoted)
	if err != nil {
		t.Error("Unquote error: ", err)
	} else {
		t.Error("Unquoted: ", unquoted)
	}
}