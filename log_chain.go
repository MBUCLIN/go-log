package golog

import (
)


// Data structure for log chains. Used to make it easier to manipulate the log chain.
type LogChain struct {
  // The log chain.
  Logs        Log
  // The current log that the log chain uses.
  current     Log
  // The default options for every nodes.
  option      LogOption
}
