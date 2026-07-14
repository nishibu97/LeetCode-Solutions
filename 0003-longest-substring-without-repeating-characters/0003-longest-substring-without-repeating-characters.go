package main

func lengthOfLongestSubstring(s string) int {
    lastSeen := make(map[byte]int)
    left := 0
    maxLength := 0

    for right := 0; right < len(s); right++ {
        char := s[right]

        if index, found := lastSeen[char]; found && index >= left {
            left = index + 1
        }

        lastSeen[char] = right

        length := right - left + 1
        if length > maxLength {
         maxLength = length
        }
    }

    return maxLength
}