package util

import (
    "crypto/rand"
    "math/big"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func GenerateShortCode(length int) (string, error) {
    b := make([]rune, length)
    for i := 0; i < length; i++ {
        n, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
        if err != nil {
            return "", err
        }
        b[i] = letters[n.Int64()]
    }
    return string(b), nil
}