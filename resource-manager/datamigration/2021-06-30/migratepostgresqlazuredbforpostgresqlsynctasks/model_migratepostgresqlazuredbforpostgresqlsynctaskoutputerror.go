package migratepostgresqlazuredbforpostgresqlsynctasks

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ MigratePostgreSqlAzureDbForPostgreSqlSyncTaskOutput = MigratePostgreSqlAzureDbForPostgreSqlSyncTaskOutputError{}

type MigratePostgreSqlAzureDbForPostgreSqlSyncTaskOutputError struct {
	Error *ReportableException `json:"error,omitempty"`

	// Fields inherited from MigratePostgreSqlAzureDbForPostgreSqlSyncTaskOutput
	Id *string `json:"id,omitempty"`
}

var _ json.Marshaler = MigratePostgreSqlAzureDbForPostgreSqlSyncTaskOutputError{}

func (s MigratePostgreSqlAzureDbForPostgreSqlSyncTaskOutputError) MarshalJSON() ([]byte, error) {
	type wrapper MigratePostgreSqlAzureDbForPostgreSqlSyncTaskOutputError
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling MigratePostgreSqlAzureDbForPostgreSqlSyncTaskOutputError: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling MigratePostgreSqlAzureDbForPostgreSqlSyncTaskOutputError: %+v", err)
	}
	decoded["resultType"] = "ErrorOutput"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling MigratePostgreSqlAzureDbForPostgreSqlSyncTaskOutputError: %+v", err)
	}

	return encoded, nil
}
