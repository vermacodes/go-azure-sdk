package tasks

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SelectedCertificateInput struct {
	CertificateName string `json:"certificateName"`
	Password        string `json:"password"`
}
