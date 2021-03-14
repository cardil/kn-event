package k8s

import (
	batchv1 "k8s.io/api/batch/v1"
)

// JobRunner will launch a Job and monitor it for completion.
type JobRunner interface {
	Run(batchv1.Job) error
}

// CreateJobRunner will create a JobRunner, or return an error.
func CreateJobRunner(kube Clients) JobRunner {
	return nil
}
