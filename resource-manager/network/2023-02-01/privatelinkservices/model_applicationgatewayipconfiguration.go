package privatelinkservices

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApplicationGatewayIPConfiguration struct {
	Etag       *string                                            `json:"etag,omitempty"`
	Id         *string                                            `json:"id,omitempty"`
	Name       *string                                            `json:"name,omitempty"`
	Properties *ApplicationGatewayIPConfigurationPropertiesFormat `json:"properties,omitempty"`
	Type       *string                                            `json:"type,omitempty"`
}
