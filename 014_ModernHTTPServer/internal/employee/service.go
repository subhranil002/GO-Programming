package course

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var (
	ErrNotFound = errors.New("employee not found")
	ErrInvalid  = errors.New("invalid employee")
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) List(ctx context.Context) ([]Employee, error) {
	return s.repo.FindAll(ctx)
}

func (s *Service) Get(ctx context.Context, id string) (Employee, error) {
	objectID, err := parseID(id)
	if err != nil {
		return Employee{}, err
	}

	employee, err := s.repo.FindByID(ctx, objectID)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return Employee{}, ErrNotFound
	}
	if err != nil {
		return Employee{}, err
	}

	return employee, nil
}

func (s *Service) Create(
	ctx context.Context,
	req CreateEmployeeRequest,
) (Employee, error) {
	if err := validateCreate(req); err != nil {
		return Employee{}, err
	}

	employee := Employee{
		ID:         bson.NewObjectID(),
		Name:       strings.TrimSpace(req.Name),
		Email:      strings.TrimSpace(req.Email),
		Phone:      strings.TrimSpace(req.Phone),
		Department: strings.TrimSpace(req.Department),
		Salary:     req.Salary,
	}

	return s.repo.Insert(ctx, employee)
}

func (s *Service) Update(
	ctx context.Context,
	id string,
	req UpdateEmployeeRequest,
) (Employee, error) {
	objectID, err := parseID(id)
	if err != nil {
		return Employee{}, err
	}

	update := bson.D{}

	if req.Name != nil {
		name := strings.TrimSpace(*req.Name)
		if name == "" {
			return Employee{}, fmt.Errorf("%w: name cannot be empty", ErrInvalid)
		}
		update = append(update, bson.E{Key: "name", Value: name})
	}

	if req.Email != nil {
		email := strings.TrimSpace(*req.Email)
		if email == "" {
			return Employee{}, fmt.Errorf("%w: email cannot be empty", ErrInvalid)
		}
		update = append(update, bson.E{Key: "email", Value: email})
	}

	if req.Phone != nil {
		update = append(update, bson.E{Key: "phone", Value: strings.TrimSpace(*req.Phone)})
	}

	if req.Department != nil {
		update = append(update, bson.E{Key: "department", Value: strings.TrimSpace(*req.Department)})
	}

	if req.Salary != nil {
		if *req.Salary < 0 {
			return Employee{}, fmt.Errorf("%w: salary cannot be negative", ErrInvalid)
		}
		update = append(update, bson.E{Key: "salary", Value: *req.Salary})
	}

	if len(update) == 0 {
		return Employee{}, fmt.Errorf("%w: no fields to update", ErrInvalid)
	}

	updated, err := s.repo.UpdateByID(ctx, objectID, update)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return Employee{}, ErrNotFound
	}
	if err != nil {
		return Employee{}, err
	}

	return updated, nil
}

func (s *Service) Delete(ctx context.Context, id string) error {
	objectID, err := parseID(id)
	if err != nil {
		return err
	}

	deletedCount, err := s.repo.DeleteByID(ctx, objectID)
	if err != nil {
		return err
	}

	if deletedCount == 0 {
		return ErrNotFound
	}

	return nil
}

func parseID(id string) (bson.ObjectID, error) {
	objectID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return bson.NilObjectID, fmt.Errorf("%w: invalid id", ErrInvalid)
	}

	return objectID, nil
}

func validateCreate(req CreateEmployeeRequest) error {
	if strings.TrimSpace(req.Name) == "" {
		return fmt.Errorf("%w: name is required", ErrInvalid)
	}

	if strings.TrimSpace(req.Email) == "" {
		return fmt.Errorf("%w: email is required", ErrInvalid)
	}

	if req.Salary < 0 {
		return fmt.Errorf("%w: salary cannot be negative", ErrInvalid)
	}

	return nil
}
