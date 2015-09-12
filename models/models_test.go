package models

import (
	"testing"

	"github.com/Unknwon/com"
	. "github.com/smartystreets/goconvey/convey"
)

func int64sToString(nums []int64) string {
	if len(nums) == 0 {
		return ""
	}

	str := ""
	for i := 0; i < len(nums); i++ {
		str += com.ToStr(nums[i]) + " "
	}
	return str[:len(str)-1] // Remove last white space
}

func Test_Fibonacci(t *testing.T) {
	Convey("Generate Fibonacci sequence", t, func() {
		Convey("Invalid parameters", func() {
			So(FibonacciSequence(-1), ShouldBeNil)
			So(FibonacciSequence(150), ShouldBeNil)
		})

		Convey("Valid parameters", func() {
			So(int64sToString(FibonacciSequence(0)), ShouldEqual, "")
			So(int64sToString(FibonacciSequence(1)), ShouldEqual, "0")
			So(int64sToString(FibonacciSequence(2)), ShouldEqual, "0 1")
			So(int64sToString(FibonacciSequence(3)), ShouldEqual, "0 1 1")
			So(int64sToString(FibonacciSequence(4)), ShouldEqual, "0 1 1 2")
			So(int64sToString(FibonacciSequence(5)), ShouldEqual, "0 1 1 2 3")
			So(int64sToString(FibonacciSequence(6)), ShouldEqual, "0 1 1 2 3 5")
			So(int64sToString(FibonacciSequence(10)), ShouldEqual, "0 1 1 2 3 5 8 13 21 34")
		})
	})
}
