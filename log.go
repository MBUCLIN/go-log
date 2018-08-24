package golog

import (
  "errors"
)


// Data structure for a log
type Log struct {
  // Contain the log itself.
  Error     error
  // log node level option.
  option    *LogOption
  // next log of the chain.
  next      *Log
}

// Set the last log
func  (l *Log) Add(next *Log) {
  var tmp *Log = l;

  for tmp.next != nil {
    tmp = tmp.next;
  }
  tmp.next = next;
}

// Return the next node or an error if there is none.
func  (l *Log) Next() (*Log, error) {
  var err error = nil;

  if l.next == nil {
    err = errors.New("EOF");
  }
  return l.next, err;
}

// Get the log error level.
func  (l *Log) GetLevel() (string) {
  return l.option.Level;
}

func  (l *Log) GetMessage() (string) {
  if l.Error != nil {
    return l.Error.Error();
  } else {
    return "";
  }
}
// Return true if log should be kept.
func  (l *Log) IsKept() (bool) {
  return l.option.Keep;
}

// Create a log.
func  newLog(message error, option *LogOption) (*Log) {
  var log *Log = new(Log);

  log.next = nil;
  log.option = option;
  log.Error = message;
  return log;
}

// Create a log exported, this created option if nil.
func  NewLog(level string, message error, option *LogOption) (*Log) {
  if option == nil {
    option = NewLogOption(level, false);
  }
  option.Level = level;
  return newLog(message, option);
}

// Create a log from a string message. There will be no code set to the log.
func  NewLogString(level string, message string, option *LogOption) (*Log) {
  var err error = errors.New(message);

  return NewLog(level, err, option);
}
