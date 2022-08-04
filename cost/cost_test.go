package cost

import (
	"testing"
	"time"
)

func TestBegin(t *testing.T) {

}

func TestPrint(t *testing.T) {

	Begin("camera-001")
	Printf("hello world")

	time.Sleep(time.Millisecond * 56)
	Printf("welcome to Beijing")
	End()
}
