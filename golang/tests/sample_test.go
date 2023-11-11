package internal_test

import (
	"testing"

	"isuct.ru/informatics2022/internal"
)

func TestSumm(t *testing.T) {
	summ := internal.Summ(2, 3)
	if summ != 5 {
		t.Fatalf(`Summ(2,3) = %d, want 5, error`, summ)
	}
}
func Test_fu(t *testing.T) {
	resol := internal.Func_resolution(1)
	if resol != 1.1066819197003217 {
		t.Fatalf(` Func_resolution(1) = %v, want 1.1066819197003217, error`, resol)
	}
}
