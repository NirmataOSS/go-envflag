package envflag

import (
	"os"
	"testing"
	"time"
)

func TestFlagString(t *testing.T) {
	flagS1 := String("S1", "S1", "hello1", "String")
	Parse()
	if *flagS1 != "hello1" {
		t.Errorf("Error: flagS1 expected %v got %v", "hello1", *flagS1)
		t.Fail()
	}
}

func TestEnvFlagString(t *testing.T) {
	os.Setenv("S2", "hello2")
	flagS2 := String("S2", "S2", "hello2", "String")
	Parse()
	if *flagS2 != "hello2" {
		t.Errorf("Error: flagS2 expected %v got %v", "hello2", *flagS2)
		t.Fail()
	}
}

func TestFlagBool(t *testing.T) {
	flagB1 := Bool("B1", "B1", false, "Bool [BOOL]")
	Parse()
	if *flagB1 != false {
		t.Errorf("Error: flagB1 expected %v got %v", false, *flagB1)
		t.Fail()
	}
}

func TestEnvFlagBool(t *testing.T) {
	os.Setenv("B2", "true")
	flagB2 := Bool("B2", "B2", false, "Bool [BOOL]")
	Parse()
	if *flagB2 != true {
		t.Errorf("Error: flagB2 expected %v got %v", true, *flagB2)
		t.Fail()
	}
}

func TestFlagInt(t *testing.T) {
	flagI1 := Int("I1", "I1", 5, "Int")
	Parse()
	if *flagI1 != 5 {
		t.Errorf("Error: flagI1 expected %v got %v", 5, *flagI1)
		t.Fail()
	}
}

func TestFlagIntEnv(t *testing.T) {
	os.Setenv("I2", "15")
	flagI2 := Int("I2", "I2", 5, "Int")
	Parse()
	if *flagI2 != 15 {
		t.Errorf("Error: flagI2 expected %v got %v", 15, *flagI2)
		t.Fail()
	}
}

func TestFlagInt64(t *testing.T) {
	flagI3 := Int64("I3", "I3", 501, "Int")
	Parse()
	if *flagI3 != 501 {
		t.Errorf("Error: flagI3 expected %v got %v", 501, *flagI3)
		t.Fail()
	}
}

func TestFlagIntEnv64(t *testing.T) {
	os.Setenv("I4", "333")
	flagI4 := Int("I4", "I4", 3, "Int")
	Parse()
	if *flagI4 != 333 {
		t.Errorf("Error: flagI4 expected %v got %v", 333, *flagI4)
		t.Fail()
	}
}

func TestFlagUInt(t *testing.T) {
	flagUI := UInt64("UI", "UI", 1235, "UInt")
	Parse()
	if *flagUI != 1235 {
		t.Errorf("Error: flagUI expected %v got %v", 1235, *flagUI)
		t.Fail()
	}
}

func TestFlagUIntEnv(t *testing.T) {
	os.Setenv("UI2", "657575")
	flagUI2 := Int("UI2", "UI2", 0, "UInt")
	Parse()
	if *flagUI2 != 657575 {
		t.Errorf("Error: flagUI2 expected %v got %v", 657575, *flagUI2)
		t.Fail()
	}
}

func TestFlagUInt64(t *testing.T) {
	flagUI3 := UInt64("UI3", "UI3", 919, "UInt64")
	Parse()
	if *flagUI3 != 919 {
		t.Errorf("Error: flagUI3 expected %v got %v", 919, *flagUI3)
		t.Fail()
	}
}

func TestFlagUIntEnv64(t *testing.T) {
	os.Setenv("UI4", "3232")
	flagUI4 := Int("UI4", "UI4", 0, "UInt64")
	Parse()
	if *flagUI4 != 3232 {
		t.Errorf("Error: flagUI4 expected %v got %v", 3232, *flagUI4)
		t.Fail()
	}
}

func TestFlagFloat64(t *testing.T) {
	flagF64 := Float64("F64", "F64", 65.12, "Float64")
	Parse()
	if *flagF64 != 65.12 {
		t.Errorf("Error: flagF64 expected %v got %v", 65.12, *flagF64)
		t.Fail()
	}
}

func TestFlagFloatEnv64(t *testing.T) {
	os.Setenv("F64-2", "981.282")
	flagF642 := Float64("F64-2", "F64-2", 0.0, "Float64")
	Parse()
	if *flagF642 != 981.282 {
		t.Errorf("Error: flag642 expected %v got %v", 981.282, *flagF642)
		t.Fail()
	}
}

func TestFlagDuration(t *testing.T) {
	d := time.Second * 20
	flagD := Duration("D", "D", d, "Duration")
	Parse()
	if *flagD != d {
		t.Errorf("Error: flagF64 expected %v got %v", d, *flagD)
		t.Fail()
	}
}

func TestFlagFloatEnvDuration(t *testing.T) {
	d := time.Minute * 5
	os.Setenv("D2", d.String())
	flagD2 := Duration("D2", "D2", 0, "Duration")
	Parse()
	if *flagD2 != d {
		t.Errorf("Error: flag642 expected %v got %v", d, *flagD2)
		t.Fail()
	}
}
