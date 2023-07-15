package flags

import "flag"

var (
	// HelpFlag prints the usage.
	HelpFlag = flag.Bool("help", false, "Print usage.")

	// ResultsFlag is a flag that passes a list of error codes to parse, if not specified we will try parse other command line arguments as error codes, if that fails we will parse stdin.
	ResultsFlag = flag.String("results", "", "A comma separated list of error codes to parse, if not specified we will try parse other command line arguments as error codes, if that fails we will parse stdin. (Environment variable: RESULTS)")
)

const FlagsUsageString string = `
	[-h|--help] [--results] [...errorCodes]`
