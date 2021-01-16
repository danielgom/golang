package main

import (
	"fmt"
	"strconv"
	"unicode/utf8"
	"unsafe"
)

func main() {

	str := "Daлiel😱🙃"
	fmt.Println(len(str))
	fmt.Println(utf8.RuneCountInString(str))

	// Can convert string to byte slice or rune slice but these are not good ideas
	// Golang makes it easy so you do not convert these to rune slices or byte slices and viceversa
	// this has an explanation:

	// Strings are inmutable, so we cannot convert something here like str[1] = 'h', we can convert everything to
	// byte slice and change it however it is going to give errors
	bytes := []byte(str)
	fmt.Println(bytes)
	bytes[2] = 'N'

	str = string(bytes)
	fmt.Println(str) // This is printed DaN�iel😱🙃 , notice the wrong encoding, this is because there are problems
	// while changing 2 bytes literals to 1 byte or any ammount of bytes

	// The main problem with converting to rune slices is memory, each
	str = "Daлiel😱🙃"
	runes := []rune(str)
	fmt.Println(int(unsafe.Sizeof(runes[0])) * len(runes)) // 32 Bytes
	fmt.Println(len(runes))                                // Same 8 runes

	// As we have seen so far string is UTF-8 Encoded automatically so it uses from 1 byte to 4 bytes per rune
	// If we convert it to a rune slice []rune, every location is 4 bytes so we waste memory

	// DECODING A STRING

	text := `Insectele reprezintă o clasă de animale nevertebrate hexapode aparținând încrengăturii Arthropoda, subîncrengătura Hexapoda. În natură există circa un milion de specii de insecte, care viețuiesc atât în mediile terestre cât și în cele acvatice. Insectele pot fi apterigote (fără aripi) sau pterigote (cu aripi).

Cea mai largă și uniform răspândită categorie taxonomică din cadrul artropodelor, insectele alcătuiesc grupul cel mai divers de animale de pe Pământ, arealul fiind cu precădere uscatul, existând cu mult peste 1.000.000 de specii, dintre care aproximativ 925.000 au fost cercetate și descrise.`
	// if we try to print byte by byte it wont work
	//Converting to byte wont work because we would be just expanding the size but the value is the same
	for i := 0; i < len(text); {
		r, size := utf8.DecodeRuneInString(text[i:]) // This returns the decoded rune plus the size of this one
		fmt.Printf("%c", r)                          // Print the next rune

		i += size // Add the size so for example in "în" the size of î is 2 so next rune starts in 3
	}
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()

	// FASTER WAY **************************************************

	// This next code creates a new array so it is not so recomended

	/*
		for _, v := range text {
			fmt.Printf("%c", v)
		}
	*/

	// Efficient, this does not create a new array and it does the same as the code above

	for i := range text {
		fmt.Printf("%c", text[i])
	}

	s := strconv.FormatFloat(12.445, 'E', -1, 32)
	fmt.Println(s)

	f, _ := strconv.ParseFloat("12.445", 2)
	fmt.Println(f)

	strings := []string{
		"cool",
		"güzel",
		"jīntiān",
		"今天",
		"read 🤓",
	}

	for _, s := range strings {
		fmt.Printf("%q\n", s)

		// Print the byte and rune length of the strings
		// Hint: Use len and utf8.RuneCountInString
		fmt.Printf("\thas %d bytes and %d runes\n",
			len(s), utf8.RuneCountInString(s))

		// Print the bytes of the strings in hexadecimal
		// Hint: Use % x verb
		fmt.Printf("\tbytes   : % x\n", s)

		// Print the runes of the strings in hexadecimal
		// Hint: Use % x verb
		fmt.Print("\trunes   :")
		for _, r := range s {
			fmt.Printf("% x", r)
		}
		fmt.Println()

		// Print the runes of the strings as rune literals
		// Hint: Use for range
		fmt.Print("\trunes   :")
		for _, r := range s {
			fmt.Printf("%q", r)
		}
		fmt.Println()

		// Print the first rune and its byte size of the strings
		// Hint: Use utf8.DecodeRuneInString
		r, size := utf8.DecodeRuneInString(s)
		fmt.Printf("\tfirst   : %q (%d bytes)\n", r, size)

		// Print the last rune of the strings
		// Hint: Use utf8.DecodeLastRuneInString
		r, size = utf8.DecodeLastRuneInString(s)
		fmt.Printf("\tfirst   : %q (%d bytes)\n", r, size)

		// Slice and print the first two runes of the strings
		_, first := utf8.DecodeRuneInString(s)
		_, second := utf8.DecodeRuneInString(s[first:])
		fmt.Printf("\tfirst 2 : %q\n", s[:first+second])

		// Slice and print the last two runes of the strings
		_, last1 := utf8.DecodeLastRuneInString(s)
		_, last2 := utf8.DecodeLastRuneInString(s[:len(s)-last1])
		fmt.Printf("\tlast 2  : %q\n", s[len(s)-last2-last1:])

		// Convert the string to []rune
		// Print the first and last two runes
		rs := []rune(s)
		fmt.Printf("\tfirst 2 : %q\n", string(rs[:2]))
		fmt.Printf("\tlast 2  : %q\n", string(rs[len(rs)-2:]))
	}
}
