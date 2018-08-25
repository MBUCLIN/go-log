package golog

import (
  "testing"
  "os"
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
    };

    t.Logf("option: %v\n", option);
  });
}

// Test NewLogOption function.
func  TestNewLogOption(t *testing.T) {
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

// Test duplicate an option
func  TestDuplicate(t *testing.T) {
  t.Run("LO=Duplicate", func(t *testing.T) {
    var option *LogOption = newOption();

    dup := option.Duplicate();

    dup.Level = "Warning";
    if option.Level == dup.Level {
      t.Errorf("LO=Duplicate: Fail: option.Level: %s expect Info\n", option.Level);
      t.Failed();
    }
  });
}

func  TestLogChainOption(t *testing.T) {
  t.Run("LCO=Struct", func(t *testing.T) {
    var option LogChainOption = LogChainOption{
      false,
      nil,
    };
    t.Logf("log chain option: %v\n", option);
  });
}

func  TestNewLogChainOption(t *testing.T) {
  t.Run("LCO=New", func(t *testing.T) {
    var option *LogChainOption = newLogChainOption(true, nil);

    if option.IsStrict() != true {
      t.Errorf("LCO=New: Fail: Invalid strict mode is disabled\n");
      t.Failed();
    }
  });
}

func  TestLogFileOption(t *testing.T) {
  t.Run("LFO=Struct", func(t *testing.T) {
    var option LogFileOption = LogFileOption{
      true,
      false,
      os.O_RDONLY,
      os.FileMode(0666),
      nil,
      nil,
    };
    t.Logf("log file option: %v\n", option);
  });
}
