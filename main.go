package main

import (
	"fmt"
	"strings"
	"io/ioutil"
	"os"
)
// Check if err
func check(e error) {
    if e != nil {
        panic(e)
    }
}
// Removing duplicated items
func unique(intSlice []string) []string {
    keys := make(map[string]bool)
    list := []string {}
    for _, entry := range intSlice {
        if _, value := keys[entry]; !value {
            keys[entry] = true
            list = append(list, entry)
        }
    }
    return list
}

// Pure functions
func leet(word string) string {
	return strings.NewReplacer("A","4", "E", "3", "I", "1", "O", "0", "S", "5", "T", "7", "B", "8").Replace(word)
}
func count1(word string) string {
	return word+"1"
}
func count2(word string) string {
	return word+"12"
}
func count3(word string) string {
	return word+"123"
}
func count4(word string) string {
	return word+"1234"
}
func count8(word string) string {
	return word+"12345678"
}
func count81(word string) string {
	return word+"87654321"
}
func year16(word string) string {
	return word+"2016"
}
func year17(word string) string {
	return word+"2017"
}
func year18(word string) string {
	return word+"2018"
}
func year19(word string) string {
	return word+"2019"
}
func year20(word string) string {
	return word+"2020"
}
func swapCase(r rune) rune {
    switch {
    case 'a' <= r && r <= 'z':
        return r - 'a' + 'A'
    case 'A' <= r && r <= 'Z':
        return r - 'A' + 'a'
    default:
        return r
    }
}
func inverter(word string) string {
	runes := []rune(string(word))
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	final_content := []string {}
	// Read base file
	blob, err := ioutil.ReadFile(os.Args[1])
	check(err)
	content := strings.Fields(string(blob))

	// Running functions for make wordlist
	for _, word := range content {
		final_content = append(final_content, leet(word))
		final_content = append(final_content, count1(word))
		final_content = append(final_content, count2(word))
		final_content = append(final_content, count3(word))
		final_content = append(final_content, count4(word))
		final_content = append(final_content, count8(word))
		final_content = append(final_content, count81(word))
		final_content = append(final_content, count81(word))
		final_content = append(final_content, year16(word))
		final_content = append(final_content, year17(word))
		final_content = append(final_content, year18(word))
		final_content = append(final_content, year19(word))
		final_content = append(final_content, year20(word))
		final_content = append(final_content, strings.ToUpper(word))
		final_content = append(final_content, strings.ToLower(word))
		final_content = append(final_content, strings.Map(swapCase, word))
		final_content = append(final_content, inverter(word))
	}

	// Save data
	file, err_create := os.Create(os.Args[2])
	check(err_create)
	defer file.Close()
	data := []byte(strings.Join(unique(final_content), "\n"))
	number_wrote, err_write := file.Write(data)
	check(err_write)
	fmt.Printf("File saved in %s: %db\n", os.Args[2], number_wrote)
}
