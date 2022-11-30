package mongo

import (
	"context"
	"fmt"
	"time"

	"github.com/abassGarane/work/domain"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type mongoRepository struct {
	client   *mongo.Client
	ctx      context.Context
	database string
	timeout  time.Duration
}

func newClient(mongoURL string, timeout int, ctx context.Context) (*mongo.Client, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURL))
	if err != nil {
		return nil, err
	}
	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}
	return client, err
}
func NewMongoRepository(mongoURL, database string, timeout int, ctx context.Context) (domain.JobRepository, error) {
	client, err := newClient(mongoURL, timeout, ctx)
	if err != nil {
		return nil, err
	}
	repo := &mongoRepository{
		client:   client,
		database: database,
		ctx:      ctx,
		timeout:  time.Duration(timeout) * time.Second,
	}
	return repo, nil
}
func (m *mongoRepository) Get(ID string) (*domain.Job, error) {
	col := m.client.Database(m.database).Collection("jobs")
	filter := bson.M{
		"id": ID,
	}
	job := &domain.Job{}
	err := col.FindOne(m.ctx, filter).Decode(&job)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.Wrap(domain.ErrNoJobListed, "repository.Job.Get")
		}
		return nil, errors.Wrap(err, "repository.Job.Get")
	}
	return job, nil
}
func (m *mongoRepository) GetAll() ([]domain.Job, error) {

	col := m.client.Database(m.database).Collection("jobs")
	cursor, err := col.Find(m.ctx, bson.D{})
	if err != nil {
		if err == mongo.ErrNilDocument {
			return nil, errors.Wrap(domain.ErrNoJobListed, "repository.Job.GetAll")
		}

		return nil, errors.Wrap(err, "repository.Job.GetAll")
	}
	jobs := []domain.Job{}
	defer cursor.Close(m.ctx)
	if err = cursor.All(m.ctx, &jobs); err != nil {
		return nil, errors.Wrap(err, "repository.Job.AddJob")
	}

	return jobs, nil
}
func (m *mongoRepository) AddJob(j *domain.Job) error {
	col := m.client.Database(m.database).Collection("jobs")
	_, err := col.InsertOne(m.ctx, &j)
	if err != nil {
		return errors.Wrap(domain.ErrInvalidJob, "repository.Job.AddJob")
	}
	return nil
}
func (m *mongoRepository) UpdateJob(j *domain.Job, id string) (*domain.Job, error) {
	col := m.client.Database(m.database).Collection("jobs")
	job := &domain.Job{}
	if err := col.FindOneAndUpdate(m.ctx, bson.M{"id": id}, j).Decode(&job); err != nil {
		return nil, errors.Wrap(err, "repository.Job.UpdateJob")
	}
	return job, nil
}
func (m *mongoRepository) DeleteJob(id string) (*domain.Job, error) {
	col := m.client.Database(m.database).Collection("jobs")
	j := &domain.Job{}
	err := col.FindOneAndDelete(m.ctx, bson.M{"id": id}).Decode(&j)
	if err != nil {
		return nil, errors.Wrap(err, "repository.Job.DeleteJob")
	}
	return j, nil
}
