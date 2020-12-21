package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type expr struct {
	symb []string
	prec []int
}

func max(a []int) int {
	m := 0
	for i := 0; i < len(a); i++ {
		if a[i] > m {
			m = a[i]
		}
	}
	return m
}

func line2Expr1(line string) expr {
	chars := strings.Split(strings.ReplaceAll(line, " ", ""), "")
	ex := expr{}
	depth := 0
	for _, c := range chars {
		if c == "(" {
			depth += 1
		} else if c == ")" {
			depth -= 1
		} else {
			ex.symb = append(ex.symb, c)
			ex.prec = append(ex.prec, depth)
		}
	}
	return ex
}

// func parenIndex(a []string, p string) int {
// 	if p == ")" {
// 		for i := 0; i < len(a); i++ {
// 			if a[i] == ")" {
// 				return i
// 			}
// 		}
// 	} else if p == "(" {
// 		for i := len(a) - 1; i >= 0; i-- {
// 			if a[i] == "(" {
// 				return i
// 			}
// 		}
// 	}
// 	return -1
// }

func parenIndex(a []string, p string) int {
	nestLvl := 0
	if p == ")" {
		for i := 0; i < len(a); i++ {
			if a[i] == "(" {
				nestLvl += 1
			}
			if a[i] == ")" {
				nestLvl -= 1
			}
			if (a[i] == ")" && nestLvl == 0) || i == len(a)-1 {
				return i
			}
		}
	} else if p == "(" {
		for i := len(a) - 1; i >= 0; i-- {
			if a[i] == ")" {
				nestLvl += 1
			}
			if a[i] == "(" {
				nestLvl -= 1
			}
			if (a[i] == "("  && nestLvl == 0) || i == 0 {
				return i
			}
		}
	}
	return -1
}

func bindAddition(ex expr, a []string) expr {
	for i := 1; i < len(a)-1; i++ {
		if a[i] == "+" {
			prev := a[:i]
			next := a[i+1:]
			this := []string{"+"}
			if next[0] == "(" {
				end := parenIndex(next, ")")
				this = append(this, next[:end+1]...)
				this = append(this, ")")
				this = append(this, next[end+1:]...)
			} else {
				this = append(this, next[0])
				this = append(this, ")")
				this = append(this, next[1:]...)
			}
			if prev[len(prev)-1] == ")" {
				beg := parenIndex(prev, "(")
				this = append(prev[beg:], this...)
				this = append([]string{"("}, this...)
				this = append(prev[:beg], this...)
			} else {
				this = append([]string{prev[len(prev)-1]}, this...)
				this = append([]string{"("}, this...)
				this = append(prev[:len(prev)-1], this...)
			}
			depths := []int{}
			depth := 0
			for i := 0; i < len(this); i++ {
				if this[i] == "(" {
					depth = depth + 1
				} else if this[i] == ")" {
					depth = depth - 1
				} else {
					depths = append(depths, depth)
				}
			}
			for i := 0; i < len(depths); i++ {
				ex.prec[i] += depths[i]
			}
		}
	}
	return ex
}

func line2Expr2(line string) expr {
	line = strings.ReplaceAll(line, " ", "")
	chars := strings.Split(line, "")
	ex := expr{}
	depth := 0
	for _, c := range chars {
		if c == "(" {
			depth += 1
		} else if c == ")" {
			depth -= 1
		} else {
			ex.symb = append(ex.symb, c)
			ex.prec = append(ex.prec, depth)
		}
	}
	ex = bindAddition(ex, chars)
	return ex
}

func (ex expr) OpTri(i int) (expr, bool) {
	if i+3 > len(ex.prec) {
		return ex, false
	}
	if !(ex.prec[i+1] == ex.prec[i] && ex.prec[i+2] == ex.prec[i]) {
		return ex, false
	}
	left := expr{ex.symb[:i+1], ex.prec[:i+1]}
	right := expr{ex.symb[i+3:], ex.prec[i+3:]}
	n1, _ := strconv.Atoi(ex.symb[i])
	n2, _ := strconv.Atoi(ex.symb[i+2])
	switch ex.symb[i+1] {
	case "+":
		left.symb[i] = strconv.FormatInt(int64(n1+n2), 10)
	case "*":
		left.symb[i] = strconv.FormatInt(int64(n1*n2), 10)
	}
	ex = left
	for j := 0; j < len(right.symb); j++ {
		ex.symb = append(ex.symb, right.symb[j])
		ex.prec = append(ex.prec, right.prec[j])
	}
	return ex, true
}

func (ex expr) SyncPrec(prec int) expr {
	if len(ex.prec) > 1 {
		// Is the last element of a given precedence isolated?
		if ex.prec[len(ex.prec)-1] == prec {
			if !(ex.prec[len(ex.prec)-2] == prec) {
				ex.prec[len(ex.prec)-1] -= 1
			}
		}
		// Is the first?
		if ex.prec[0] == prec {
			if !(ex.prec[1] == prec) {
				ex.prec[0] -= 1
			}
		}
	}
	// What about those between?
	for i := 1; i < len(ex.prec)-1 && len(ex.prec) > 2; i++ {
		isolated := !(ex.prec[i-1] == prec || ex.prec[i+1] == prec)
		if ex.prec[i] == prec && isolated {
			ex.prec[i] -= 1
		}
	}
	return ex
}

func evalExpr(ex expr) expr {
	prec := max(ex.prec)
	for prec >= 0 {
		i := 0
		for i < len(ex.symb) {
			if ex.prec[i] == prec {
				if e, ok := ex.OpTri(i); ok {
					ex = e
					i--
				}
			}
			i = i + 1
		}
		ex = ex.SyncPrec(prec)
		prec--
	}
	return ex
}

func sumExprs(path string, part int) int {
	file, _ := os.Open(path)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	sum := 0
	if part == 1 {
		for scanner.Scan() {
			line := scanner.Text()
			n, _ := strconv.Atoi(evalExpr(line2Expr1(line)).symb[0])
			sum += n
		}
	} else if part == 2 {
		for scanner.Scan() {
			line := scanner.Text()
			n, _ := strconv.Atoi(evalExpr(line2Expr2(line)).symb[0])
			sum += n
		}

	}
	return sum
}

func testSampleData() {
	path := "18_small1.txt"
	file, _ := os.Open(path)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	num := 0
	ans1 := [...]int{71, 51, 26, 437, 12240, 13632}
	ans2 := [...]int{231, 51, 46, 1445, 669060, 23340}
	for scanner.Scan() {
		line := scanner.Text()
		a1, _ := strconv.Atoi(evalExpr(line2Expr1(line)).symb[0])
		a2, _ := strconv.Atoi(evalExpr(line2Expr2(line)).symb[0])
		m1, m2 := a1 == ans1[num], a2 == ans2[num]
		fmt.Print(num+1, ". ", a1, " (", m1, ") ", a2, " (", m2, ")\n")
		if m1 == false {
			fmt.Println("ERROR:", a1, "should be", ans1[num])
		}
		if m2 == false {
			fmt.Println("ERROR:", a2, "should be", ans2[num])
		}
		num += 1
	}
}

func main() {
	// testSampleData()
	// line := "((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2)"
	// ex1 := line2Expr1(line)
	// ex2 := line2Expr2(line)
	// evalExpr(ex1)
	// fmt.Println(evalExpr(ex1))
	// fmt.Println(evalExpr(ex1), evalExpr(ex2))
	// path := "18_small1.txt"
	path := "18.txt"
	fmt.Println("Part 1:", sumExprs(path, 1))
	fmt.Println("Part 2:", sumExprs(path, 2))
}
