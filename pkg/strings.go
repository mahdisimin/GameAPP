package pkg

import (
	"crypto/md5"
	"encoding/hex"
)

func HashTextFunc(text string) string {
	plainPass := text
	hashPass := md5.Sum([]byte(plainPass))
	hashText := hex.EncodeToString(hashPass[:])
	return hashText
}
