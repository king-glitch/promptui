package main

import (
	"fmt"
	"strings"

	"github.com/king-glitch/promptui"
)

func main() {
	var options = []string{
		"Monday",
		"Tuesday",
		"Wednesday",
		"Thursday",
		"Friday",
		"Saturday",
		"Sunday",
	}

	prompt := promptui.MultiSelect{
		Label: "Select Day",
		Items: options,
		Searcher: func(
			input string,
			index int,
		) bool {
			option := options[index]
			return (len(input) == 0) || strings.TrimSpace(input) == strings.TrimSpace(option)
		},
	}

	indexes, err := prompt.Run()

	if err != nil {
		fmt.Printf(
			"Prompt failed %v\n",
			err,
		)
		return
	}

	for _, index := range indexes {
		fmt.Printf(
			"You choose %q\n",
			options[index],
		)
	}
}
