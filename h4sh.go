// h4sh by farinap5
// A utility to create hash signatures from strings.

package main

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"crypto/sha1"
	"fmt"
	"io"
	"os"
)

var comm string
var e string

func _sha_1(comm string) {
	sha_1 := sha1.New()
	sha_1.Write([]byte(comm))
	print("SHA-1 <- ", comm, "\n")
	fmt.Printf("sha1: %x\n", sha_1.Sum(nil))
}
func _md5(comm string) {
	md5 := md5.New()
	io.WriteString(md5, comm)
	print("MD5 <- ", comm, "\n")
	fmt.Printf("md5: %x\n", md5.Sum(nil))
}
func _sha_256(comm string) {
	sha_256 := sha256.New()
	sha_256.Write([]byte(comm))
	print("SHA-256 <- ", comm, "\n")
	fmt.Printf("sha256: %x\n", sha_256.Sum(nil))
}
func _sha_512(comm string)  {
	sha_512 := sha512.New()
	sha_512.Write([]byte(comm))
	print("SHA-512 <- ", comm, "\n")
	fmt.Printf("sha512: %x\n", sha_512.Sum(nil))
}
func main() {
	e = `
     -H4sh-
  Usage Method:
> md5 "string"
> sha1 "string"
> sha256 "string"
> sha512 "string"

`
	if len(os.Args) > 2 {
		if os.Args[1] == "md5" {
			comm = os.Args[2]
			_md5(comm)
		}
		if os.Args[1] == "sha256" {
			comm = os.Args[2]
			_sha_256(comm)
		}
		if os.Args[1] == "sha512" {
			comm = os.Args[2]
			_sha_512(comm)
		}
		if os.Args[1] == "sha1" {
		comm = os.Args[2]
		_sha_1(comm)
		}
	} else {
		print(e)
	}
}
