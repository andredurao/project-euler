/* XOR decryption
Each character on a computer is assigned a unique code and the preferred standard is ASCII (American Standard Code for Information Interchange). For example, uppercase A = 65, asterisk (*) = 42, and lowercase k = 107.
A modern encryption method is to take a text file, convert the bytes to ASCII, then XOR each byte with a given value, taken from a secret key. The advantage with the XOR function is that using the same encryption key on the cipher text, restores the plain text; for example, 65 XOR 42 = 107, then 107 XOR 42 = 65.
For unbreakable encryption, the key is the same length as the plain text message, and the key is made up of random bytes. The user would keep the encrypted message and the encryption key in different locations, and without both "halves", it is impossible to decrypt the message.
Unfortunately, this method is impractical for most users, so the modified method is to use a password as a key. If the password is shorter than the message, which is likely, the key is repeated cyclically throughout the message. The balance for this method is using a sufficiently long password key for security, but short enough to be memorable.
Your task has been made easy, as the encryption key consists of three lower case characters. Using p059_cipher.txt (right click and 'Save Link/Target As...'), a file containing the encrypted ASCII codes, and the knowledge that the plain text must contain common English words, decrypt the message and find the sum of the ASCII values in the original text.
*/

package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var p = fmt.Println

func loadRunes() (list []rune) {
	txt, _ := ioutil.ReadFile("p059_cipher.txt")
	chars := strings.Split(string(txt), ",")
	for _, char := range chars {
		currentChar, _ := strconv.Atoi(char)
		list = append(list, rune(currentChar))
	}
	return
}

func passwd(a, b, c int) [3]rune {
	var newPassword [3]rune
	newPassword[0] = rune(a)
	newPassword[1] = rune(b)
	newPassword[2] = rune(c)
	return newPassword
}

// try to decrypt the first [size] chars or all if size = -1
func decrypt(encryptedText []rune, password [3]rune, size int) (result string) {
	if size == -1 {
		size = len(encryptedText)
	}
	for i := 0; i < size; i++ {
		currentChar := encryptedText[i] ^ password[i%3]
		result += string(currentChar)
	}
	return
}

// Analysis: Manually look for the word "the" among possible passwords
func analysis(encryptedText []rune) {
	for a := 97; a <= 122; a++ {
		for b := 97; b <= 122; b++ {
			for c := 97; c <= 122; c++ {
				password := passwd(a, b, c)
				test := decrypt(encryptedText, password, 50)
				if strings.Index(test, "the") > 0 {
					p(password, test)
				}
			}
		}
	}
	// Result: 101 120 112
}

func sum(encryptedText []rune) (result int) {
	password := [3]rune{101, 120, 112}
	decryptedText := decrypt(encryptedText, password, -1)
	runes := []rune(decryptedText)
	for _, value := range runes {
		result += int(value)
	}
	// p(decryptedText)
	return
}

func main() {
	p("Problem 59")
	p('a', 'z')
	encryptedText := loadRunes()
	// analysis(encryptedText) Result: 101 120 112
	result := sum(encryptedText)
	p(result)
}
