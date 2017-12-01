package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
)

func ReadFile(filename string) string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return ""
	}

	return string(data)
}

// compile patterns
var value_patt = regexp.MustCompile(`value (\d+) goes to bot (\d+)`)
var hand_off_patt = regexp.MustCompile(`bot (\d+) gives low to (bot|output) (\d+) and high to (bot|output) (\d+)`)

type Bot struct {
	low, high  int
	directions []string
}
type Bots map[string]*Bot

func (from *Bot) GiveChip(to *Bot, hl string) {
	// transfer value
	var tv int
	if hl == "low" {
		tv = from.low
		from.low = 0
	} else {
		tv = from.high
		from.high = 0
	}

	// check which slot the chip needs to go in, moving values if needed
	if to.low == 0 {
		to.low = tv
	} else if tv > to.low {
		to.high = tv
	} else {
		to.high = to.low
		to.low = tv
	}
}

type Output struct {
	values []int
}
type Outputs map[string]*Output

func (from *Bot) GiveChipToOutput(to *Output, hl string) {
	if hl == "low" {
		to.values = append(to.values, from.low)
		from.low = 0
	} else {
		to.values = append(to.values, from.high)
		from.high = 0
	}
}

func main() {
	directions := ReadFile("input.txt")

	var bots = make(Bots)
	var outputs = make(Outputs)

	// loop through each starting value
	start_values := value_patt.FindAllStringSubmatch(directions, -1)
	for _, v := range start_values {
		// create bot if it does not exist
		if bots[v[2]] == nil {
			bots[v[2]] = &Bot{}
		}
		// microchip value
		mc_val, _ := strconv.Atoi(v[1])
		// inputs are bots too #inputlivesmatter
		input := Bot{low: mc_val}
		// give bot microchip
		input.GiveChip(bots[v[2]], "low")
	}

	// loop through each hand off
	hand_offs := hand_off_patt.FindAllStringSubmatch(directions, -1)
	for _, h := range hand_offs {
		if bots[h[1]] == nil {
			bots[h[1]] = &Bot{}
		}
		// give directions to bot
		bots[h[1]].directions = h[2:]
	}

	for {
		// loop through each bot and see what it is to do
		for _, bot := range bots {
			if bot.low == 0 || bot.high == 0 {
				continue
			}

			d := bot.directions
			if d[0] == "bot" {
				bot.GiveChip(bots[d[1]], "low")
			} else {
				if outputs[d[1]] == nil {
					outputs[d[1]] = &Output{}
				}
				bot.GiveChipToOutput(outputs[d[1]], "low")
			}
			if d[2] == "bot" {
				bot.GiveChip(bots[d[3]], "high")
			} else {
				if outputs[d[3]] == nil {
					outputs[d[3]] = &Output{}
				}
				bot.GiveChipToOutput(outputs[d[3]], "high")
			}

			bot.directions = []string{}

			// check if one of the bots in the directions was responsible for comparing 17 and 61
			if bots[d[1]].low == 17 && bots[d[1]].high == 61 {
				fmt.Printf("Bot %v compared 61 and 17\n", d[1])
				return
			}
			if bots[d[3]].low == 17 && bots[d[3]].high == 61 {
				fmt.Printf("Bot %v compared 61 and 17\n", d[3])
				return
			}
		}
	}
}
