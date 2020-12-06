// Ensure that each blank-line delimited block is missing the "cid" field at most.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type passport struct {
	byr, iyr, eyr, hgt, hcl, ecl, pid, cid string
}

func (pport passport) Display() {
	fmt.Println("byr: ", pport.byr, "\t", pport.ValidateByr(), "(1920, 2002)")
	fmt.Println("iyr: ", pport.iyr, "\t", pport.ValidateIyr(), "(2010, 2020)")
	fmt.Println("eyr: ", pport.eyr, "\t", pport.ValidateEyr(), "(2020, 2030)")
	fmt.Println("hgt: ", pport.hgt, "\t", pport.ValidateHgt(), "(150, 193)cm, (59,76)in")
	fmt.Println("hcl: ", pport.hcl, "\t", pport.ValidateHcl(), "6 digits: 0-9, a-f")
	fmt.Println("ecl: ", pport.ecl, "\t", pport.ValidateEcl(), "amb, blu, brn, gry, grn, hzl, oth")
	fmt.Print("pid:  ", pport.pid, "  ", pport.ValidatePID(), " 9 digits", "\n")
	fmt.Println("cid: ", pport.cid)
}

func (pport passport) Validate() bool {
	if pport.ValidateByr() && pport.ValidateIyr() && pport.ValidateEyr() &&
		pport.ValidateHgt() && pport.ValidateHcl() &&
		pport.ValidateEcl() && pport.ValidatePID() {
		return true
	} else {
		return false
	}
}

func (pport passport) ValidateByr() bool {
	if pport.byr == "" {
		return false
	} else if numStrInRng(pport.byr, [2]int{1920, 2002}) {
		return true
	} else {
		return false
	}

}

func (pport passport) ValidateIyr() bool {
	if pport.iyr == "" {
		return false
	} else if numStrInRng(pport.iyr, [2]int{2010, 2020}) {
		return true
	} else {
		return false
	}
}

func (pport passport) ValidateEyr() bool {
	if pport.eyr == "" {
		return false
	} else if numStrInRng(pport.eyr, [2]int{2020, 2030}) {
		return true
	} else {
		return false
	}
}

func (pport passport) ValidateHgt() bool {
	if pport.hgt == "" {
		return false
	} else {
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
}

func (pport passport) ValidateHcl() bool {
	if pport.hcl == "" {
		return false
	} else {
		hclAry := strings.Split(pport.hcl, "")
		if len(hclAry) == 7 && hclAry[0] == "#" {
			for _, char := range hclAry[1:] {
				if !strings.ContainsAny(char, "0123456789abcdef") {
					return false
				}
			}
			return true
		} else {
			return false
		}
	}
}

func (pport passport) ValidateEcl() bool {
	if pport.ecl == "" {
		return false
	} else {
		clrs := [7]string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
		for _, clr := range clrs {
			if pport.ecl == clr {
				return true
			}
		}
		return false
	}
}

func (pport passport) ValidatePID() bool {
	if pport.pid == "" {
		return false
	} else {
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
}

func numStrInRng(numStr string, rng [2]int) bool {
	numInt, _ := strconv.Atoi(numStr)
	if numInt >= rng[0] && numInt <= rng[1] {
		return true
	} else {
		return false
	}
}

func (pport passport) SetFields(line string) passport {
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
		// Check the passport upon encountering a blank line
		if line == "" {
			if pport.Validate() {
				valid += 1
				fmt.Println("----------")
				fmt.Println("Valid passport", valid)
				fmt.Println("----------")
				pport.Display()
			}
			pport = passport{}
		} else {
			pport = pport.SetFields(line)
		}
	}
	return valid
}

func main() {
	path := ("4.txt")
	fmt.Println(countValidPassports(path))
}
