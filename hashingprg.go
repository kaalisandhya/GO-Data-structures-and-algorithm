1.package main

import (
        "fmt"
        "hash/fnv"
)

func hash(s string) uint32 {
        h := fnv.New32a()
        h.Write([]byte(s))
        return h.Sum32()
}

func main() {
        fmt.Println(hash("Hello hashing"))
        fmt.Println(hash("Hello hashing"))
}
304492899
304492899

Program exited.

2.package main

import (
    "fmt"

    "golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

func main() {
    password := "secret"
    hash, _ := HashPassword(password) // ignore error for the sake of simplicity

    fmt.Println("Password:", password)
    fmt.Println("Hash:    ", hash)

    match := CheckPasswordHash(password, hash)
    fmt.Println("Match:   ", match)
}
Password: secret
Hash:     $2a$14$WuL1zZ6wgn5h0iGrbJ2mb.WdtJOB0zH9Jr.B17kMDewS4E306RYjO
Match:    true

Program exited.
3.package main

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"io"
)

func main() {
	input := "Lorem Ipsum dolor sit Amet"
	md5 := md5.New()
	sha_256 := sha256.New()
	sha_512 := sha512.New()
	io.WriteString(md5, input)
	sha_256.Write([]byte(input))
	sha_512.Write([]byte(input))
	sha_512_256 := sha512.Sum512_256([]byte(input))
	hmac512 := hmac.New(sha512.New, []byte("secret"))
	hmac512.Write([]byte(input))

	
	fmt.Printf("md5:\t\t%x\n", md5.Sum(nil))

	
	fmt.Printf("sha256:\t\t%x\n", sha_256.Sum(nil))

	
	fmt.Printf("sha512:\t\t%x\n", sha_512.Sum(nil))

	
	fmt.Printf("sha512_256:\t%x\n", sha_512_256)

	
	fmt.Printf("hmac512:\t%s\n", base64.StdEncoding.EncodeToString(hmac512.Sum(nil)))
}

