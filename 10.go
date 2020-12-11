package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type adapter struct {
	jo    int
	dnext int
	priors   int
}

func newAdapter(jo int) adapter {
	a := adapter{jo, 0, 0}
	return a
}

func listSortedAdapters(path string) []adapter {
	file, _ := os.Open(path)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var joltOuts []int
	for scanner.Scan() {
		jo, _ := strconv.Atoi(scanner.Text())
		joltOuts = append(joltOuts, jo)
	}
	sort.Ints(joltOuts)
	var adapters []adapter
	outlet := newAdapter(0)
	adapters = append(adapters, outlet)
	for _, v := range joltOuts {
		adapters = append(adapters, newAdapter(v))
	}
	device := newAdapter(adapters[len(adapters)-1].jo + 3)
	adapters = append(adapters, device)
	return adapters
}

func min(a []int) int {
	m := a[0]
	for _, v := range a {
		if v < m {
			m = v
		}
	}
	return m
}

func adapterDist(adapters []adapter) []adapter {
	for i := 1; i < len(adapters); i++ {
		switch adapters[i].jo - adapters[i-1].jo {
		case 1:
			adapters[i-1].dnext = 1
		case 2:
			adapters[i-1].dnext = 2
		case 3:
			adapters[i-1].dnext = 3
		}
	}
	return adapters
}

func adapterDistSums(adapters []adapter) map[string]int {
	distSums := make(map[string]int)
	for _, a := range adapters {
		switch a.dnext {
		case 1:
			distSums["one"] += 1
		case 2:
			distSums["two"] += 1
		case 3:
			distSums["three"] += 1
		}
	}
	return distSums
}

func numAdapterChains(adapters []adapter) []int {
	numChains := []int{1}
	for i := 1; i < (len(adapters) - 1); i++ {
		if (adapters[i].dnext == 1) && (adapters[i+1].dnext == 1) {
			numChains = append(numChains, 2)
		}
	}
	return numChains
}

func sumSlice(s []int) int {
	sum := 0
	for _, n := range s {
		sum += n
	}
	return sum
}

func countPaths(adapters []adapter) int {
	for i := 1; i < len(adapters); i++ {
		sliceLow := 0
		if i - 3 < 0 {
			sliceLow = 0
		} else {
			sliceLow = i - 3
		}
		for _, posPrior := range adapters[sliceLow:i] {
			if adapters[i].jo - posPrior.jo <= 3 {
				if posPrior.priors > 1 {
					adapters[i].priors += posPrior.priors
				} else {
					adapters[i].priors += 1
				}
			}
		}
	}
	return adapters[len(adapters)-1].priors
}

func main() {
	// path := "10_small1.txt"
	// path := "10_small2.txt"
	path := "10.txt"
	adapters := adapterDist(listSortedAdapters(path))
	fmt.Println("Part 1:", adapterDistSums(adapters)["one"] * adapterDistSums(adapters)["three"])
	fmt.Println("Part 2:", countPaths(adapters))
}
