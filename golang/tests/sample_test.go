package internal_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"isuct.ru/informatics2022/internal"
)

func TestSumm(t *testing.T) {
	summ := internal.Summ(2, 3)
	if summ != 5 {
		t.Fatalf(`Summ(2,3) = %d, want 5, error`, summ)
	}
}
func TestFu(t *testing.T) {
	resol := internal.FuncResolution(1)
	assert.InDelta(t, 1.1066819197003217, resol, 0.02)
}
