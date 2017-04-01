package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"
)

const (
	//black
	red = iota
	green
	brown
	blue
	magenta
	cyan
	gray
	dark_gray
	light_red
	light_green
	yellow
	light_blue
	light_magenta
	light_cyan
	white
	end
)

var allColors map[int]string

func init() {
	allColors = make(map[int]string)
	//allColors[black] = "\033[22;30m"
	allColors[red] = "\033[22;31m"
	allColors[green] = "\033[22;32m"
	allColors[brown] = "\033[22;33m"
	allColors[blue] = "\033[22;34m"
	allColors[magenta] = "\033[22;35m"
	allColors[cyan] = "\033[22;36m"
	allColors[gray] = "\033[22;37m"
	allColors[dark_gray] = "\033[01;30m"
	allColors[light_red] = "\033[01;31m"
	allColors[light_green] = "\033[01;32m"
	allColors[yellow] = "\033[01;33m"
	allColors[light_blue] = "\033[01;34m"
	allColors[light_magenta] = "\033[01;35m"
	allColors[light_cyan] = "\033[01;36m"
	allColors[white] = "\033[01;37m"
	allColors[end] = "\033[0m"
}

func getJsonKeys(s string) []string {
	var feilds map[string]interface{}
	var keys []string
	err := json.Unmarshal([]byte(s), &feilds)
	if err != nil {
		return keys
	}

	for k := range feilds {
		keys = append(keys, "\""+k+"\"")
	}
	return keys
}

func printWithColor(s string, subs ...string) {
	tmps := s
	keys := getJsonKeys(s)
	for _, sub := range subs {
		keys = append(keys, sub)
	}

	sort.Strings(keys)
	for i, k := range keys {
		var color int
		color = i % 16
		tmps = strings.Replace(tmps, k, allColors[color]+k+allColors[end], -1)
	}
	fmt.Printf(tmps)
}
func main() {
	flag.Parse()
	args := flag.Args()
	//fmt.Println(args)
	reader := bufio.NewReader(os.Stdin)
	for {
		log, _ := reader.ReadString('\n')
		//fmt.Println(log)
		printWithColor(log, args...)
	}
}
