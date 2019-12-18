package dayone

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
)

const blockSize = 64
const size = 16

//ExecuteDayFour ...
func ExecuteDayFour() {
	fmt.Printf("The Trailing Decimals are: %v", findInput("yzbqklnj"))
}

func findInput(input string) (trailingDecimal int) {
	for i := 0; i < i+1; i++ {
		hash := generateMD5Hash(input + strconv.Itoa(i))
		if hex.EncodeToString(hash)[0:6] == "000000" {
			trailingDecimal = i
			break
		}
	}

	return
}

func generateMD5Hash(input string) []byte {
	hasher := md5.New()
	hasher.Write([]byte(input))
	return hasher.Sum(nil)
}
