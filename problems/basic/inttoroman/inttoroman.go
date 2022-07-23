package inttoroman

var (
	M = []string{"", "M", "MM", "MMM"}
	C = []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	X = []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	I = []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
)

func IntToRoman(num int) string {
	return M[num/1000] + C[(num%1000)/100] + X[(num%100)/10] + I[num%10]
}
