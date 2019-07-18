package main

import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
)

func dotest(array1 []int, array2 []int, exp bool) {
    var ans = Comp(array1, array2)
    Expect(ans).To(Equal(exp))
}

var _ = Describe("Tests Tour", func() {

    It("should handle basic cases", func() {
        var a1 = []int{121, 144, 19, 161, 19, 144, 19, 11}
        var a2 = []int{11*11, 121*121, 144*144, 19*19, 161*161, 19*19, 144*144, 19*19}
        dotest(a1, a2, true)
        a1 = []int{121, 144, 19, 161, 19, 144, 19, 11}
        a2 = []int{11*21, 121*121, 144*144, 19*19, 161*161, 19*19, 144*144, 19*19}
        dotest(a1, a2, false)
        a1 = nil
        a2 = []int{11*11, 121*121, 144*144, 19*19, 161*161, 19*19, 144*144, 19*19}
        dotest(a1, a2, false)
    })
})
