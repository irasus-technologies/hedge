/*******************************************************************************
* Contributors: BMC Helix, Inc.
*
* (c) Copyright 2020-2025 BMC Helix, Inc.
 
* SPDX-License-Identifier: Apache-2.0
*******************************************************************************/


package training

import (
	"hedge/edge-ml-service/pkg/dto/job"
)

type JobServiceProvider interface {
	GetTrainingJobStatus(job *job.TrainingJobDetails) error
	SubmitTrainingJob(jobConfig *job.TrainingJobDetails) error
	DownloadModel(localModelDirectory string, fileId string) error
	UploadFile(remoteFile string, localFile string) error
}
