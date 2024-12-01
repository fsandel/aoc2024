package day1

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestDay1(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Day1 Suite")
}

var _ = Describe("Day1", func() {
	left := []int{3, 4, 2, 1, 3, 3}
	right := []int{4, 3, 5, 3, 9, 3}

	Context("calcSimilarity", func() {
		It("should get the output from the example", func() {
			Expect(calcSimilarity(left, right)).To(Equal(31))
		})
	})

	Context("calcListDiff", func() {
		It("should get the output from the example", func() {
			Expect(calcListDiff(left, right)).To(Equal(11))
		})
	})
})
