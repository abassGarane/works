package domain

import (
	"errors"
	"time"

	"github.com/teris-io/shortid"
)

var (
	ErrNoJobListed     = errors.New("No job found")
	ErrInvalidJob      = errors.New("Invalid job")
	ErrJobArchived     = errors.New("Job has been archived")
	ErrCompanyDelisted = errors.New("Company has been delisted")
)

type jobService struct {
	jobRepo JobRepository
}

func NewJobService(repo JobRepository) JobService {
	return &jobService{
		jobRepo: repo,
	}
}

func (j *jobService) Get(ID string) (*Job, error) {
	return j.jobRepo.Get(ID)
}
func (j *jobService) GetAll() ([]Job, error) {
	return j.jobRepo.GetAll()
}
func (j *jobService) AddJob(job *Job) error {
	ID := shortid.MustGenerate()
	job.ID = ID
	job.CreatedAt = time.Now().UTC().Unix()
	return j.jobRepo.AddJob(job)
}
func (j *jobService) UpdateJob(job *Job, id string) (*Job, error) {
	return j.jobRepo.UpdateJob(job, id)
}
func (j *jobService) DeleteJob(id string) (*Job, error) {
	return j.jobRepo.DeleteJob(id)
}
