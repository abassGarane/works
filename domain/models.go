package domain

type Job struct {
	ID               string `json:"id" bson:"id"`
	Title            string `json:"title" bson:"title"`
	Remote           bool   `json:"remote" bson:"remote"`
	Company          string `json:"company" bson:"company"`
	Responsibilities string `json:"responsibilities" bson:"responsibilities"`
	Requirements     string `json:"requirements" bson:"requirements"`
	Offerings        string `json:"offerings" bson:"offerings"`
	Type             string `json:"type" bson:"type"`
	Experience       string `json:"experience" bson:"experience"`
	CreatedAt        int64  `json:"created_at" bson:"created_at"`
}

type Company struct {
	ID       string            `json:"id" bson:"id"`
	Name     string            `json:"name" bson:"name"`
	Jobs     []*Job            `json:"jobs" bson:"jobs"`
	Salaries map[string]string `json:"salaries" bson:"salaries"`
	Location string            `json:"location" bson:"location"`
}

// type JobType int

// const (
// 	FullTime JobType = iota
// 	Contract
// 	Internship
// 	PartTime
// 	Temporary
// )

// type ExperienceLevel int

// const (
// 	MidLevel ExperienceLevel = iota
// 	EntryLevel
// 	Senior
// 	NoExperience
// )
