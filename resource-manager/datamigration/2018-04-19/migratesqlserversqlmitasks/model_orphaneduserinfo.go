package migratesqlserversqlmitasks

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OrphanedUserInfo struct {
	DatabaseName *string `json:"databaseName,omitempty"`
	Name         *string `json:"name,omitempty"`
}
