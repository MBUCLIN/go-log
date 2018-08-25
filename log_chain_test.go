package golog

import (
  "testing"
  "errors"
  "bytes"
)


func  TestLogChain(t *testing.T) {
  t.Run("LC=Struct", func(t *testing.T) {
    var logChain LogChain = LogChain{
      nil,
      nil,
      nil,
    };
    t.Logf("log chain: %v\n", logChain);
  });
}

func  TestNewLogChain(t *testing.T) {
  var log_lev []string = []string{
    "Warning", "Warning", "Error", "Panic", "Error",
  };
  var log_mes []string = []string{
    "WARNING", "WARNING", "ERROR", "PANIC", "ERROR",
  };
  var log_kee []bool = []bool{
    false, true, true, true, false,
  };

  t.Run("LC=New", func(t *testing.T) {
    var log_chain = NewLogChain("Info", true);

    for i := 0; i < len(log_lev); i++ {
      log_chain.LogString(log_lev[i], log_mes[i], log_kee[i]);
      if log_chain.GetCurrentLevel() != log_lev[i] && log_chain.IsKept() {
        t.Errorf("LC=New: Fail: Invalid level: %s expect %s\n", log_chain.GetCurrentLevel(), log_lev[i]);
        t.Failed();
      }
      if log_chain.GetCurrentMessage() != log_mes[i] {
        t.Errorf("LC=New: Fail: Invalid message: %s expect %s\n", log_chain.GetCurrentMessage(), log_mes[i]);
        t.Failed();
      }
    }
  });
  t.Run("LC=New", func(t *testing.T) {
    var log_chain = NewLogChain("Info", true);

    for i := 0; i < len(log_lev); i++ {
      log_chain.LogError(log_lev[i], errors.New(log_mes[i]), log_kee[i]);
      if log_chain.GetCurrentLevel() != log_lev[i] && log_chain.IsKept() {
        t.Errorf("LC=New: Fail: Invalid level: %s expect %s\n", log_chain.GetCurrentLevel(), log_lev[i]);
        t.Failed();
      }
      if log_chain.GetCurrentMessage() != log_mes[i] {
        t.Errorf("LC=New: Fail: Invalid message: %s expect %s\n", log_chain.GetCurrentMessage(), log_mes[i]);
        t.Failed();
      }
    }
  });
}

func  TestGoThrew(t *testing.T) {
  var log_lev []string = []string{
    "Info", "Warning", "Error", "Panic",
  };
  var log_mes []string = []string{
    "INFO", "WARNING", "ERROR", "PANIC",
  };
  var log_kee []bool = []bool{
    false, true, true, true,
  };
  t.Run("LC=GoThrew", func(t *testing.T) {
    var log_chain *LogChain = NewLogChain("Info", true);

    for i := 0; i < len(log_lev); i++ {
      log_chain.LogString(log_lev[i], log_mes[i], log_kee[i]);
    }
    log_chain.GoThrew();
    for i := 0; i < len(log_lev); i++ {
      if log_chain.GetCurrentLevel() != log_lev[i] {
        t.Errorf("LC=GoThrew: Fail: Current level: %s expect %s\n", log_chain.GetCurrentLevel(), log_lev[i]);
        t.Failed();
      }
      if log_chain.GetCurrentMessage() != log_mes[i] {
        t.Errorf("LC=GoThrew: Fail: Current message: %s expect %s\n", log_chain.GetCurrentMessage(), log_mes[i]);
        t.Failed();
      }
      _ = log_chain.Next();
    }
  });
}

func  TestReadOneByOne(t *testing.T) {
  var log_lev []string = []string{
    "Info", "Warning", "Error", "Panic",
  };
  var log_mes []string = []string{
    "INFO", "WARNING", "ERROR", "PANIC",
  };
  var log_kee []bool = []bool{
    false, true, true, true,
  };
  var contents [][]byte = [][]byte{
    []byte("Info: INFO\n"),
    []byte("Warning: WARNING\n"),
    []byte("Error: ERROR\n"),
    []byte("Panic: PANIC\n"),
  };
  var content []byte = []byte(nil);
  for i := 0; i < len(contents); i++ {
    content = append(content, contents[i]...);
  }

  t.Run("LC=Read", func(t *testing.T) {
    var log_chain = NewLogChain("Info", false);

    for i := 0; i < len(log_lev); i++ {
      log_chain.LogString(log_lev[i], log_mes[i], log_kee[i]);
    }
    log_chain.GoThrew();

    var rd_content []byte = []byte(nil);
    for {
      buf := make([]byte, 1);
      n, err := log_chain.Read(buf);
      rd_content = append(rd_content, buf[:n]...);
      if err != nil {
        if err.Error() != "EOF" {
          t.Errorf("LC=Read: Fail: error reached: %s expect %s\n", err.Error(), "EOF");
          t.Failed();
        }
        break ;
      }
    }
    if bytes.Equal(content, rd_content) != true {
      t.Errorf("LC=Read: Fail: read content: \n%s\nexpect: \n%s\n", rd_content, content);
      t.Errorf("LC=Read: Fail: read length: %d expect %d\n", len(rd_content), len(content));
      t.Failed();
    }
  });
}

func  TestReadSmallSize(t *testing.T) {
  var log_lev []string = []string{
    "Info", "Warning", "Error", "Panic",
  };
  var log_mes []string = []string{
    "INFO", "WARNING", "ERROR", "PANIC",
  };
  var log_kee []bool = []bool{
    false, true, true, true,
  };
  var contents [][]byte = [][]byte{
    []byte("Info: INFO\n"),
    []byte("Warning: WARNING\n"),
    []byte("Error: ERROR\n"),
    []byte("Panic: PANIC\n"),
  };
  var content []byte = []byte(nil);
  for i := 0; i < len(contents); i++ {
    content = append(content, contents[i]...);
  }

  t.Run("LC=Read", func(t *testing.T) {
    var log_chain = NewLogChain("Info", true);

    for i := 0; i < len(log_lev); i++ {
      log_chain.LogString(log_lev[i], log_mes[i], log_kee[i]);
    }
    log_chain.GoThrew();

    var rd_content []byte = []byte(nil);
    for {
      buf := make([]byte, 32);
      t.Logf(string(buf));
      n, err := log_chain.Read(buf);
      rd_content = append(rd_content, buf[:n]...);
      if err != nil {
        if err.Error() != "EOF" {
          t.Errorf("LC=Read: Fail: error reached: %s expect %s\n", err.Error(), "EOF");
          t.Failed();
        }
        break ;
      }
    }
    if bytes.Equal(content, rd_content) != true {
      t.Errorf("LC=Read: Fail: read content: \n%s\nexpect: \n%s\n", rd_content, content);
      t.Errorf("LC=Read: Fail: read length: %d expect %d\n", len(rd_content), len(content));
      t.Failed();
    }
  });
}

func TestReadBigSize(t *testing.T) {
  var log_lev []string = []string{
    "Info", "Warning", "Error", "Panic",
  };
  var log_mes []string = []string{
    "INFO", "WARNING", "ERROR", "PANIC",
  };
  var log_kee []bool = []bool{
    false, true, true, true,
  };
  var contents [][]byte = [][]byte{
    []byte("Info: INFO\n"),
    []byte("Warning: WARNING\n"),
    []byte("Error: ERROR\n"),
    []byte("Panic: PANIC\n"),
  };
  var content []byte = []byte(nil);
  for i := 0; i < len(contents); i++ {
    content = append(content, contents[i]...);
  }

  t.Run("LC=Read", func(t *testing.T) {
    var log_chain = NewLogChain("Info", true);

    for i := 0; i < len(log_lev); i++ {
      log_chain.LogString(log_lev[i], log_mes[i], log_kee[i]);
    }
    log_chain.GoThrew();

    var rd_content []byte = []byte(nil);
    for {
      buf := make([]byte, 4096);
      t.Logf(string(buf));
      n, err := log_chain.Read(buf);
      rd_content = append(rd_content, buf[:n]...);
      if err != nil {
        if err.Error() != "EOF" {
          t.Errorf("LC=Read: Fail: error reached: %s expect %s\n", err.Error(), "EOF");
          t.Failed();
        }
        break ;
      }
    }
    if bytes.Equal(content, rd_content) != true {
      t.Errorf("LC=Read: Fail: read content: \n%s\nexpect: \n%s\n", rd_content, content);
      t.Errorf("LC=Read: Fail: read length: %d expect %d\n", len(rd_content), len(content));
      t.Failed();
    }
  });
}
