package golog

import(
  "testing"
)

func  TestLogFile(t *testing.T) {
  t.Run("LF=Struct", func(t *testing.T) {
    var lf LogFile = LogFile{
      nil,
      nil,
      nil,
    };
    t.Logf("Log file: %v\n", lf);
  });
}
