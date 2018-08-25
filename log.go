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
  // number of bytes readed from the log.
  byteread  int
  // next log of the chain.
  next      *Log
}

func  (l *Log) ResetReader() {
  l.byteread = 0;
}

func  (l *Log) Logify() ([]byte) {
  var logstr = l.GetLevel() + ": " + l.GetMessage() + "\n";

  return []byte(logstr);
}

// Read a log.
func  (l *Log) Read(buffer []byte) (int, error) {
  var bufsize int = len(buffer);
  var log []byte = l.Logify();
  var length = bufsize;
  var sublog []byte = log[l.byteread:];

  // set buffer writing length.
  if len(sublog) < bufsize {
    length = len(sublog);
  }
  for i := 0; i < length; i++ {
    buffer[i] = sublog[i];
  }
  l.byteread += length;
  if l.byteread == len(log) {
    return length, errors.New("EOF");
  }
  return length, nil;
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

  l.byteread = 0;
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
  log.byteread = 0;
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
