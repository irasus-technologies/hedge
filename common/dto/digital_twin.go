/*******************************************************************************
* Contributors: BMC Helix, Inc.
*
* (c) Copyright 2020-2025 BMC Helix, Inc.
 
* SPDX-License-Identifier: Apache-2.0
*******************************************************************************/


package dto

type DigitalTwin struct {
	DigitalTwinId string
	Scenes        []Scene
}

type Scene struct {
	SceneId       string                   `json:"sceneId"`
	ImageId       string                   `json:"imageId"`
	DeviceId      string                   `json:"deviceId"`
	SceneType     string                   `json:"sceneType"`
	DisplayAttrib []map[string]interface{} `json:"displayAttrib"`
}

type Image struct {
	ImgId        string
	Object       string // scene,profile,device
	ObjName      string
	RemoteFileId string
}

type ChildDevice struct {
	DeviceId    string
	Description string
	Status      string
}
