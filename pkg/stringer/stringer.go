package stringer

import (
    "strconv"
    "unicode/utf8"
)

func Reverse(input string) (result string) {
    for _, char := range input {
        result = string(char) + result
    }
    return result
}

func Length(input string, digits bool, runes bool) (count int, kind string) {
    switch {
    case digits:
      return CountNumbers(input), "digit"
    case runes:
      return utf8.RuneCountInString(input), "rune"
    default:
      return len(input), "char"
    }

if !digits && runes {
        return len(input), "char"
    }
    return CountNumbers(input), "digit"
}

func CountNumbers(input string) (count int) {
    for _, char := range input {
        _, err := strconv.Atoi(string(char))
        if err == nil {
            count++
        }
    }
    return count
}

