package tempconv_test

import (
	"fmt"
	"gobook/ch2/ex01/tempconv"
	"testing"
)

func TestCToF(t *testing.T) {
	if tempconv.CToF(tempconv.BoilingC) != 212 {
		t.Error(fmt.Sprint(tempconv.CToF(tempconv.BoilingC)))
	}
}

func TestFToC(t *testing.T) {
	if tempconv.FToC(212) != 100 {
		t.Error(fmt.Sprint(tempconv.FToC(212)))
	}
}

func TestCToK(t *testing.T) {
	if tempconv.CToK(tempconv.BoilingC) != 373.15 {
		t.Error(fmt.Sprint(tempconv.CToK(tempconv.BoilingC)))
	}
}

func TestKToC(t *testing.T) {
	if tempconv.KToC(373.15) != 100 {
		t.Error(fmt.Sprint(tempconv.KToC(373.15)))
	}
}

func TestFToK(t *testing.T) {
	if tempconv.FToK(212) != 373.15 {
		t.Error(fmt.Sprint(tempconv.FToK(212)))
	}
}

func TestKToF(t *testing.T) {
	if tempconv.KToF(373.15) != 212 {
		t.Error(fmt.Sprint(tempconv.KToF(373.15)))
	}
}
