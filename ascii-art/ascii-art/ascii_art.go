package ascii_art

import (
	"crypto/md5"
	"encoding/hex"
	"net/http"
	"os"
	"strings"
)

const (
	b1 = "a51f800619146db0c42d26db3114c99f" // standard
	b2 = "d44671e556d138171774efbababfc135" // shadow
	b3 = "0021f26ad06f2f73a0cfa7b7d38d1434" // thinkertoy
)

func MainAsciiArt(s string, bannerName string) (string, int) {
	var res string
	path := "./banners/" + bannerName + ".txt"
	banner, err := os.ReadFile(path)
	if err != nil {
		return "", http.StatusBadRequest
	}
	switch bannerName {
	case "standard":
		if CheckHash(string(banner)) != b1 {
			return "", http.StatusInternalServerError
		}
	case "shadow":
		if CheckHash(string(banner)) != b2 {
			return "", http.StatusInternalServerError
		}
	case "thinkertoy":
		if CheckHash(string(banner)) != b3 {
			return "", http.StatusInternalServerError
		}
	default:
		return "", http.StatusBadRequest
	}
	symbols := strings.Split(strings.ReplaceAll(string(banner), "\r", ""), "\n\n")
	for _, word := range strings.Split(s, "\r\n") {
		for lines := 0; lines < 8; lines++ {
			for _, v := range word {
				res += strings.Split(symbols[v-32], "\n")[lines]
			}
			res += "\n"
		}
	}
	return res, 200
}

func CheckHash(s string) string {
	hash := md5.New()
	hash.Write([]byte(s))
	return hex.EncodeToString(hash.Sum(nil))
}
