# rune

[![License][License-Image]][License-Url]
[License-Url]: https://opensource.org/licenses/Apache-2.0
[License-Image]: https://img.shields.io/badge/license-Apache%202.0-blue.svg?maxAge=2592000

[![Build][Build-Status-Image]][Build-Status-Url]
[Build-Status-Url]: http://travis-ci.org/mikijov/rune
[Build-Status-Image]: https://travis-ci.org/mikijov/rune.svg?branch=master

[![ReportCard][ReportCard-Image]][ReportCard-Url]
[ReportCard-Url]: https://goreportcard.com/report/github.com/mikijov/rune
[ReportCard-Image]: https://goreportcard.com/badge/github.com/mikijov/rune

Programming language I use to experiment with syntax and programming patterns.

## Backlog

### v1

- [ ] tests
- [ ] arrays and slices
- [ ] error handling
- [ ] break and continue

### Future

- [ ] structs
- [ ] imports/modules
- [ ] allow passing Environment to user functions
- [ ] implement variadic functions
- [ ] support for variadic user functions
- [ ] maps

## Syntax

```
console := import("console")

fibbonaci :func(count :int) :int = {
fibbonaci := func(count :int) :int {
func fibbonaci(count :int) :int {
    n2 := 1;
    n1 := 1;
    count := count - 2;
    loop while count > 0 {
        count = count - 1;
        temp := n2 + n1;
        n2 = n1;
        n1 = temp;
    }
    return n1;
}

fibbonaci :func(count :int) :int = { n2 := 1; n1 := 1; count := count - 2; loop while count > 0 { count = count - 1; temp := n2 + n1; n2 = n1; n1 = temp; } return n1; }

console.print("Enter value for n:");
count := console.input();
console.print("fibbonaci({count})={fibbonaci(count)}");

max := func(a, b :int) :int { if a > b a else b }

func quicksort(A :list, lo, hi :int) {
    if lo < hi {
        p := partition(A, lo, hi);
        quicksort(A, lo, p - 1);
        quicksort(A, p + 1, hi);
    }
}

partition := func(A :list, lo, hi :int) :int {
    pivot := A[hi];
    i := lo - 1;

    j := lo;
    loop while j < hi {
        if A[j] < pivot then {
            i = i + 1;
            temp := A[i], A[i] = A[j], A[j] = temp;
        }
    }
    if A[hi] < A[i + 1] {
        temp := A[i + 1], A[i + 1] = A[hi], A[hi] = temp;
    }
    return i + 1;
}

func newCustomer(adr :string) :map {
    return {
        name: "Default",;
        address: if (not adr empty) adr else "N/A";
    }
}

func someFibNumbers() :int[] {
    return [
        fibbonaci(1),
        fibbonaci(2),
        fibbonaci(3),
        fibbonaci(4),
        fibbonaci(5),
    ]
}
```

<!-- vim: set expandtab ts=4 sw=4 sts=4 tw=80 : -->
