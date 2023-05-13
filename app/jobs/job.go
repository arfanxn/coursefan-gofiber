package jobs

import (
	"github.com/sirupsen/logrus"
)

type JobHandler interface {
	Handle()
}

type Job struct {
	Payload any
	Action  func() error
}

func (job *Job) Handle() {
	err := job.Action()
	if err != nil {
		logrus.Error(err)
		// TODO: handle error
	}
}
