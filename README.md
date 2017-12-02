# rune

Simple declarations working. Implement assignments and tests.




```
console := import("console")

func fibbonaci(count :int) :int
    n2 := 1
    n1 := 1
    count := count - 2
    loop while count > 0
        count = count - 1
        temp := n2 + n1
        n2 = n1
        n1 = temp
    return n1

fibbonaci :func(count :int) :int = n2 := 1, n1 := 1, count := count - 2, loop while count > 0 (count = count - 1, temp := n2 + n1, n2 = n1, n1 = temp) return n1)

console.print("Enter value for n:")
count := console.input()
console.print("fibbonaci({count})={fibbonaci(count)}")

max := func(a, b :int) :int if a > b a else b

func quicksort(A :list, lo, hi :int)
    if lo < hi
        p := partition(A, lo, hi)
        quicksort(A, lo, p - 1)
        quicksort(A, p + 1, hi)

partition := func(A :list, lo, hi :int) :int
    pivot := A[hi]
    i := lo - 1

    j := lo
    loop while j < hi
        if A[j] < pivot then
            i = i + 1
            temp := A[i], A[i] = A[j], A[j] = temp
    if A[hi] < A[i + 1]
        temp := A[i + 1], A[i + 1] = A[hi], A[hi] = temp
    return i + 1

func newCustomer(adr :string) :map
    return {
        name: "Default",
        address: if (not adr empty) adr else "N/A",
    }

func someFibNumbers() :int[]
    return [
        fibbonaci(1),
        fibbonaci(2),
        fibbonaci(3),
        fibbonaci(4),
        fibbonaci(5),
    ]


```

# crune

```
console := import("console")

fibbonaci := func(count :int) :int {
    int n2 = 1;
    int n1 = 1;
    for index := count - 2; count > 0; ++count {
        int temp = n2 + n1;
        n2 = n1;
        n1 = temp;
    }
    return n1;
}

printf("Enter value for n:");
count := console.input();
printf("fibbonaci(%d)=%d", count, fibbonaci(count));

max := func(a, b :int) :int { if a > b a else b }

quicksort := func(A :array, lo, hi :int) {
    if lo < hi {
        int p = partition(A, lo, hi);
        quicksort(A, lo, p - 1);
        quicksort(A, p + 1, hi);
    }
}

partition := func(A :array, lo, hi :int) :int {
    pivot := A[hi];
    i := lo - 1;

    while j < hi {
        if A[j] < pivot {
            i = i + 1;
            temp := A[i]; A[i] = A[j]; A[j] = temp;
        }
    }
    if A[hi] < A[i + 1] {
        temp := A[i + 1]; A[i + 1] = A[hi]; A[hi] = temp;
    }
    return i + 1;
}

fibbonaci := func(count :int) :int {n2 := 1; n1 := 1; count := count - 2; while count > 0 { count = count - 1; temp := n2 + n1; n2 = n1; n1 = temp;} return n1;}

```


```
finFunc :func(count :int) :int
type finFunc :func(count :int) :int
```

<!-- vim: set expandtab ts=4 sw=4 sts=4 tw=80 : -->
