package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/nikita-petko/ctr-result-parser/descriptions"
	"github.com/nikita-petko/ctr-result-parser/flags"
	"github.com/nikita-petko/ctr-result-parser/nn"
)

var applicationName string
var buildMode string
var commitSha string

// Pre-setup, runs before main.
func init() {
	flags.SetupFlags(applicationName, buildMode, commitSha)
}

func parseResult(result string) (uint32, bool) {
	// Allow hex like 0x and no 0x.
	if strings.TrimPrefix(result, "0x") != result {
		result = result[2:]
	}

	// Clean the result so we only have 0-9 and a-f.
	result = strings.ToLower(result)
	result = regexp.MustCompile("[^0-9a-f]").ReplaceAllString(result, "")

	// Parse the result.
	parsedResult, err := strconv.ParseUint(result, 16, 32)
	if err != nil {
		// Try to parse as a decimal.
		parsedResult, err = strconv.ParseUint(result, 10, 32)
		if err != nil {
			os.Stderr.WriteString(fmt.Sprintf("Warning: Failed to parse result \"%s\" as hex or decimal. Reason: %s\n", result, err.Error()))

			return 0, false
		}
	}

	return uint32(parsedResult), true
}

func parseResults(results string) []uint32 {
	split := strings.Split(results, ",")
	parsedResults := make([]uint32, len(split))

	for i, result := range split {
		v, ok := parseResult(result)
		if ok {
			parsedResults[i] = v
		} else {
			// Remove the result.
			parsedResults = append(parsedResults[:i], parsedResults[i+1:]...)
		}
	}

	return parsedResults
}

func printResults(results []uint32) {
	for i, result := range results {
		nnResult := nn.NewResult(result)

		fmt.Printf("Result #%d:\n%s\nDescription: %s\n\n", i+1, nnResult.ToString(false), descriptions.GetDescription(nnResult))
	}
}

// Main entrypoint.
func main() {
	if *flags.HelpFlag {
		flag.Usage()

		return
	}

	// 4 ways of reading results:
	// 1. From the argument -results
	//    ctr-result-parser -results 0x1234567890,0x1234567890,0x1234567890
	// 2. If no arguments are specified (no parsed args) but there is still args to parse, we will assume a call like this:
	// 		ctr-result-parser 0x1234567890 0x1234567890 0x1234567890
	// 3. If no arguments are specified (no parsed args) and there is no args to parse, we will assume a call like this:
	// 		echo 0x1234567890,0x1234567890,0x1234567890 | ctr-result-parser
	// 4. If no arguments are specified (no parsed args) and there is no args to parse and no stdin, we will assume a call like this:
	// 		ctr-result-parser -> Will prompt for input

	// Try argument parsing.
	if *flags.ResultsFlag != "" {
		// Parse the results.
		results := parseResults(*flags.ResultsFlag)

		// Print the results.
		printResults(results)

		return
	}

	// Try raw argument parsing.
	if len(flag.Args()) > 0 {
		// Parse the results.
		results := parseResults(strings.Join(flag.Args(), ","))

		// Print the results.
		printResults(results)

		return
	}

	// Try stdin parsing.
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		// Read the results.
		bytes, err := io.ReadAll(os.Stdin)
		if err != nil {
			panic(err)
		}

		// Parse the results.
		results := parseResults(string(bytes))

		// Print the results.
		printResults(results)

		return
	}

	// Prompt for input.
	fmt.Println("Please enter the results (comma separated):")

	// Read the results.
	var results string
	_, err := fmt.Scanln(&results)
	if err != nil {
		panic(err)
	}

	// Parse the results.
	parsedResults := parseResults(results)

	// Print the results.
	printResults(parsedResults)
}
