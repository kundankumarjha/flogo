package zendesk

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

// THIS IS ADDED
// log is the default package logger which we'll use to log
var log = logger.GetLogger("activity-zendesk")

// MyActivity is a stub for your Activity implementation
type MyActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &MyActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *MyActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// THIS HAS CHANGED
// Eval implements activity.Activity.Eval
func (a *MyActivity) Eval(context activity.Context) (done bool, err error) {

	// Get the activity data from the context
	firstName := context.GetInput("firstName").(string)
	lastName := context.GetInput("lastName").(string)

	// Use the log object to log the greeting
	log.Debugf("The Flogo engine says [%s] to [%s]", firstName, lastName)

	// Set the result as part of the context
	context.SetOutput("name", "The Flogo engine says Your name is : "+firstName+" "+lastName)

	// Signal to the Flogo engine that the activity is completed

	return true, nil
}
