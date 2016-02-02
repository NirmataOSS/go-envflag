package envflag

import (
	"os"
	"testing"

	"bitbucket.org/nirmata/go-envflag"
)

const (
	boolValDefaultFlag   = true
	f64ValDefaultFlag    = 3.14
	intValDefaultFlag    = 3140
	int64ValDefaultFlag  = 31400
	uintValDefaultFlag   = 314000
	uint64ValDefaultFlag = 3140000

	boolValDefaultFlagStr   = "true"
	f64ValDefaultFlagStr    = "3.14"
	intValDefaultFlagStr    = "3140"
	int64ValDefaultFlagStr  = "31400"
	uintValDefaultFlagStr   = "314000"
	uint64ValDefaultFlagStr = "3140000"

	boolValDefaultFlag2   = true
	f64ValDefaultFlag2    = 4.13
	intValDefaultFlag2    = 4130
	int64ValDefaultFlag2  = 41300
	uintValDefaultFlag2   = 413000
	uint64ValDefaultFlag2 = 4130000

	boolValDefaultFlag2Str   = "true"
	f64ValDefaultFlag2Str    = "4.13"
	intValDefaultFlag2Str    = "4130"
	int64ValDefaultFlag2Str  = "41300"
	uintValDefaultFlag2Str   = "413000"
	uint64ValDefaultFlag2Str = "4130000"

	boolValDefaultEnv   = false
	f64ValDefaultEnv    = 8.13
	intValDefaultEnv    = 8130
	int64ValDefaultEnv  = 81300
	uintValDefaultEnv   = 813000
	uint64ValDefaultEnv = 8130000

	boolValDefaultEnvStr   = "false"
	f64ValDefaultEnvStr    = "8.13"
	intValDefaultEnvStr    = "8130"
	int64ValDefaultEnvStr  = "81300"
	uintValDefaultEnvStr   = "813000"
	uint64ValDefaultEnvStr = "8130000"
)

/* For testing purposes, we cannot declare these variables here since
we're creating environment variables in the code.  We should call
these functions after declaring the environment variables. Then
call parse

var (
	flagB      = envflag.Bool("B", "B", boolValDefaultFlag, "Bool [BOOL]")
	flagF64    = envflag.Float64("F64", "F64", f64ValDefaultFlag, "F64 [F64]")
	flagInt    = envflag.Int("Int", "Int", intValDefaultFlag, "Int [INT]")
	flagInt64  = envflag.Int64("Int64", "Int64", int64ValDefaultFlag, "Int64 [INT64]")
	flagUint   = envflag.UInt("Uint", "Uint", uintValDefaultFlag, "Uint [UINT]")
	flagUint64 = envflag.UInt64("Uint64", "Uint64", uint64ValDefaultFlag, "Uint64 [UINT64]")
)
*/

func TestEnvFlagBool(t *testing.T) {

	envflag.Set("logtostderr", "true")
	envflag.Set("v", "3")

	if err := os.Setenv("B", boolValDefaultEnvStr); err != nil {
		t.Log("B Setenv err:", err)
		t.Fail()
	}

	flagB := envflag.Bool("B", "B", boolValDefaultFlag, "Bool [BOOL]")
	envflag.Parse()

	if *flagB != boolValDefaultEnv {

		t.Errorf("Error: Unexpected values after processing")
		t.Fail()
	}
}

func TestEnvFlagF64(t *testing.T) {

	envflag.Set("logtostderr", "true")
	envflag.Set("v", "3")

	if err := os.Setenv("F64", f64ValDefaultEnvStr); err != nil {
		t.Log("F64 Setenv err:", err)
		t.Fail()
	}

	flagF64 := envflag.Float64("F64", "F64", f64ValDefaultFlag, "F64 [F64]")

	if *flagF64 != f64ValDefaultEnv {

		t.Errorf("Error: Unexpected values after processing")
		t.Fail()
	}
}

func TestEnvFlagInt(t *testing.T) {

	envflag.Set("logtostderr", "true")
	envflag.Set("v", "3")

	if err := os.Setenv("Int", intValDefaultEnvStr); err != nil {
		t.Log("Int Setenv err:", err)
		t.Fail()
	}

	flagInt := envflag.Int("Int", "Int", intValDefaultFlag, "Int [INT]")
	envflag.Parse()

	if *flagInt != int(intValDefaultEnv) {

		t.Errorf("Error: Unexpected values after processing")
		t.Fail()
	}
}

func TestEnvFlagInt64(t *testing.T) {

	envflag.Set("logtostderr", "true")
	envflag.Set("v", "3")

	if err := os.Setenv("Int64", int64ValDefaultEnvStr); err != nil {
		t.Log("Int64 Setenv err:", err)
		t.Fail()
	}
	flagInt64 := envflag.Int64("Int64", "Int64", int64ValDefaultFlag, "Int64 [INT64]")
	envflag.Parse()

	if *flagInt64 != int64ValDefaultEnv {

		t.Errorf("Error: Unexpected values after processing")
		t.Fail()
	}
}

func TestEnvFlagUint(t *testing.T) {

	envflag.Set("logtostderr", "true")
	envflag.Set("v", "3")

	if err := os.Setenv("Uint", uintValDefaultEnvStr); err != nil {
		t.Log("Uint Setenv err:", err)
		t.Fail()
	}

	flagUint := envflag.UInt("Uint", "Uint", uintValDefaultFlag, "Uint [UINT]")
	envflag.Parse()

	if *flagUint != uint(uintValDefaultEnv) {

		t.Errorf("Error: Unexpected values after processing")
		t.Fail()
	}
}

// TestEnvFlag tests if we can get flags and environment variables,
// and if the environment variables can override the flags
func TestEnvFlagUint64(t *testing.T) {
	envflag.Set("logtostderr", "true")
	envflag.Set("v", "3")

	if err := os.Setenv("Uint64", uint64ValDefaultEnvStr); err != nil {
		t.Log("Uint64 Setenv err:", err)
		t.Fail()
	}

	flagUint64 := envflag.UInt64("Uint64", "Uint64", uint64ValDefaultFlag, "Uint64 [UINT64]")
	envflag.Parse()

	if *flagUint64 != uint64ValDefaultEnv {

		t.Errorf("Error: Unexpected values after processing")
		t.Fail()
	}
}

func TestEnvFlagBoolNoOverride(t *testing.T) {

	envflag.Set("logtostderr", "true")
	envflag.Set("v", "3")

	flagB2 := envflag.Bool("B2", "B2", boolValDefaultFlag2, "Bool2 [BOOL2]")
	envflag.Parse()

	if *flagB2 != boolValDefaultFlag2 {

		t.Errorf("Error: Unexpected values after processing")
		t.Fail()
	}
}

func TestEnvFlagF64NoOverride(t *testing.T) {

	envflag.Set("logtostderr", "true")
	envflag.Set("v", "3")

	flagF642 := envflag.Float64("F642", "F642", f64ValDefaultFlag2, "F642 [F642]")
	envflag.Parse()

	if *flagF642 != f64ValDefaultFlag2 {

		t.Errorf("Error: Unexpected values after processing")
		t.Fail()
	}
}

func TestEnvFlagIntNoOverride(t *testing.T) {

	envflag.Set("logtostderr", "true")
	envflag.Set("v", "3")

	flagInt2 := envflag.Int("Int2", "Int2", intValDefaultFlag2, "Int2 [INT2]")
	envflag.Parse()

	if *flagInt2 != int(intValDefaultFlag2) {

		t.Errorf("Error: Unexpected values after processing")
		t.Fail()
	}
}

func TestEnvFlagInt64NoOverride(t *testing.T) {

	envflag.Set("logtostderr", "true")
	envflag.Set("v", "3")

	flagInt642 := envflag.Int64("Int642", "Int642", int64ValDefaultFlag2, "Int642 [INT642]")
	envflag.Parse()

	if *flagInt642 != int64ValDefaultFlag2 {

		t.Errorf("Error: Unexpected values after processing")
		t.Fail()
	}
}

func TestEnvFlagIntUintNoOverride(t *testing.T) {

	envflag.Set("logtostderr", "true")
	envflag.Set("v", "3")

	flagUint2 := envflag.UInt("Uint2", "Uint2", uintValDefaultFlag2, "Uint2 [UINT2]")
	envflag.Parse()

	if *flagUint2 != uint(uintValDefaultFlag2) {

		t.Errorf("Error: Unexpected values after processing")
		t.Fail()
	}
}

func TestEnvFlagIntUint64NoOverride(t *testing.T) {

	envflag.Set("logtostderr", "true")
	envflag.Set("v", "3")

	flagUint642 := envflag.UInt64("Uint642", "Uint642", uint64ValDefaultFlag2, "Uint642 [UINT642]")
	envflag.Parse()

	if *flagUint642 != uint64ValDefaultFlag2 {

		t.Errorf("Error: Unexpected values after processing")
		t.Fail()
	}
}

func TestPrintDefaults(t *testing.T) {

	envflag.Set("logtostderr", "true")
	envflag.Set("v", "3")

	envflag.UInt("Uint21", "Uint21", uintValDefaultFlag2, "Uint21 [UINT21]")

	envflag.UInt64("Uint6421", "Uint6421", uint64ValDefaultFlag2, "Uint6421 [UINT6421]")
	t.Logf(">>>>> Printing Defaults <<<<<")
	envflag.PrintDefaults()
}

func TestPrintUsage(t *testing.T) {

	envflag.Set("logtostderr", "true")
	envflag.Set("v", "3")

	t.Logf(">>>>> Printing Usage <<<<<")
	envflag.Usage()
}
