package entitytypes

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = NicEntity{}

type NicEntity struct {
	Properties *NicEntityProperties `json:"properties,omitempty"`

	// Fields inherited from Entity
	Id         *string     `json:"id,omitempty"`
	Name       *string     `json:"name,omitempty"`
	SystemData *SystemData `json:"systemData,omitempty"`
	Type       *string     `json:"type,omitempty"`
}

var _ json.Marshaler = NicEntity{}

func (s NicEntity) MarshalJSON() ([]byte, error) {
	type wrapper NicEntity
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling NicEntity: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling NicEntity: %+v", err)
	}
	decoded["kind"] = "Nic"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling NicEntity: %+v", err)
	}

	return encoded, nil
}
