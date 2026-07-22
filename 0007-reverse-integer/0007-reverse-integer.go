package main

func reverse(x int) int {
    const IntMax = 1<<31 - 1
    const IntMin = -1 << 31

    rev := 0
    for x != 0 {
        pop := x % 10
        x /= 10
        if rev > IntMax/10 || (rev == IntMax/10 && pop > 7) {
            return 0
        }
        if rev < IntMin/10 || (rev == IntMin/10 && pop < -8) {
            return 0
        }
        rev = rev*10 + pop
    }
    return rev
}