package golog

import (
  "os"
)

/*
** This part is for a log option.
** The options that are log relative.
*/

const defaultLevel = "Info";
const defaultTime = false;
const defaultColor = false;
const defaultKeep = false;
const defaultStrict = false;

// Data structure for logs options, define the behaviour of the log. 
type LogOption struct {
  // Default level for the log.
  Level             string
  // Display time in the log.
  Time              bool
  // Enable colors (Not implemented in v1.0).
  Color             bool
  // The log will be kept instead of destroyed from chain uppon writting.
  Keep              bool
  // If the log has been kept, we don't want to rewrite it later on.
  written           bool
}

// Create a new Option and set the option level default values.
func  newOption() (*LogOption) {
  var option *LogOption = new(LogOption);

  option.Level = defaultLevel;
  option.Time = defaultTime;
  option.Color = defaultColor;
  option.Keep = defaultKeep;
  option.written = false;
  return option;
}

// Create a new LogOption and set it's level and is the log node is kept.
func  NewLogOption(level string, keep bool) (*LogOption) {
  var option *LogOption = newOption();

  option.Level = level;
  option.Keep = keep;
  return option;
}

// Duplicate an option.
func  (lo *LogOption) Duplicate() (*LogOption) {
  option := new(LogOption);

  option.Level = lo.Level;
  option.Time = lo.Time;
  option.Color = lo.Color;
  option.Keep = lo.Keep;
  return option;
}

/*
** From this point, we are in the log chain options.
** The options that define the log chain behavior.
*/

// Data structure for log chain option, define the behavior for a log chain.
type LogChainOption struct {
  // Enable the strict mode
  // The whole chain uses the same option instance.
  // Except for kept records when the shared option Keep value is false.
  strict          bool
  // The default / shared option instance.
  // If strict mode is false and no option is specified for the new log, this option is duplicated.
  option          *LogOption
}

// Create the LogChain option.
func  newLogChainOption(strictmode bool, option *LogOption) (*LogChainOption) {
  var lc_option *LogChainOption = new(LogChainOption);

  lc_option.strict = strictmode;
  if option == nil {
    option = newOption();
  }
  lc_option.option = option;
  return lc_option;
}

// Set time mode
func  (lco *LogChainOption) SetTime(v bool) {
  lco.option.Time = v;
}

// Set color mode
func  (lco *LogChainOption) SetColor(v bool) {
  lco.option.Color = v;
}

// Set keep mode
func  (lco *LogChainOption) SetKeep(v bool) {
  lco.option.Keep = v;
}

// return strict mode
func  (lco *LogChainOption) IsStrict() (bool) {
  return lco.strict;
}
// return keep mode
func  (lco *LogChainOption) IsKept() (bool) {
  return lco.option.Keep;
}

/*
** Here is the log file options.
** defining the log file behavior.
*/

// Define the log file manager behavior.
type  LogFileOption struct {
  // Delete log nodes uppon writting.
  clearOnWrite      bool
  // Only one chain is used for all types.
  OneChainMode      bool
  // the open option
  FileOption        int
  // the file permission
  FilePerm          os.FileMode
  // mapping of log filenames
  FileLevel         map[string]string
  // chains default options.
  option            *LogChainOption
}
