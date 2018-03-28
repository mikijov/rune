# rune

[![License][License-Image]][License-Url]
[![Build][Build-Status-Image]][Build-Status-Url]
[![ReportCard][ReportCard-Image]][ReportCard-Url]

Programming language I use to experiment with syntax and programming patterns.
Still very much in early stages. However if you find any of it interesting I
welcome you stopping by and asking a question (make an issue for now)...

## Backlog

### v1

- [x] tests
- [x] structs
- [ ] arrays
  - [x] simple array syntax
  - [ ] allocation
  - [ ] slices
- [ ] interfaces
  - [ ] method declaration and invocation
  - [ ] interface declaration
  - [ ] interface conversion/assignment
- [ ] error handling
- [ ] break and continue
- [ ] string type

### Future

- [ ] struct literals
- [ ] custom zero value for types
- [ ] imports/modules
- [ ] allow passing Environment to user functions
- [ ] implement variadic functions
- [ ] support for variadic user functions
- [ ] maps

## Syntax

### Type definition:
```
type Integer :int

type DemoStruct :struct {
    field1 :int
    field2 :real
    field3 :bool
}
```

#### Custom zero value.
```
# -1 is the zero value
type weirdZero :int = -1
```

### Arrays.
```
type myArray :int[]

a1 :int[] = [1, 2]

a2 := [1, 2, 3] # implicit type int[]

a3 := int[33] # allocate array of 33 integers, initilized to zero

v := a2[1] # second item in zero based array, i.e. '2' from a2
```

### Objects and Interfaces
```
struct S1 {
    f1 :int
}

func S1.setF1(value :int) {
    this.f1 = value
}

func S1.getF1() :int {
    return this.f1

func S1.add(value :int) {
    this.f1 += value
}

# Struct and field definition automatically defines and implements interface.
# Note that the field name is derived from setter and getter, not from the
# field name. I.e. the field could have had different name but the same
# interface would be created.
interface S1 {
    f1 :int
    add(value :int)
}

S1 implements S1 {
    f1.get = getF1
    f1.set = setF1
    add = add
}

interface I1 {
    f1 :int
    add(value :int)
}

# Automatically maps S1 methods to I1. Methods are mapped by name. Field setters
# and getters are mapped with (get|set)(capitalized field name).
S1 implements I1;

interface I2 {
    field1 :int # virtual field, but has to be mapped to getter and setter
    anotherAdd(value :int)
}

# Manual implementation of interface where methods are assigned to methods and
# properties.
S1 implements I2 {
    field1.get = getF1
    field1.set = setF1
    anotherAdd = myAdd
}
```

### Sample code.
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

[License-Url]: https://opensource.org/licenses/Apache-2.0
[License-Image]: https://img.shields.io/badge/license-Apache%202.0-blue.svg?maxAge=2592000
[Build-Status-Url]: http://travis-ci.org/mikijov/rune
[Build-Status-Image]: https://travis-ci.org/mikijov/rune.svg?branch=master
[ReportCard-Url]: https://goreportcard.com/report/github.com/mikijov/rune
[ReportCard-Image]: https://goreportcard.com/badge/github.com/mikijov/rune
<!-- vim: set expandtab ts=4 sw=4 sts=4 tw=80 : -->
