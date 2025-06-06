/*******************************************************************************
* Contributors: BMC Helix, Inc.
*
* (c) Copyright 2020-2025 BMC Helix, Inc.
 
* SPDX-License-Identifier: Apache-2.0
*******************************************************************************/

package redis

import (
	"encoding/json"
)

type marshalFunc func(in interface{}) (out []byte, err error)
type unmarshalFunc func(in []byte, out interface{}) (err error)

func marshalObject(in interface{}) (out []byte, err error) {
	return json.Marshal(in)
}

func unmarshalObject(in []byte, out interface{}) (err error) {
	return json.Unmarshal(in, out)
}
