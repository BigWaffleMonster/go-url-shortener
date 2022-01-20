package helperfunctions

import (
  "math/rand"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_1234567890"

func GenerateRandomString() string {
    b := make([]byte, 10)
    for i := range b {
        b[i] = letterBytes[rand.Intn(len(letterBytes))]
    }
    return string(b)
}
