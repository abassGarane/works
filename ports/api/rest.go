package api

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/abassGarane/work/domain"
	"github.com/gofiber/fiber/v2"
)

type jobHandler struct {
	service domain.JobService
}

func NewJobHandler(jobService domain.JobService) JobHandler {
	return &jobHandler{
		service: jobService,
	}
}

func (j *jobHandler) Get(c *fiber.Ctx) error {
	id := c.Params("id")
	job, err := j.service.Get(id)
	if err != nil {
		return c.Status(fiber.ErrNotFound.Code).JSON(fiber.Map{
			"message": "Job not found",
		})
	}
	return c.Status(fiber.StatusOK).JSON(job)
}

func (j *jobHandler) GetAll(c *fiber.Ctx) error {
	jobs, err := j.service.GetAll()
	log.Println("Request recieved from ", c.OriginalURL(), " ", c.Method())
	if len(jobs) == 0 {
		fmt.Println("No jobs found")
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Jobs not found",
		})
	}
	c.Set("Access-Control-Allow-Origin", "*")
	c.Set("Access-Control-Allow-Headers", "Content-Type")
	if err != nil {
		return c.Status(fiber.ErrNotFound.Code).JSON(fiber.Map{
			"message": "Jobs not found",
		})
	}
	return c.Status(fiber.StatusOK).JSON(&jobs)
}

func (j *jobHandler) UpdateJob(c *fiber.Ctx) error {
	id := c.Params("id")
	job := &domain.Job{}
	//deserialization
	if err := c.BodyParser(&job); err != nil {
		return c.Status(fiber.ErrNotFound.Code).JSON(fiber.Map{
			"message": "Invalid job structure",
		})
	}
	//db
	job, err := j.service.UpdateJob(job, id)
	if err != nil {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{
			"message": "Invalid job structure",
		})
	}
	return c.Status(fiber.StatusOK).JSON(job)
}

func (j *jobHandler) AddJob(c *fiber.Ctx) error {
	job := &domain.Job{}
	rawBody := c.Body()
	if len(rawBody) == 0 {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{
			"message": fmt.Sprintf("Error %s", fiber.ErrInternalServerError.Error()),
		})
	}
	err := json.Unmarshal(rawBody, &job)
	if err := c.BodyParser(&job); err != nil {
		return c.Status(fiber.ErrNotFound.Code).JSON(fiber.Map{
			"message": "Invalid job structure",
		})
	}
	err = j.service.AddJob(job)
	if err != nil {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{
			"message": fiber.ErrInternalServerError.Message,
		})
	}
	return c.SendStatus(fiber.StatusOK)
}

func (j *jobHandler) DeleteJob(c *fiber.Ctx) error {
	id := c.Params("id")
	job, err := j.service.DeleteJob(id)
	if err != nil {
		return c.Status(fiber.ErrNotFound.Code).JSON(fiber.Map{
			"message": "Invalid job structure",
		})
	}
	return c.Status(fiber.StatusAccepted).JSON(job)
}

// func middleware for cors
func (j *jobHandler) SetCorsHeaders(c *fiber.Ctx) fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Set("Access-Control-Allow-Origin", "*")
		return c.Next()
	}
}
