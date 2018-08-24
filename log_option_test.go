package golog

import (
  "testing"
)

// Does not really test anything, just test build.
func  TestLogOption(t *testing.T) {
  t.Run("LO=Struct", func(t *testing.T) {
    var option LogOption = LogOption{
      "Info",
      false,
      false,
      true,
      false,
      false,
    };

    t.Logf("option: %v\n", option);
  });
}

// Test NewLogOption function.
func  TestNewLogOption(t * testing.T) {
  t.Run("LO=New", func(t *testing.T) {
    var option *LogOption = NewLogOption("Warning", true);

    if option.Level != "Warning" {
      t.Errorf("LO=New: Fail: level is not set to Warning but: %s\n", option.Level);
      t.Failed();
    }
    if option.Keep != true {
      t.Errorf("LO=New: Fail: keep boolean is not true\n");
      t.Failed();
    }
  })
}
