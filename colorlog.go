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

func getJsonKeys(s string) ([]string, string) {
	var feilds map[string]interface{}
	var keys []string
	err := json.Unmarshal([]byte(s), &feilds)
	if err != nil {
		return keys, ""
	}

	level := ""
	for k, v := range feilds {
		keys = append(keys, "\""+k+"\"")
		if k == "level" {
			level = v.(string)
		}
	}
	return keys, level
}

func printWithColor(s string, subs []string) {
	tmps := s
	keys, level := getJsonKeys(s)
	sort.Strings(keys)
	index := 0
	for _, k := range keys {
		index++

		if index%8 == 1 {
			index++
		}

		if index%8 == 0 {
			index += 2
		}

		colorType := ColorType(index)
		switch level {
		case "error":
			tmps = strings.Replace(tmps, k, Colors[ColorOff]+Colors[colorType]+k+Colors[ColorOff]+Colors[Red], -1)
		case "warn":
			tmps = strings.Replace(tmps, k, Colors[ColorOff]+Colors[colorType]+k+Colors[ColorOff]+Colors[Yellow], -1)
		case "debug":
			tmps = strings.Replace(tmps, k, Colors[ColorOff]+Colors[colorType]+k+Colors[ColorOff]+Colors[Green], -1)
		default:
			tmps = strings.Replace(tmps, k, Colors[colorType]+k+Colors[ColorOff], -1)
		}
	}

	switch level {
	case "error":
		tmps = Colors[Red] + tmps + Colors[ColorOff]
	case "warn":
		tmps = Colors[Yellow] + tmps + Colors[ColorOff]
	case "debug":
		tmps = Colors[Green] + tmps + Colors[ColorOff]
	}

	sort.Strings(subs)
	for i, sub := range subs {
		colorType := ColorType(i + 2 + 8*3)
		tmps = strings.Replace(tmps, sub, Colors[colorType]+Colors[UnderlineBlack]+sub+Colors[ColorOff]+Colors[ColorOff], -1)
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
		printWithColor(log, args)
	}
}
