package cdksopssecrets

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
)

// Options for expiration notifications on secret keys.
//
// When enabled, CDK reads unencrypted keys ending with the configured suffix
// (e.g. `gitlab_token_expiration`) from the local `sopsFilePath` and
// synthesizes one-time EventBridge Scheduler schedules that publish to SNS
// before each expiration date.
type ExpirationOptions struct {
	// Number of days before the expiration date to send the SNS notification, or multiple reminder offsets to synthesize one schedule per value.
	// Default: 14.
	//
	DaysBeforeExpiration interface{} `field:"optional" json:"daysBeforeExpiration" yaml:"daysBeforeExpiration"`
	// Enable expiration notifications.
	// Default: false.
	//
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// The suffix used to identify expiration date keys in the secret.
	//
	// For example, a suffix of `_expiration` will match any key like
	// `gitlab_token_expiration` and treat its value as the expiration date
	// for `gitlab_token`.
	// Default: '_expiration'.
	//
	ExpirationSuffix *string `field:"optional" json:"expirationSuffix" yaml:"expirationSuffix"`
	// An existing SNS topic to publish expiration notifications to.
	//
	// If not provided, a new SNS topic will be created automatically.
	// Default: - A new SNS topic is created.
	//
	NotificationTopic awssns.ITopic `field:"optional" json:"notificationTopic" yaml:"notificationTopic"`
	// A subscriber to attach to the expiration notification topic.
	//
	// Works for both an auto-created topic and a provided `notificationTopic`.
	// Default: - No subscriber is added.
	//
	Subscriber awssns.ITopicSubscription `field:"optional" json:"subscriber" yaml:"subscriber"`
}

