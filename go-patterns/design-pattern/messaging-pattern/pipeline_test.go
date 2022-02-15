package messaging_pattern

import (
	"testing"
)

func TestPipeline(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	_ = pipeline(nums, echo, square, echo, sum, echo)
}

func TestSample(t *testing.T) {
	t.Logf("log")
	t.Errorf("error")
}
