package golog

import (
)


// Data structure for a log
type Log struct {
  // Contain the log itself.
  Error     error
  // log node level option.
  option    LogOption
  // next log of the chain.
  next      *log
}
