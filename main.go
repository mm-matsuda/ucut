package main

import (
	"fmt"
	"os"
	"log"

	"gopkg.in/urfave/cli.v1"
	"bufio"
	"strings"
	"strconv"
	"errors"
)

var Delimiter = "\000"

func splitFields(fieldSpec string) ([]int, error) {
	var fields []int

	specifications := strings.Split(fieldSpec, ",")
	for _, spec := range specifications {
		if strings.Contains(spec, "-") {
			specRange := strings.Split(spec, "-")
			if len(specRange) != 2 {
				return nil, errors.New("bad fields specification")
			}
			rangeStart, err := strconv.Atoi(specRange[0])
			if err != nil {
				return nil, err
			}
			rangeEnd, err := strconv.Atoi(specRange[1])
			if err != nil {
				return nil, err
			}
			for i := rangeStart-1; i < rangeEnd; i++ {
				fields = append(fields, i)
			}
		} else {
			s, err := strconv.Atoi(spec)
			if err != nil {
				return nil, err
			}
			fields = append(fields, s-1)
		}
	}

	return fields, nil
}

func main() {
	app := cli.NewApp()
	app.Name = "ucut"
	app.Usage = "useful cut at fields"

	app.Flags = []cli.Flag {
		cli.StringFlag{
			Name: "lang",
			Value: "english",
			Usage: "language for the greeting",
		},
		cli.StringFlag{
			Name: "bytes, b",
			Usage: "The `list` specifies byte positions.",
		},
		cli.StringFlag{
			Name: "character, c",
			Usage: "The `list` specifies character positions.",
		},
		cli.StringFlag{
			Name: "delimiter, d",
			Usage: "Use `delim` as the field delimiter character instead of the tab character.",
		},
		cli.StringFlag{
			Name: "fields, f",
			Usage: "The `list` specifies fields, separated in the input by the field delimiter character (see the -d option.)  Output fields are separated by a single occurrence of the field delimiter character.",
	    },
		cli.StringFlag{
			Name: "n",
			Usage: "Do not split multi-byte characters.  Characters will only be output if at least one byte is selected, and, after a prefix of zero or more unselected bytes, the rest of the bytes that form the character are selected.",
		},
		cli.StringFlag{
			Name: "s",
			Usage: "Suppress lines with no field delimiter characters.  Unless specified, lines with no delimiters are passed through unmodified. ",
		},
	}

	app.Action = func(context *cli.Context) error {
		var fields []int
		scanner := bufio.NewScanner(os.Stdin)

		if context.IsSet("delimiter") {
			Delimiter = context.String("delimiter")
		}
		if !context.IsSet("fields") {

		}
		fields, err := splitFields(context.String("fields"))
		if err != nil {
			return err
		}
		for scanner.Scan() {
			words := strings.Split(scanner.Text(), Delimiter)
			wordsLen := len(words)
			for _, f := range fields {
				if f >= wordsLen {
					continue
				}
				fmt.Printf("%s ", words[f])
			}
			fmt.Print("\n")
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
