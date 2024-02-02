package utils

import (
	"crypto/rand"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/google/uuid"
)

func PublicIP() string {
	resp, err := http.Get("https://ifconfig.co/ip")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	ip, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("Public IP address:", string(ip))
	return string(ip)
}

func GetUUID() string {
	uuid := uuid.New()
	id := uuid.ID()

	return strconv.Itoa(int(id))
}

func RandomString(n int) string {
	s, r := make([]rune, n), []rune(alphaNumRunes)

	for i := range s {
		p, _ := rand.Prime(rand.Reader, len(r))
		x, y := p.Uint64(), uint64(len(r))
		s[i] = r[x%y]
	}

	return string(s)
}
