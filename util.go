package sms

import "fmt"

// Constants for responses
const (
	PendingConfirmation = "PendingConfirmation"
	PendingValidation   = "PendingValidation"
	InvalidRequest      = "InvalidRequest"
	NotSupported        = "NotSupported"
	SUCCESS             = "Success"
	FAILED              = "Failed"
	QUEUED              = "Queued"
	SENT                = "Sent"
)

// GetAPIHost returns either sandbox or prod
func GetAPIHost(env string) string {
	return getHost(env, "api")
}

// GetSmsURL is the sms endpoint
func GetSmsURL(env string) string {
	return GetAPIHost(env) + "/version1/messaging"
}

// GetCreateSubURL returns the Subscription create endpoint
func GetCreateSubURL(env string) string {
	return GetAPIHost(env) + "/version1/subscription/create"
}

func getHost(env, service string) string {
	if env != "sandbox" {
		return fmt.Sprintf("https://%s.africastalking.com", service)
	}
	return fmt.Sprintf("https://%s.sandbox.africastalking.com", service)

}
