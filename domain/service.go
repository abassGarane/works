package domain

type JobService interface {
	Get(ID string) (*Job, error)
	GetAll() ([]Job, error)
	AddJob(*Job) error
	UpdateJob(j *Job, id string) (*Job, error)
	DeleteJob(id string) (*Job, error)
}
