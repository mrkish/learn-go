package main

func Comp(array1 []int, array2 []int) bool {
    var same boolean
    for _, i; range array1 {
        var first := array1[i]
        for _, j; range array2 {
            second := array2[j]
            if first == second || first * first == second {
                 same = true
            } else {
                same = false
            }
        } 
    }
    return same
}

