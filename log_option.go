package golog

import (
)


const defaultLevel = "Info";
const defaultTime = false;
const defaultColor = false;
const defaultKeep = false;
const defaultStrict = false;

// Data structure for logs options, define the behaviour of the log or log chain. 
type LogOption struct {
  // Default level for the log.
  Level         string
  // Display time in the log.
  Time          bool
  // Enable colors (Not implemented in v1.0).
  Color         bool
  // The log will be kept instead of destroyed from chain uppon writting.
  Keep          bool
  // All following logs will have the same options: except for Keep that can be changed at will.
  Strict        bool
}
