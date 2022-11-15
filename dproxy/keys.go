/*
  types - © 2018-Present - SouthWinds Tech Ltd - www.southwinds.io
  Licensed under the Apache License, Version 2.0 at http://www.apache.org/licenses/LICENSE-2.0
  Contributors to this project, hereby assign copyright in this code to the project,
  to be licensed under the same terms as the rest of the code.
*/

package dproxy

import (
	"time"
)

const (
	WebHookInfoType = "DPROXY_WH_INFO"
	ReleaseType     = "DPROXY_RELEASE"
	ReleaseKey      = "DPROXY_RELEASE_?"
)

var (
	WebHookInfoProto = &WebhookInfo{
		WebhookToken: "dwfcj3485j353",
		ReferrerURL:  "referrer.com",
		IpSafeList:   []string{"127.0.0.1"},
		Filter:       "",
	}
	ReleaseProto = &Release{
		Origin:       "origin.com",
		DeploymentId: "1234",
		BucketName:   "bucket",
		FolderName:   "v1",
		Time:         time.Time{},
	}
)
