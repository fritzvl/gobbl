package gbl

// EVLog contains the string for the log event type
var EVLog = "log"

// EVFlagSet contains the string for the flag set event type
var EVFlagSet = "flag_set"

// EVFlagAccess contains the string for the flag access event type
var EVFlagAccess = "flag_access"

// EVFlagClear contains the string for the flag clear event type
var EVFlagClear = "flag_clear"

// EVRequestStart contains the string for the request start event type
var EVRequestStart = "request_start"

// EVRequestEnd contains the string for the request end event type
var EVRequestEnd = "request_end"

// EVHandlerCall contains the string for the handler call event type
var EVHandlerCall = "handler_call"

// Event represents a single bot event
type Event struct {
	/*
		Event Type
		This notes the type of event that this object contains, can be one of the following
			- log
			- flag_set
			- flag_access
			- flag_clear
			- request_start
			- request_end
			- handler_call
	*/
	Type string `json:"type"`

	// If relevant, this is the context from which the event originated
	Context *Context `json:"ctx"`

	FlagSet    *FlagSet    `json:"fs"`
	FlagAccess *FlagAccess `json:"fa"`
	FlagClear  *FlagClear  `json:"fc"`

	Log *LogEntry `json:"log"`

	HandlerCall *HandlerCall `json:"hc"`
}

// FlagSet contains the data for a flag set event
type FlagSet struct {
	Flag  string      `json:"flag"`
	Value interface{} `json:"value"`
}

// FlagAccess contains the data for a flag access event
type FlagAccess struct {
	Flag             string `json:"flag"`
	IsExistenceCheck bool   `json:"isExistenceCheck"`
}

// FlagClear contains the data for a flag clear event
type FlagClear struct {
	Flags []string `json:"flags"`
}

// HandlerCall contains the data for a handler call event
type HandlerCall struct {
	Handler       string `json:"handler"`
	StackPosition int    `json:"stackPos"`
}
