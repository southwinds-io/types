/*
  types - © 2018-Present - SouthWinds Tech Ltd - www.southwinds.io
  Licensed under the Apache License, Version 2.0 at http://www.apache.org/licenses/LICENSE-2.0
  Contributors to this project, hereby assign copyright in this code to the project,
  to be licensed under the same terms as the rest of the code.
*/

package doorman

import (
	"fmt"
	"strings"
)

// InRoute the definition of an inbound route
type InRoute struct {
	// Name the name of the route
	Name string `bson:"_id" json:"name" yaml:"name" example:"SUPPLIER_A_IN_ROUTE"`
	// Description a description indicating the purpose of the route
	Description string `bson:"description "json:"description" yaml:"description" example:"the inbound route for supplier A"`
	// ServiceHost the remote host from where inbound files should be downloaded
	ServiceHost string `bson:"service_host" json:"service_host" yaml:"service_host" example:"s3.supplier-a.com"`
	// ServiceId a unique identifier for the S3 service where inbound files should be downloaded
	ServiceId string `bson:"service_id" json:"service_id" yaml:"service_id"`
	// BucketName the name of the S3 bucket containing files to download
	BucketName string `bson:"bucket_name" json:"bucket_name" yaml:"bucket_name"`
	// User the username to authenticate against the remote ServiceHost
	User string `bson:"user "json:"user" yaml:"user"`
	// Pwd the password to authenticate against the remote ServiceHost
	Pwd string `bson:"pwd" json:"pwd" yaml:"pwd"`
	// list of authorised publishers
	AllowedAuthors []string `json:"allowed_authors"`
	// Sign whether to (re)sign the imported packages
	Sign bool `json:"sign"`
}

func (r InRoute) GetName() string {
	return r.Name
}

func (r InRoute) Validate() error {
	if len(r.ServiceHost) == 0 {
		return fmt.Errorf("inbound route %s service_host is mandatory", r.Name)
	}
	if !strings.HasPrefix(r.ServiceHost, "s3") {
		return fmt.Errorf("inbound route %s invalid service_host, should start with s3:// or s3s://", r.Name)
	}
	if (len(r.User) > 0 && len(r.Pwd) == 0) || (len(r.User) == 0 && len(r.Pwd) > 0) {
		return fmt.Errorf("inbound route %s requires both username and password to be provided, or none of them", r.Name)
	}
	return nil
}

func (r InRoute) Creds() string {
	return fmt.Sprintf("%s:%s", r.User, r.Pwd)
}
