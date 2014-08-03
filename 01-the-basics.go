package main

import (
    "fmt"
    "math"
    "math/rand"
    "strings"
    "time"
    "unicode/utf8"
)

const (
   A  = 10
   B  = "hello"
   C  = 10.2
)

func init() { // I'm always called before main()
    rand.Seed(time.Now().UTC().UnixNano())
}

func main() {
    // --------------------
    // http://blog.golang.org/gos-declaration-syntax
    // --------------------
    var a1 int; /* <- note ; */ a1 = 10
    var a2 int = 5
    var a3 = 5
    a4 := 5

    _ = a1; _ = a2; _ = a3; _ = a4

    // Strings
    s1 := "hello world"
    s2 := "ok" + B
    _ = s1; _ = s2

    // Maps
    monthDays := map[string]int{
		"Jan": 31, "Feb": 28, "Mar": 31,
		"Apr": 30, "May": 31, "Jun": 30,
		"Jul": 31, "Aug": 31, "Sep": 30,
		"Oct": 31, "Nov": 30, "Dec": 31,
	}
    _ = monthDays

    // --------------------
    // Arrays, see Slices below
    // --------------------
    // ! think of arrays as a struct w/ index fields instead of name
    // Arrays pass as *values*
    b0 := [3]int{0, 1, 2}
    b1 := [...]int{0, 1, 2} // have compiler count elements
    var b2 [3]int // [0,0,0]
    b3 := [2][2]int{ {1, 2}, {3, 4} }
    _ = b0; _ = b1; _ = b2; _ = b3;


    // --------------------
    // CONDITIONALS
    // --------------------
    if (false && true) || (true && false) {
        // ...
    } else if false || false {
        // ...
    } else {
        // ...
    }

    // http://golang.org/doc/effective_go.html#switch
    // Form 1: Matching values
    sw := 5
    switch (sw) {
    case 1,3,5:
        // ... it's odd!
    case 2,4,6:
        // ... it's even!
    default:
        // 0 || > 6!
    }

    // Form 2: check bunch of boolean expressions
    switch {
    case (false && true) || (true && false):
        // ... no break required!
    case ("a" == "A") || (a1 > 100):
        // ...
    default:
        // ...
    }

    // --------------------
    // LOOPS, there is only `for`
    // --------------------
    for i:=0; i < 4; i++ { _ = i }
    // no while/do-while loops, all unified under for
    for { /* like for(;;) */
        break
    }

    // using range: https://code.google.com/p/go-wiki/wiki/Range
    for i := range b1 { _ = i}
    for i := range [...]int{1,2,3} { _ = i }
    for k,v := range monthDays { _ = k; _ = v }
    for _,v := range monthDays { _ = v }
    for k,_ := range monthDays { _ = k }

    l1 := 0
    for (l1 != 3) { l1 = rand.Intn(5) }  // like while loop

    // -=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
    // Common Stuff I do..
    // -=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-

    // --------------------
    // Math http://golang.org/pkg/math
    // --------------------
    _ = 1 + (4 * 9 / 8 * (13 % 5))
    _ = math.Pow(5,2) * math.Pi
    _ = math.Floor(10.2)
    _ = math.Ceil(10.2)
    // no Round() https://code.google.com/p/go/issues/detail?id=4594
    _ = math.Abs(-14.12)

    // --------------------
    // Strings http://golang.org/pkg/strings/
    // --------------------
    s := "   Hello 世界. I'm utf8! "
    st := strings.Trim(s, " ")

    _ = strings.Split(s, ".")
    _ = strings.ToLower(s)
    _ = strings.ToUpper(s)
    _ = strings.Replace(s, "世界", "World", -1 /* num to replace, <0 = all */)
    _ = strings.Contains("世", s) // true!
    _ = (len(st) != utf8.RuneCountInString(st)) // utf8!

    // --------------------
    // Random Numbers http://golang.org/pkg/math/rand/
    // --------------------
    rN := 5
    _ = rand.Intn(rN)  // [0, rN)
    _ = rand.Int()     // [0, 2^31)
    _ = rand.Int63()   // [0, 2^63)
    _ = rand.Float32() // [0, 1.0)
    _ = rand.Float64() // [0, 1.0)

    fmt.Print("Done\n")
}
