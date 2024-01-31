package migratesqlserversqlmitasks

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MigrateSqlServerSqlMITaskOutput interface {
}

// RawMigrateSqlServerSqlMITaskOutputImpl is returned when the Discriminated Value
// doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawMigrateSqlServerSqlMITaskOutputImpl struct {
	Type   string
	Values map[string]interface{}
}

func unmarshalMigrateSqlServerSqlMITaskOutputImplementation(input []byte) (MigrateSqlServerSqlMITaskOutput, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling MigrateSqlServerSqlMITaskOutput into map[string]interface: %+v", err)
	}

	value, ok := temp["resultType"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "AgentJobLevelOutput") {
		var out MigrateSqlServerSqlMITaskOutputAgentJobLevel
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MigrateSqlServerSqlMITaskOutputAgentJobLevel: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "DatabaseLevelOutput") {
		var out MigrateSqlServerSqlMITaskOutputDatabaseLevel
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MigrateSqlServerSqlMITaskOutputDatabaseLevel: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "ErrorOutput") {
		var out MigrateSqlServerSqlMITaskOutputError
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MigrateSqlServerSqlMITaskOutputError: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "LoginLevelOutput") {
		var out MigrateSqlServerSqlMITaskOutputLoginLevel
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MigrateSqlServerSqlMITaskOutputLoginLevel: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "MigrationLevelOutput") {
		var out MigrateSqlServerSqlMITaskOutputMigrationLevel
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MigrateSqlServerSqlMITaskOutputMigrationLevel: %+v", err)
		}
		return out, nil
	}

	out := RawMigrateSqlServerSqlMITaskOutputImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
