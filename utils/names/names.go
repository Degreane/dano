// names generates a name for given integer
package names

import (
	"fmt"
	"math"

	"gonum.org/v1/gonum/stat/combin"
)

var (
	nameArr []rune = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

func claculateElementRangeK(k int) int {
	for i := 0; i < len(nameArr); i++ {
		bin := combin.Binomial(len(nameArr), i)
		if k < bin {
			return i
		}
	}
	return 0
}
func getTestName(in int) {
	// get test name does just resolve to the boundaries of the integer

	for i := 0; i < in; i++ {
		cs := combin.Combinations(len(nameArr), i)
		bin := claculateElementRangeK(in)
		fmt.Printf("With I = %d Boundaries=%d, Binomial=%d\n", i, len(cs), bin)
	}
	fmt.Printf("Calculating Binomial of 27 %d\n", claculateElementRangeK(27))
	cs := combin.Combinations(len(nameArr), claculateElementRangeK(27))
	for i := 0; i <= 27; i++ {
		fmt.Printf("I=%d %c%c\n", i, nameArr[cs[i][0]], nameArr[cs[i][1]])
	}

}
func _calculateBoundaries(in int) (int, int, int) {
	// calculateBoundaries returns min,max boundaries for the given input int
	_len := len(nameArr)
	var _min, _max, _radix int = 0, 0, 0
	// fmt.Println("{IN<", in, ">}")
	if in > 0 && in <= _len {
		return 1, _len, 1
	}
	for i := 1; ; i++ {
		if in <= _max {
			return _min, _max, _radix
		} else {
			_min = _max + 1
			_max = _max + int(math.Pow(float64(_len), float64(i)))
			_radix = int(math.Pow(float64(_len), float64(i-1)))
		}
	}
}
func GetName(in int, prefix string) string {

	_len := len(nameArr)
	if in <= 0 {
		return ""
	}
	if in > 0 && in <= _len {
		return fmt.Sprintf("%s%c", prefix, nameArr[in-1])
	}
	_min, _max, _radix := _calculateBoundaries(in)

	var _str string = fmt.Sprint(prefix)
	var _sub int

	if _min <= in && in <= _max {

		_sub = (in - _radix)

		var i int = 1
		for i = 1; _sub > _min-1; i++ {
			_sub = _sub - (_min - 1)

		}

		_pref := fmt.Sprintf("%s%c", prefix, nameArr[i-1])
		return GetName(_sub, _pref)
	}
	return _str
}
