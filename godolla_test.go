package godolla

import (
  "testing"
)

var gd godolla = godolla{}

func TestSerialWith10(t *testing.T) {
  str := gd.Serial(10)
  if str != "($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)" {
    t.Error("Expected ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10), got ", str)
  }
}

func TestSingleSerial(t *testing.T) {
  str := gd.Serial(1)
  if str != "($1)" {
    t.Error("Expected ($1), got ", str)
  }
}