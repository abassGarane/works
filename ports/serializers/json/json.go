package json

import (
	"encoding/json"

	"github.com/abassGarane/work/domain"
	"github.com/pkg/errors"
)

type Job struct{}

func (j *Job) Decode(input []byte) (*domain.Job, error) {
	job := &domain.Job{}
	err := json.Unmarshal(input, &job)
	if err != nil {
		return nil, errors.Wrap(err, "serializers.Job.Decode")
	}
	return job, nil
}

func (j *Job) Encode(input *domain.Job) ([]byte, error) {
	return json.Marshal(input)
}
