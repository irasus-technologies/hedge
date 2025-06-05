/*******************************************************************************
* Contributors: BMC Helix, Inc.
*
* (c) Copyright 2020-2025 BMC Helix, Inc.
 
* SPDX-License-Identifier: Apache-2.0
*******************************************************************************/


package dto

import "fmt"

type UserRole struct {
	UserKongUsername string `json:"userId,omitempty"`
	RoleName         string `json:"role,omitempty"`
}

func (userRole *UserRole) TableName() string {
	return "hedge.user_roles"
}
func (userRole UserRole) ToString() string {
	return fmt.Sprintf("userKongUsername: %s\nroleName: %s\n", userRole.UserKongUsername, userRole.RoleName)
}
