package usecase

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/florian-renfer/freelancing-application-monitor/internal/domain"
	"github.com/google/uuid"
)

type mockUserRepoStore struct {
	domain.UserRepository

	result domain.User
	err    error
}

func (m mockUserRepoStore) Create(_ context.Context, _ domain.User) (domain.User, error) {
	return m.result, m.err
}

type mockCreateUserPresenter struct {
	result CreateUserOutput
}

func (m mockCreateUserPresenter) Output(_ domain.User) CreateUserOutput {
	return m.result
}

func TestExecute(t *testing.T) {
	t.Parallel()

	type args struct {
		input CreateUserInput
	}

	tests := []struct {
		name          string
		args          args
		repository    domain.UserRepository
		presenter     CreateUserPresenter
		expected      CreateUserOutput
		expectedError any
	}{
		{
			name: "Create user successfully",
			args: args{
				input: CreateUserInput{
					Email:       "max.mustermann@mail.de",
					Password:    "changeme",
					Firstname:   "Max",
					Lastname:    "Mustermann",
					DateOfBirth: time.Date(1, 1, 1970, 0, 0, 0, 0, time.UTC),
				},
			},
			repository: mockUserRepoStore{
				result: domain.NewUser(
					uuid.MustParse("080bdfc8-f03b-45fa-8088-5ca94b1502e8"),
					"max.mustermann@mail.de",
					"",
					"Max",
					"Mustermann",
					time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
					time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
					time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
				),
				err: nil,
			},
			presenter: mockCreateUserPresenter{
				result: CreateUserOutput{
					ID:          uuid.MustParse("080bdfc8-f03b-45fa-8088-5ca94b1502e8"),
					Email:       "max.mustermann@mail.de",
					Firstname:   "Max",
					Lastname:    "Mustermann",
					DateOfBirth: time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
					CreatedAt:   time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
					UpdatedAt:   time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
				},
			},
			expected: CreateUserOutput{
				ID:          uuid.MustParse("080bdfc8-f03b-45fa-8088-5ca94b1502e8"),
				Email:       "max.mustermann@mail.de",
				Firstname:   "Max",
				Lastname:    "Mustermann",
				DateOfBirth: time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
				CreatedAt:   time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
				UpdatedAt:   time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			expectedError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var uc = NewCreateUserInteractor(tt.repository, tt.presenter, time.Second)

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
