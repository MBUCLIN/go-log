package golog

import (
  "testing"
  "errors"
  "bytes"
)

func  TestLog(t *testing.T) {
  t.Run("L=Struct", func(t *testing.T) {
    var log Log = Log{
      errors.New("YOLO"),
      nil,
      0,
      nil,
    };
    t.Logf("log: %v\n", log);
  });
}

func  TestNewLog(t *testing.T) {
  t.Run("L=New", func(t *testing.T) {
    var log = NewLogString("Warning", "STRING", nil);

    if log.Error == nil {
      t.Errorf("L=New: Fail: error is not set\n");
      t.Failed();
    }
    if log.Error.Error() != "STRING" {
      t.Errorf("L=New: Fail: error has invalid message: %s\n", log.Error.Error());
      t.Failed();
    }
  });
  t.Run("L=New", func(t *testing.T) {
    var log = NewLog("Error", errors.New("STRING"), nil);

    if log.Error == nil {
      t.Errorf("L=New: Fail: error is not set\n");
      t.Failed();
    }
    if log.Error.Error() != "STRING" {
      t.Errorf("L=New: Fail: error has invalid message: %s\n", log.Error.Error());
      t.Failed();
    }
  });
}

func  TestNewLogWithOption(t *testing.T) {
  var option *LogOption = NewLogOption("Warning", true);
  t.Run("L=New", func(t *testing.T) {
    var log = NewLogString("Warning", "STRING", option);

    if log.GetLevel() != "Warning" {
      t.Errorf("L=New: Fail: Wrong log level\n");
      t.Failed();
    }
    if log.IsKept() != true {
      t.Errorf("option: %v\n", option);
      t.Errorf("L=New: Fail: Log should not be kept\n");
      t.Failed();
    }
  });
  t.Run("L=New", func(t *testing.T) {
    var log = NewLog("Error", errors.New("STRING"), option);

    if log.GetLevel() != "Error" {
      t.Errorf("L=New: Fail: Wrong log level\n");
      t.Failed();
    }
    if log.IsKept() != true {
      t.Errorf("L=New: Fail: Log should not be kept\n");
      t.Failed();
    }
  });
}

func  TestNext(t *testing.T) {

  lmap := make(map[string]string);

  lmap["Info"] = "This is an info";
  lmap["Warning"] = "This is a warning";
  lmap["Error"] = "This is an error.";
  lmap["Panic"] = "Good buy and thank you for the fish!";

  t.Run("L=Next", func(t *testing.T) {
    var headLog *Log = nil;

    for level, message := range lmap {
      var option *LogOption = NewLogOption("Warning", false);
      var log *Log = nil;

      log = NewLogString(level, message, option);
      if headLog == nil {
        headLog = log;
      } else {
        headLog.Add(log);
      }
    }
    var tmp *Log = headLog;

    for {
      if tmp == nil {
        break ;
      }
      if tmp.GetMessage() != lmap[tmp.GetLevel()] {
        t.Errorf("L=Next: Fail: invalid message: |%s| expect: |%s|\n", tmp.GetMessage(), lmap[tmp.GetLevel()]);
        t.Failed();
      }
      tmp, _ = tmp.Next();
    }
  });
}

func  TestRead(t *testing.T) {
  var log = NewLogString("Info", "This is my message", NewLogOption("Info", false));
  var bufsize []int = []int{
    1, 16, 128,
  };
  t.Run("L=Read", func(t *testing.T) {

    for i := 0; i < len(bufsize); i++ {
      var content []byte = []byte(nil);
      log.ResetReader();
      for {
        buf := make([]byte, bufsize[i]);
        n, err := log.Read(buf);
        content = append(content, buf[:n]...);
        if err != nil {
          if err.Error() != "EOF" {
            t.Errorf("L=Read: Fail: Unexpected error: %s\n", err.Error());
            t.Failed();
          }
          break ;
        }
      }
      if bytes.Equal(content, log.Logify()) != true {
        t.Errorf("L=Read: Fail: content: %s\nexpect: %s\n", string(content), string(log.Logify()));
        t.Failed();
      }
    }
  });
}
