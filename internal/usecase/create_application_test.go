package usecase

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/florian-renfer/freelancing-application-monitor/internal/domain"
	"github.com/google/uuid"
)

type (
	mockApplicationRepository struct {
		domain.ApplicationRepository
		result domain.Application
		err    error
	}

	mockCreateApplicationPresenter struct {
		result CreateApplicationOutput
	}
)

func (m mockApplicationRepository) Create(context.Context, domain.Application) (domain.Application, error) {
	return m.result, m.err
}

func (m mockApplicationRepository) FindAll(context.Context) ([]domain.Application, error) {
	return []domain.Application{}, nil
}

func (m mockCreateApplicationPresenter) Output(_ domain.Application) CreateApplicationOutput {
	return m.result
}

func TestCreateApplication(t *testing.T) {
	t.Parallel()

	type args struct {
		input CreateApplicationInput
	}

	tests := []struct {
		name          string
		args          args
		repository    domain.ApplicationRepository
		presenter     CreateApplicationPresenter
		expected      CreateApplicationOutput
		expectedError any
	}{
		{
			name: "Create account successfully",
			args: args{
				input: CreateApplicationInput{
					Title:       "Fullstack Software Developer",
					Description: "We are looking for a fullstack software developer to support our development team.",
					Url:         "https://freeelancermap.de/projekte/fullstack-software-developer",
					State:       domain.APPLIED,
					AppliedAt:   time.Date(2025, 11, 25, 17, 15, 14, 41, time.UTC),
				},
			},
			repository: mockApplicationRepository{
				result: domain.NewApplication(
					uuid.MustParse("cc4fb575-3d00-4492-ac97-d8d395f6c268"),
					"Fullstack Software Developer",
					"We are looking for a fullstack software developer to support our development team.",
					"https://freeelancermap.de/projekte/fullstack-software-developer",
					domain.APPLIED,
					time.Date(2025, 11, 25, 17, 15, 14, 41, time.UTC),
					time.Date(2025, 11, 25, 17, 19, 32, 46, time.UTC),
					time.Date(2025, 11, 25, 17, 19, 32, 46, time.UTC),
				),
				err: nil,
			},
			presenter: mockCreateApplicationPresenter{
				result: CreateApplicationOutput{
					Id:          uuid.MustParse("cc4fb575-3d00-4492-ac97-d8d395f6c268"),
					Title:       "Fullstack Software Developer",
					Description: "We are looking for a fullstack software developer to support our development team.",
					Url:         "https://freeelancermap.de/projekte/fullstack-software-developer",
					State:       domain.APPLIED,
					AppliedAt:   time.Date(2025, 11, 25, 17, 15, 14, 41, time.UTC),
					CreatedAt:   time.Date(2025, 11, 25, 17, 19, 32, 46, time.UTC),
					UpdatedAt:   time.Date(2025, 11, 25, 17, 19, 32, 46, time.UTC),
				},
			},
			expected: CreateApplicationOutput{
				Id:          uuid.MustParse("cc4fb575-3d00-4492-ac97-d8d395f6c268"),
				Title:       "Fullstack Software Developer",
				Description: "We are looking for a fullstack software developer to support our development team.",
				Url:         "https://freeelancermap.de/projekte/fullstack-software-developer",
				State:       domain.APPLIED,
				AppliedAt:   time.Date(2025, 11, 25, 17, 15, 14, 41, time.UTC),
				CreatedAt:   time.Date(2025, 11, 25, 17, 19, 32, 46, time.UTC),
				UpdatedAt:   time.Date(2025, 11, 25, 17, 19, 32, 46, time.UTC),
			},
			expectedError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var uc = NewCreateApplicationInteractor(tt.repository, tt.presenter, time.Second)

			result, err := uc.Execute(context.TODO(), tt.args.input)
			if (err != nil) && (err.Error() != tt.expectedError) {
				t.Errorf("[TestCase '%s'] Result: '%v' | ExpectedError: '%v'", tt.name, err, tt.expectedError)
			}

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("[TestCase '%s'] Result: '%v' | Expected: '%v'", tt.name, result, tt.expected)
			}
		})
	}
}
