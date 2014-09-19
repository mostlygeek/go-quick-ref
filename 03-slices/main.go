package main

func main() {
    // --------------------
    // SLICES http://blog.golang.org/go-slices-usage-and-internals
    // A slice is a descriptor of an array segment. It consists of a pointer to the array,
    // the length of the segment, and its capacity
    //
    // length    - number of elements of the array referred to by the slice
    // capacity  - number of elements in the underlying array
    // --------------------
    sl0 := []int{0,1,2,3,4}     // notice empty []
    sl1 := make([]int, 3, 5)    // 3=length  5=capacity
    _ = len(sl1) // 3
    _ = cap(sl1) // 5

    // creating a slice from another slice
    sl3 := sl0[1:4] // [1 2 3]
    sl4 := sl0[:4]  // [0 1 2 3]
    sl5 := sl0[:]   // [0 1 2 3 4]

    // run time panics: only use range between [0, len(s10)]
    sl6 := sl0[0:len(sl0)]

    _ = sl3; _ = sl4; _ = sl5; _ = sl6;

    // note: index values outside of bounds = runtime panic, sl4 := sl0[1,len(sl0)+1]
    source := []int{9,10,11,12,13}
    destination1 := make([]int, len(source), (cap(source)+1)*2)
    destination2 := make([]int, (cap(source)+1)*2)
    copy(destination1, source) // [9 10 11 12 13]           - len(destination1) == 5
    copy(destination2, source) // [9 10 11 12 13 0 0 0 0 0] - len(destination2) == 10
}
