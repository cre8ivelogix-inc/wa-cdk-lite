// Well Architected CDK Constructs - Lite
package wacdklite


// Well Architected Bucket alarms.
type WaBucketAlarms string

const (
	// This enum is used to represent an alarm name for 4XX errors e.g. 404 Not Found.
	WaBucketAlarms_FOURXX_ERRORS_ALARM WaBucketAlarms = "FOURXX_ERRORS_ALARM"
	// This enum is used to represent an alarm name for 5XX errors e.g. 500 Internal Server Error.
	WaBucketAlarms_FIVEXX_ERRORS_ALARM WaBucketAlarms = "FIVEXX_ERRORS_ALARM"
)

