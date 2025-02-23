package number_to_vietnamese_text

import (
	"fmt"
	"regexp"
	"strings"
)

var digits = []string{
	"không", "một", "hai", "ba", "bốn", "năm", "sáu", "bảy", "tám", "chín",
}

var multipleThousand = []string{
	"", "nghìn", "triệu", "tỷ", "nghìn tỷ", "triệu tỷ", "tỷ tỷ",
}

type FormatMoney string

const (
	VNDFull  FormatMoney = "VNDFull"
	VNDShort FormatMoney = "VNDShort"
	VND                  = "VND"
	USD      FormatMoney = "USD"
)

const (
	VNDFullWord  string = " việt nam đồng"
	VNDWord      string = " VND"
	VNDShortWord string = " đồng"
	USDWord      string = " USD"
)

type NumberToVietnameseWordOption struct {
	money       bool
	formatMoney FormatMoney
	ucFirst     bool
}

type NumberToVietnameseWord struct {
	number  int64
	options NumberToVietnameseWordOption
}

func (n *NumberToVietnameseWord) Convert() string {
	words := ToVietnameseWords(n.number)
	if n.options.ucFirst {
		words = capitalize(words)
	}

	if n.options.formatMoney == VND {
		return words + VNDWord
	}
	if n.options.formatMoney == VNDFull {
		return words + VNDFullWord
	}
	if n.options.formatMoney == VNDShort {
		return words + VNDShortWord
	}
	if n.options.formatMoney == USD {
		return words + USDWord
	}

	return words
}

func capitalize(s string) string {
	if len(s) == 0 {
		return ""
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

func readPair(b, c int) string {
	switch b {
	case 0:
		if c == 0 {
			return ""
		}
		return " lẻ " + digits[c]
	case 1:
		if c == 0 {
			return "mười"
		} else if c == 5 {
			return "mười lăm"
		}
		return "mười " + digits[c]
	default:
		if c == 0 {
			return digits[b] + " mươi"
		} else if c == 1 {
			return digits[b] + " mươi mốt"
		} else if c == 4 {
			return digits[b] + " mươi tư"
		} else if c == 5 {
			return digits[b] + " mươi lăm"
		}
		return digits[b] + " mươi " + digits[c]
	}
}

func readTriple(triple string, showZeroHundred bool) string {
	a, b, c := int(triple[0]-'0'), int(triple[1]-'0'), int(triple[2]-'0')

	switch {
	case a == 0 && b == 0 && c == 0:
		return ""
	case a == 0 && showZeroHundred:
		return "không trăm " + readPair(b, c)
	case a == 0 && b == 0:
		return digits[c]
	case a == 0:
		return readPair(b, c)
	default:
		return digits[a] + " trăm " + readPair(b, c)
	}
}

func ToVietnameseWords(number int64) string {
	if number == 0 {
		return "không"
	}

	if number < 0 {
		return "âm " + ToVietnameseWords(-number)
	}

	s := fmt.Sprintf("%d", number)
	padding := (3 - len(s)%3) % 3
	s = strings.Repeat("0", padding) + s

	var groups []string
	for i := 0; i < len(s); i += 3 {
		groups = append(groups, s[i:i+3])
	}

	showZeroHundred := false
	for _, g := range groups {
		if g != "000" {
			showZeroHundred = true
			break
		}
	}

	var result string
	for i, g := range groups {
		idx := len(groups) - 1 - i
		part := readTriple(g, showZeroHundred && i > 0)
		if part != "" {
			result += part + " " + multipleThousand[idx] + " "
		}
	}

	re := regexp.MustCompile(`\s+`)
	return strings.TrimSpace(re.ReplaceAllString(result, " "))
}
