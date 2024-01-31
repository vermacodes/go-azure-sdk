package migratepostgresqlazuredbforpostgresqlsynctasks

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ MigratePostgreSqlAzureDbForPostgreSqlSyncTaskOutput = MigratePostgreSqlAzureDbForPostgreSqlSyncTaskOutputMigrationLevel{}

type MigratePostgreSqlAzureDbForPostgreSqlSyncTaskOutputMigrationLevel struct {
	EndedOn             *string                  `json:"endedOn,omitempty"`
	SourceServer        *string                  `json:"sourceServer,omitempty"`
	SourceServerType    *ScenarioSource          `json:"sourceServerType,omitempty"`
	SourceServerVersion *string                  `json:"sourceServerVersion,omitempty"`
	StartedOn           *string                  `json:"startedOn,omitempty"`
	State               *ReplicateMigrationState `json:"state,omitempty"`
	TargetServer        *string                  `json:"targetServer,omitempty"`
	TargetServerType    *ScenarioTarget          `json:"targetServerType,omitempty"`
	TargetServerVersion *string                  `json:"targetServerVersion,omitempty"`

	// Fields inherited from MigratePostgreSqlAzureDbForPostgreSqlSyncTaskOutput
	Id *string `json:"id,omitempty"`
}

var _ json.Marshaler = MigratePostgreSqlAzureDbForPostgreSqlSyncTaskOutputMigrationLevel{}

func (s MigratePostgreSqlAzureDbForPostgreSqlSyncTaskOutputMigrationLevel) MarshalJSON() ([]byte, error) {
	type wrapper MigratePostgreSqlAzureDbForPostgreSqlSyncTaskOutputMigrationLevel
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling MigratePostgreSqlAzureDbForPostgreSqlSyncTaskOutputMigrationLevel: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling MigratePostgreSqlAzureDbForPostgreSqlSyncTaskOutputMigrationLevel: %+v", err)
	}
	decoded["resultType"] = "MigrationLevelOutput"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling MigratePostgreSqlAzureDbForPostgreSqlSyncTaskOutputMigrationLevel: %+v", err)
	}

	return encoded, nil
}
