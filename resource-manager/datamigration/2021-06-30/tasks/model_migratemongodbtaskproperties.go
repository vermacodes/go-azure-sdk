package tasks

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ProjectTaskProperties = MigrateMongoDbTaskProperties{}

type MigrateMongoDbTaskProperties struct {
	Input  *MongoDbMigrationSettings `json:"input,omitempty"`
	Output *[]MongoDbProgress        `json:"output,omitempty"`

	// Fields inherited from ProjectTaskProperties
	ClientData *map[string]string   `json:"clientData,omitempty"`
	Commands   *[]CommandProperties `json:"commands,omitempty"`
	Errors     *[]ODataError        `json:"errors,omitempty"`
	State      *TaskState           `json:"state,omitempty"`
}

var _ json.Marshaler = MigrateMongoDbTaskProperties{}

func (s MigrateMongoDbTaskProperties) MarshalJSON() ([]byte, error) {
	type wrapper MigrateMongoDbTaskProperties
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling MigrateMongoDbTaskProperties: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling MigrateMongoDbTaskProperties: %+v", err)
	}
	decoded["taskType"] = "Migrate.MongoDb"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling MigrateMongoDbTaskProperties: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &MigrateMongoDbTaskProperties{}

func (s *MigrateMongoDbTaskProperties) UnmarshalJSON(bytes []byte) error {
	type alias MigrateMongoDbTaskProperties
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into MigrateMongoDbTaskProperties: %+v", err)
	}

	s.ClientData = decoded.ClientData
	s.Commands = decoded.Commands
	s.Errors = decoded.Errors
	s.Input = decoded.Input
	s.State = decoded.State

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling MigrateMongoDbTaskProperties into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["output"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Output into list []json.RawMessage: %+v", err)
		}

		output := make([]MongoDbProgress, 0)
		for i, val := range listTemp {
			impl, err := unmarshalMongoDbProgressImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Output' for 'MigrateMongoDbTaskProperties': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Output = &output
	}
	return nil
}
