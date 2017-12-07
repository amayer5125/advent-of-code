package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var program_description = regexp.MustCompile(`([a-z]+) \((\d+)\)(?: -> (.*))?`)

type Program struct {
	Name     string
	Weight   int
	Children []*Program
}

func (p *Program) TotalWeight() (w int) {
	w += p.Weight

	if p.HasChildren() {
		for _, v := range p.Children {
			w += v.TotalWeight()
		}
	}

	return
}

func (p *Program) HasChildren() bool {
	return len(p.Children) > 0
}

func (p *Program) OutputWeightTree(prefix string) {
	fmt.Printf("%s%d %s[%d]\n", prefix, p.TotalWeight(), p.Name, p.Weight)

	if p.HasChildren() {
		for _, v := range p.Children {
			v.OutputWeightTree(prefix + "    ")
		}
	}
}

func getRoot(d []string) (p string) {
	var parents []string
	var isChild = make(map[string]bool)

	for _, v := range d {
		if !strings.Contains(v, " -> ") {
			continue
		}

		parts := program_description.FindStringSubmatch(v)

		parents = append(parents, parts[1])

		children := strings.Split(parts[3], ", ")
		for _, i := range children {
			isChild[i] = true
		}
	}

	for _, v := range parents {
		if isChild[v] {
			continue
		}

		p = v

		break
	}

	return
}

func parseTree(d []string, r string) *Program {
	p := Program{}

	for k, v := range d {
		if !strings.HasPrefix(v, r+" ") {
			continue
		}

		parts := program_description.FindStringSubmatch(v)

		p.Name = parts[1]
		p.Weight, _ = strconv.Atoi(parts[2])

		if len(parts[3]) > 0 {
			// remove current item from program description list
			d = append(d[:k], d[k+1:]...)

			// get each child recusivly
			children := strings.Split(parts[3], ", ")
			for _, v := range children {
				p.Children = append(p.Children, parseTree(d, v))
			}
		}

		break
	}

	return &p
}

func main() {
	var d []string
	var root *Program

	// get the input
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		d = append(d, scanner.Text())
	}

	r := getRoot(d)

	root = parseTree(d, r)

	for _, v := range root.Children {
		v.OutputWeightTree("")
	}
}
