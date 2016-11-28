package envflag

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"

	pflag "github.com/spf13/pflag"
)

// MergeFlags will merge flag.FlagSet to the pflag.FlagSet
func MergeFlags() {
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
}

// Parse parses the FlagSet
func Parse() {
	/**
	* Adding flags to goflag as well, glog requires flag.Parse before
	* logging starts. If we don't add any flag in flag, then flag.Parse
	* fails saying flag is not defined.
	*
	* So, for any flag we define we add to pflags and goflags both.
	* Usage only prints pflags
	 */

	flag.Parse()
	pflag.Parse()
	MergeFlags()
}

// Lookup returns a flag value, or nil if it does not exist
func Lookup(name string) *pflag.Flag {
	return pflag.Lookup(name)
}

// Parsed returns true if FlagSet already parsed, false otherwise
func Parsed() bool {
	return pflag.Parsed()
}

// PrintDefaults prints default value for all elements
func PrintDefaults() {
	pflag.PrintDefaults()
}

// PrintValues prints current values for all flags
func PrintValues() {
	fmt.Println("Using values: ")
	pflag.VisitAll(func(f *pflag.Flag) {
		fmt.Printf("        %s = %s\n", f.Name, f.Value)
	})

	fmt.Println()
}

// Set sets the name/value pair in the flag set
func Set(name, value string) error {
	pflag.Set(name, value)
	return nil
}

// Usage prints the flag usage
func Usage() {
	pflag.Usage()
}

// VisitAll invokes pflag.VisitAll
func VisitAll(fn func(*pflag.Flag)) {
	pflag.VisitAll(fn)
}

func verifyNames(flagName, envName string) {
	if flagName == "" && envName == "" {
		panic("No flag or environment name given for " + envName)
	}
}

func lookupEnv(name string) string {
	if name == "" {
		return ""
	}

	return os.Getenv(name)
}

// Bool creates a new entry in the flag set of boolean value.
// The environment value used as a default, if it exists
func Bool(flagName, envName string, value bool, usage string) *bool {
	verifyNames(flagName, envName)
	envValStr := lookupEnv(envName)
	if envValStr != "" {
		value, _ = strconv.ParseBool(envValStr)
	}

	flag.Bool(flagName, value, usage)
	return pflag.Bool(flagName, value, usage)
}

// Float64 creates a new entry in the flag set with float64 value.
// The environment value used as a default, if it exists
func Float64(flagName, envName string, value float64, usage string) *float64 {
	verifyNames(flagName, envName)
	envValStr := lookupEnv(envName)
	if envValStr != "" {
		value, _ = strconv.ParseFloat(envValStr, 64)
	}

	flag.Float64(flagName, value, usage)
	return pflag.Float64(flagName, value, usage)
}

// Int creates a new entry in the flag set with Int value.
// The environment value used as a default, if it exists
func Int(flagName, envName string, value int, usage string) *int {
	verifyNames(flagName, envName)
	envValStr := lookupEnv(envName)
	if envValStr != "" {
		value, _ = strconv.Atoi(envValStr)
	}

	flag.Int(flagName, value, usage)
	return pflag.Int(flagName, value, usage)
}

// Int64 creates a new entry in the flag set with Int64 value.
// The environment value used as a default, if it exists
func Int64(flagName, envName string, value int64, usage string) *int64 {
	verifyNames(flagName, envName)
	envValStr := lookupEnv(envName)
	if envValStr != "" {
		value, _ = strconv.ParseInt(envValStr, 10, 64)
	}

	flag.Int64(flagName, value, usage)
	return pflag.Int64(flagName, value, usage)
}

// String creates a new entry in the flag set with String value.
// The environment value used as a default, if it exists
func String(flagName, envName, value, usage string) *string {
	verifyNames(flagName, envName)
	envVal := lookupEnv(envName)
	if envVal != "" {
		value = envVal
	}

	flag.String(flagName, value, usage)
	return pflag.String(flagName, value, usage)
}

// UInt creates a new entry in the flag set with UInt value.
// The environment value used as a default, if it exists
func UInt(flagName, envName string, value uint, usage string) *uint {
	verifyNames(flagName, envName)
	envValStr := lookupEnv(envName)
	if envValStr != "" {
		ui64, _ := strconv.ParseUint(envValStr, 10, 64)
		value = uint(ui64)
	}

	flag.Uint(flagName, value, usage)
	return pflag.Uint(flagName, value, usage)
}

// UInt64 creates a new entry in the flag set with UInt64 value.
// The environment value used as a default, if it exists
func UInt64(flagName, envName string, value uint64, usage string) *uint64 {
	verifyNames(flagName, envName)
	envValStr := lookupEnv(envName)
	if envValStr != "" {
		value, _ = strconv.ParseUint(envValStr, 10, 64)
	}

	flag.Uint64(flagName, value, usage)
	return pflag.Uint64(flagName, value, usage)
}

// Duration creates a new entry in the flag set with Duration value.
// The environment value used as a default, if it exists
func Duration(flagName, envName string, value time.Duration, usage string) *time.Duration {
	verifyNames(flagName, envName)
	envValStr := lookupEnv(envName)
	if envValStr != "" {
		value, _ = time.ParseDuration(envValStr)
	}

	flag.Duration(flagName, value, usage)
	return pflag.Duration(flagName, value, usage)
}
