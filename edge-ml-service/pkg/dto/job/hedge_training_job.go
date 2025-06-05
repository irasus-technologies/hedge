/*******************************************************************************
* Contributors: BMC Helix, Inc.
*
* (c) Copyright 2020-2025 BMC Helix, Inc.
 
* SPDX-License-Identifier: Apache-2.0
*******************************************************************************/

package job

type HedgeTrainJob struct {
	JobName   string `json:"job_name"`
	ImagePath string `json:"image_path"`
	DataFile  string `json:"data_file"`
}

type HedgeJobStatus struct {
	JobName string `json:"job_name"`
	Status  string `json:"status"`
	Message string `json:"message"`
}
