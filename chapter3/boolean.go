package boolean

// btoi return `0` if b is false, and `1` otherwise
func btoi(b bool) int {
    if b {
        return 1
    }

    return 0
}

// itob return false if `i` is 0, and true otherwise
func itob(i int) bool {
    return i != 0
}

