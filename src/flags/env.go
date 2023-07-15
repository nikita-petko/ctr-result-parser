package flags

func applyEnvironmentVariableFlags() {
	getEnvironmentVariableOrFlag("RESULTS", ResultsFlag)
}
