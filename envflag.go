package envflag

import (
	goflag "flag"
	"os"
	"strconv"

	flag "github.com/spf13/pflag"
)

const (
	// BitSize - conversions are 64-bit
	BitSize = 64

	// BaseNum - conversions are base-10
	BaseNum = 10
)

func MergeFlags() {
	flag.CommandLine.AddGoFlagSet(goflag.CommandLine)	
}

// Parse parses the FlagSet
func Parse() {
	/*     Adding flags to goflag as well, glog requires flag.Parse before
	 *      logging starts. If we don't add any flag in goflag, then flag.Parse
	 *      fails saying flag is not defined.
	 *      So, any flag we define we add to pflags and goflags both. Usage only
	 *      prints pflags
	 */

	flag.Parse()
	goflag.Parse()
}

func Lookup(name string) *flag.Flag{
	return flag.Lookup(name)
}
// Parsed returns true if FlagSet already parsed, false otherwise
func Parsed() bool {

	return flag.Parsed()
}

// PrintDefaults prints default value for all elements
func PrintDefaults() {

	flag.PrintDefaults()
}

// Set sets the name/value pair in the flag set
func Set(name, value string) error {

	flag.Set(name, value)
	return nil
}

func Usage() {

	flag.Usage()
}

func verifyNames(flagName, envName string) {
	if flagName == "" && envName == "" {
		panic("No flag name or env name given")
	}
}

func lookupFlag(name string) bool {

	if len(name) > 0 {
		if flag.Lookup(name) != nil {
			return true
		}
	}

	return false
}

func lookupEnv(name string) bool {

	if len(name) > 0 {
		envVal := os.Getenv(name)
		if len(envVal) > 0 {
			return true
		}
	}

	return false
}

// Bool creates a new entry in the flag set with boolean value.
// The environment value is overridden with the flag value if it exists
func Bool(flagName, envName string, value bool, usage string) *bool {

	verifyNames(flagName, envName)
	flagValExists := lookupFlag(flagName)
	envValExists := lookupEnv(envName)

	if envValExists {
		envValStr := os.Getenv(envName)
		envVal, _ := strconv.ParseBool(envValStr)
		if flagValExists {
			flag.Set(flagName, envValStr)
		} else {
			flag.Bool(flagName, envVal, usage)
		}
		return &envVal
	}
	goflag.Bool(flagName, value, usage)
	return flag.Bool(flagName, value, usage)
}

// Float64 creates a new entry in the flag set with float64 value.
// The environment value is overridden with the flag value if it exists
func Float64(flagName, envName string, value float64, usage string) *float64 {

	verifyNames(flagName, envName)
	flagValExists := lookupFlag(flagName)
	envValExists := lookupEnv(envName)

	if envValExists {
		envValStr := os.Getenv(envName)
		envVal, _ := strconv.ParseFloat(envValStr, BitSize)
		if flagValExists {
			flag.Set(flagName, envValStr)
		} else {
			flag.Float64(flagName, envVal, usage)
		}
		return &envVal
	}
	goflag.Float64(flagName, value, usage)
	return flag.Float64(flagName, value, usage)
}

// Int creates a new entry in the flag set with Int value.
// The environment value is overridden with the flag value if it exists
func Int(flagName, envName string, value int, usage string) *int {

	verifyNames(flagName, envName)
	flagValExists := lookupFlag(flagName)
	envValExists := lookupEnv(envName)

	if envValExists {
		envValStr := os.Getenv(envName)
		envVal, _ := strconv.ParseInt(envValStr, BaseNum, BitSize)
		if flagValExists {
			flag.Set(flagName, envValStr)
		} else {
			flag.Int(flagName, int(envVal), usage)
		}
		envValInt := int(envVal)
		return &envValInt
	}
	goflag.Int(flagName, value, usage)
	return flag.Int(flagName, value, usage)
}

// Int64 creates a new entry in the flag set with Int64 value.
// The environment value is overridden with the flag value if it exists
func Int64(flagName, envName string, value int64, usage string) *int64 {

	verifyNames(flagName, envName)
	flagValExists := lookupFlag(flagName)
	envValExists := lookupEnv(envName)

	if envValExists {
		envValStr := os.Getenv(envName)
		envVal, _ := strconv.ParseInt(envValStr, BaseNum, BitSize)
		if flagValExists {
			flag.Set(flagName, envValStr)
		} else {
			flag.Int64(flagName, envVal, usage)
		}
		return &envVal
	}
	goflag.Int64(flagName, value, usage)
	return flag.Int64(flagName, value, usage)
}

// String creates a new entry in the flag set with String value.
// The environment value is overridden with the flag value if it exists
func String(flagName, envName, value, usage string) *string {

	verifyNames(flagName, envName)
	flagValExists := lookupFlag(flagName)
	envValExists := lookupEnv(envName)

	if envValExists {
		envVal := os.Getenv(envName)
		if flagValExists {
			flag.Set(flagName, envVal)
		} else {
			flag.String(flagName, envVal, usage)
		}
		return &envVal
	}
	goflag.String(flagName, value, usage)
	return flag.String(flagName, value, usage)
}

// UInt creates a new entry in the flag set with UInt value.
// The environment value is overridden with the flag value if it exists
func UInt(flagName, envName string, value uint, usage string) *uint {

	verifyNames(flagName, envName)
	flagValExists := lookupFlag(flagName)
	envValExists := lookupEnv(envName)

	if envValExists {
		envValStr := os.Getenv(envName)
		envVal, _ := strconv.ParseUint(envValStr, BaseNum, BitSize)
		envValUint := uint(envVal)
		if flagValExists {
			flag.Set(flagName, envValStr)
		} else {
			flag.Uint(flagName, envValUint, usage)
		}
		return &envValUint
	}
	goflag.Uint(flagName, value, usage)
	return flag.Uint(flagName, value, usage)
}

// UInt64 creates a new entry in the flag set with UInt64 value.
// The environment value is overridden with the flag value if it exists
func UInt64(flagName, envName string, value uint64, usage string) *uint64 {

	verifyNames(flagName, envName)
	flagValExists := lookupFlag(flagName)
	envValExists := lookupEnv(envName)

	if envValExists {
		envValStr := os.Getenv(envName)
		envVal, _ := strconv.ParseUint(envValStr, BaseNum, BitSize)
		if flagValExists {
			flag.Set(flagName, envValStr)
		} else {
			flag.Uint64(flagName, envVal, usage)
		}
		return &envVal
	}
	goflag.Uint64(flagName, value, usage)
	return flag.Uint64(flagName, value, usage)
}
