package golog

import(
  "os"
)

// Define a log file manager
type  LogFile struct {
  // The log chains used by the log file.
  Chains  map[string]*LogChain
  // The files link for each chains
  // The file for a chain does not exists if the chain has never been written to file.
  Files   map[string]*os.File
  // The options for the log file
  option  *LogFileOption
}

