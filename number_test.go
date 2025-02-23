package number_to_vietnamese_text

import (
	"testing"
)

func TestToVietnameseWords(t *testing.T) {
	tests := []struct {
		input    int64
		expected string
	}{
		{0, "không"},
		{5, "năm"},
		{15, "mười lăm"},
		{123, "một trăm hai mươi ba"},
		{1000, "một nghìn"},
		{1000000, "một triệu"},
		{2000000, "hai triệu"},
		{-55, "âm năm mươi lăm"},
	}

	for _, tt := range tests {
		result := ToVietnameseWords(tt.input)
		if result != tt.expected {
			t.Errorf("For input %d, expected %s but got %s", tt.input, tt.expected, result)
		}
	}
}

func TestNumberToVietnameseWord_Convert(t *testing.T) {
	tests := []struct {
		input    NumberToVietnameseWord
		expected string
	}{
		{NumberToVietnameseWord{123, NumberToVietnameseWordOption{false, VND, false}}, "một trăm hai mươi ba VND"},
		{NumberToVietnameseWord{5000, NumberToVietnameseWordOption{true, VNDFull, true}}, "Năm nghìn việt nam đồng"},
		{NumberToVietnameseWord{1000000, NumberToVietnameseWordOption{false, USD, false}}, "một triệu USD"},
	}

	for _, tt := range tests {
		result := tt.input.Convert()
		if result != tt.expected {
			t.Errorf("For input %+v, expected %s but got %s", tt.input, tt.expected, result)
		}
	}
}
