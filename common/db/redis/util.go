/*******************************************************************************
* Contributors: BMC Helix, Inc.
*
* (c) Copyright 2020-2025 BMC Helix, Inc.
 
* SPDX-License-Identifier: Apache-2.0
*******************************************************************************/

package redis

import (
	"github.com/gomodule/redigo/redis"
)

func unlinkCollection(conn redis.Conn, col string) error {
	_ = conn.Send("MULTI")
	s := scripts["unlinkZsetMembers"]
	_ = s.Send(conn, col)
	s = scripts["unlinkCollection"]
	_ = s.Send(conn, col)
	_, err := conn.Do("EXEC")
	return err
}
