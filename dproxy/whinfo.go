/*
  types - © 2018-Present - SouthWinds Tech Ltd - www.southwinds.io
  Licensed under the Apache License, Version 2.0 at http://www.apache.org/licenses/LICENSE-2.0
  Contributors to this project, hereby assign copyright in this code to the project,
  to be licensed under the same terms as the rest of the code.
*/

package dproxy

// WebhookInfo the information for the proxy to authenticate the Webhook caller
type WebhookInfo struct {
	// WebhookToken the opaque string used for webhook authentication
	WebhookToken string
	// ReferrerURL the URL the referrer of the token should come from to be considered valid
	// this is the URL of the S3 bucket where the spec files were uploaded
	ReferrerURL string
	// IpSafeList one or more IP addresses that are considered valid for authentication of the Webhook caller
	// this is the real IP of the webhook caller
	IpSafeList []string `json:"ip_safe_list" yaml:"ip_safe_list"`
	// Filter a regular expression to filter publication events and prevent the doorman to be invoked if the
	// S3 resource matches the filter; if no defined, no filter is applied
	Filter string
}

func (i WebhookInfo) Validate() error {
	return nil
}
