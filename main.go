package main

import (
	"fmt"
	"strings"
	"io/ioutil"
	"os"
	"strconv"
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
func count1to8(word string) []string {
        list := []string {word}
        for i := 0; i < 9; i++ {
                list = append(list,list[len(list)-1]+strconv.Itoa(i))
        }
        return list
}
func year90(word string) []string {
        list := []string {word}
        for i := 99; i > 89; i-- {
                list = append(list,word+strconv.Itoa(i))
        }
        return list
}
func year2000(word string) []string {
        list := []string {word}
        for i := 2020; i > 1999; i-- {
                list = append(list,word+strconv.Itoa(i))
        }
        return list
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
func mess(word string, all_list []string) []string {
	list := []string {}
	for _, i := range all_list {
		list = append(list, word + i)
	}
	return list
}
func combine(all_list []string) []string {
	final_words := []string {}

	for _, i := range all_list {
		final_words = append(final_words, mess(i, all_list)...)
	}

	return final_words
}

func main() {
	final_content := []string {}
	// Read base file
	blob, err := ioutil.ReadFile(os.Args[1])
	check(err)
	content := strings.Fields(string(blob))

	// Running functions for make wordlist
	for _, word := range content {
		final_content = content
		final_content = append(final_content, combine(content)...)
		final_content = append(final_content, leet(word))
		final_content = append(final_content, count1to8(word)...)
		final_content = append(final_content, year90(word)...)
		final_content = append(final_content, year2000(word)...)
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
