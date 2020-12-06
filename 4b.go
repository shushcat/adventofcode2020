// Ensure that each blank-line delimited block is missing the "cid" field at most.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	// "unicode"
)

type passport struct {
	byr, iyr, eyr, hgt, hcl, ecl, pid, cid string
}

func (pport passport) Validate() bool {
	if pport.byr == "" || pport.iyr == "" || pport.eyr == "" || pport.hgt == "" ||
		pport.hcl == "" || pport.ecl == "" || pport.pid == "" {
		return false
	} else if !numStrInRng(pport.byr, [2]int{1920, 2002}) ||
		!numStrInRng(pport.iyr, [2]int{2010, 2020}) ||
		!numStrInRng(pport.eyr, [2]int{2020, 2030}) {
		return false
	} else if !pport.ValidateHgt() || !pport.ValidateHcl() ||
		!pport.ValidateEcl() || !pport.ValidatePID() {
		return false
	} else {
		return true
	}
}

func (pport passport) ValidateHgt() bool {
	hgtAry := strings.Split(pport.hgt, "")
	unit := strings.Join(hgtAry[len(hgtAry)-2:], "")
	hgtStr := strings.Join(hgtAry[:len(hgtAry)-2], "")
	if unit == "cm" {
		return numStrInRng(hgtStr, [2]int{150, 193})
	} else if unit == "in" {
		return numStrInRng(hgtStr, [2]int{59, 76})
	} else {
		return true
	}
}

func (pport passport) ValidateHcl() bool {
	hclAry := strings.Split(pport.hcl, "")
	if len(hclAry) == 7 && hclAry[0] == "#" {
		for _, char := range hclAry[1:] {
			if !strings.ContainsAny(char, "0123456789abcdef") {
				return false
			}
		}
	} else {
		return false
	}
	return true
}

func (pport passport) ValidateEcl() bool {
	clrs := [7]string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
	for _, clr := range clrs {
		if pport.ecl == clr {
			return true
		}
	}
	return false
}

func (pport passport) ValidatePID() bool {
	pidAry := strings.Split(pport.pid, "")
	if len(pidAry) == 9 {
		for _, char := range pidAry {
			if !strings.ContainsAny(char, "0123456789") {
				return false
			}
		}
	} else {
		return false
	}
	return true
}

func numStrInRng(numStr string, rng [2]int) bool {
	numInt, _ := strconv.Atoi(numStr)
	if numInt >= rng[0] && numInt <= rng[1] {
		return true
	} else {
		return false
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

func main() {
	path := ("4.txt")
	fmt.Println(countValidPassports(path))
	// var pport passport
	// pport.hcl = "#74a3e3"
	// pport.ecl = "blu"
	// pport.pid = "000009000"
	// fmt.Println(pport.ValidatePID())
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
				fmt.Println("Valid:  \t", pport)
			} else {
				fmt.Println("Invalid:\t", pport)
			}
			pport = passport{}
		} else {
			pport = pport.setFields(line)
		}
	}
	return valid
}
