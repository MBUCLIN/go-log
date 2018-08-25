package golog

import (
  "errors"
)


// Data structure for log chains. Used to make it easier to manipulate the log chain.
// A log within the chain canno't be changed.
type LogChain struct {
  // The log chain.
  logs        *Log
  // The current log that the log chain uses.
  current     *Log
  // The chain option for every nodes.
  option      *LogChainOption
}

// Create a new empty log chain.
func  NewLogChain(level string, strictmode bool) (*LogChain) {
  var lc *LogChain = new(LogChain);
  var option *LogOption = NewLogOption(level, false);

  lc.logs = nil;
  lc.current = nil;
  lc.option = newLogChainOption(strictmode, option);
  return lc;
}

// Push the log to the log chain, the current log is set to the pushed log.
func  (lc *LogChain) Add(log *Log) {
  if lc.logs == nil {
    lc.logs = log;
    lc.current = log;
  } else {
    lc.logs.Add(log);
    lc.current.byteread = 0;
    lc.current = log;
  }
}

// Get next log.
func  (lc *LogChain) Next() (error) {
  if lc.current == nil {
    return errors.New("EOF");
  }
  lc.current.byteread = 0;
  next, err := lc.current.Next();
  if err != nil {
    return err;
  }
  lc.current = next;
  return nil;
}

// Get the log option and level from the option.
func  (lc *LogChain) getNewLogOption(level string, keep bool) (*LogOption, string) {
  var option = lc.option.option;
  var keeplevel = option.Level;

  if lc.option.strict == true && option.Keep == false && keep == true {
    var tmpopt *LogOption = NewLogOption(level, keep);

    tmpopt.Level = level;
    tmpopt.Time = option.Time;
    tmpopt.Color = option.Color;
    option = tmpopt;
    keeplevel = level;
  } else if lc.option.strict == false {
    option = option.Duplicate();
    keeplevel = level;
  }
  return option, keeplevel;
}

// Create an new log from an error.
func  (lc *LogChain) LogError(level string, message error, keep bool) {
  option, levelset := lc.getNewLogOption(level, keep);

  var log *Log = NewLog(levelset, message, option);

  lc.Add(log);
}

// Create a log from a string.
func  (lc *LogChain) LogString(level, message string, keep bool) {
  option, levelset := lc.getNewLogOption(level, keep);

  var log *Log = NewLogString(levelset, message, option);

  lc.Add(log);
}

// Get the current log's level.
func (lc *LogChain) GetCurrentLevel() string {
  if lc.current == nil {
    return lc.GetLevel();
  }
  return lc.current.GetLevel();
}

// Get the current log's message.
func (lc *LogChain) GetCurrentMessage() string {
  if lc.current == nil {
    return "";
  }
  return lc.current.GetMessage();
}

// Get the default level of the chain.
func (lc *LogChain) GetLevel() string {
  return lc.option.option.Level;
}

// Get the current log's keep status.
func  (lc *LogChain) IsKept() bool {
  if lc.current == nil {
    return lc.option.option.Keep;
  }
  return lc.current.IsKept();
}

// Set current to the first node of the chain, it indicates that you will go threw the chain.
// The current going threw log is lost at the first call of Add, LogString or LogError.
func  (lc *LogChain) GoThrew() {
  lc.current = lc.logs;
}

// Reads the log chain, can be used as io.Reader
// Call GoThrew before the first call to this function to read from the start.
// Any new nodes added to the chain will break the reading process.
func  (lc *LogChain) Read(buffer []byte) (int, error) {
  var bufsize int = len(buffer);

  if lc.current == nil {
    return 0, errors.New("EOF");
  }
  var readsize int = 0;
  for {
    var subbuf []byte = make([]byte, bufsize - readsize);

    // Read log
    n, err := lc.current.Read(subbuf);
    for i := 0; i < n; i++ {
      buffer[readsize + i] = subbuf[i];
    }
    readsize += n;
    if err != nil {
      if err.Error() != "EOF" {
        return readsize, err;
      }
      // Next log
      err = lc.Next();
      if err != nil {
        return readsize, err;
      }
      if readsize == bufsize {
        break ;
      }
    }
    if readsize == bufsize {
      break ;
    }
  }
  return readsize, nil;
}
