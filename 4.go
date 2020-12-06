// Ensure that each blank-line delimited block is missing the "cid" field at most.
package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

type passport struct {
	byr, iyr, eyr, hgt, hcl, ecl, pid, cid string
}

func (pport passport) Validate() bool {
	if pport.byr == "" || pport.iyr == "" || pport.eyr == "" ||pport.hgt == "" ||
		pport.hcl == "" || pport.ecl == "" || pport.pid == "" {
		return false
	} else {
		return true
	}
}

func (pport passport) setFields(line string) passport {
	pairs := strings.Split(line, " ")
	for _, pair := range pairs {
		pair := strings.Split(pair, ":")
		switch pair[0] {
		case "byr":
			pport.byr = pair[1]
		case "iyr":
			pport.iyr = pair[1]
		case "eyr":
			pport.eyr = pair[1]
		case "hgt":
			pport.hgt = pair[1]
		case "hcl":
			pport.hcl = pair[1]
		case "ecl":
			pport.ecl = pair[1]
		case "pid":
			pport.pid = pair[1]
		case "cid":
			pport.cid = pair[1]
		}
	}
	return pport
}

func countValidPassports(path string) int {
	file, _ := os.Open(path)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	valid := 0
	var pport passport
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			if pport.Validate() {
				valid += 1
				fmt.Println(pport, "is a valid passport")
			} else {
				fmt.Println(pport, "is NOT a valid passport")
			}
			pport = passport{}
		} else {
			pport = pport.setFields(line)
		}
	}
	return valid
}

func main() {
	path := ("4.txt")
	fmt.Println(countValidPassports(path))
}

