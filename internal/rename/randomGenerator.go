package rename

import (
	"math/rand"
	"time"
)

func generateRandomString(length int, seedAppend int) string {
	password := ""

	characters := "qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM123456789!@#$%&"

	rand.Seed(time.Now().Unix() + int64(seedAppend))

	for i := 0; i < length; i++ {
		rnum := rand.Intn(len(characters))
		password += string(characters[rnum])
	}

	return password
}
