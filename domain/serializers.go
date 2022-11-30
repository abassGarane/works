package domain

type JobSerializer interface {
	Decode([]byte) (*Job, error)
	Encode(*Job) ([]byte, error)
}
