package action

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/florian-renfer/freelancing-application-monitor/internal/domain"
	"github.com/florian-renfer/freelancing-application-monitor/internal/infastructure/log"
	"github.com/florian-renfer/freelancing-application-monitor/internal/infastructure/validation"
	"github.com/florian-renfer/freelancing-application-monitor/internal/usecase"
	"github.com/google/uuid"
)

const createApplicationEndpoint = "/v1/applications"

type mockCreateApplication struct {
	result usecase.CreateApplicationOutput
	err    error
}

func (m mockCreateApplication) Execute(_ context.Context, _ usecase.CreateApplicationInput) (usecase.CreateApplicationOutput, error) {
	return m.result, m.err
}

func TestCreateApplication(t *testing.T) {
	t.Parallel()

	validator, _ := validation.NewValidatorFactory(validation.InstanceGoPlayground)

	type args struct {
		payload []byte
	}

	tests := []struct {
		name               string
		args               args
		ucMock             usecase.CreateApplicationUseCase
		expectedBody       string
		expectedStatusCode int
	}{
		{
			name: "Create application successfully",
			args: args{
				payload: []byte(`
					{
						"title": "Lead Developer | Java",
						"description": "We're looking for a lead developer for a new Java application.",
						"url": "https://freelancermap.de/projekte/lead-developer-java",
						"state": "APPLIED",
						"applied_at": "2025-12-06T18:28:44.0500+01:00"
					}
				`),
			},
			ucMock: mockCreateApplication{
				result: usecase.CreateApplicationOutput{
					Id:          uuid.MustParse("72325c96-bdea-4d60-a296-f84f2a5e3ef9"),
					Title:       "Lead Developer | Java",
					Description: "We're looking for a lead developer for a new Java application.",
					Url:         "https://freelancermap.de/projekte/lead-developer-java",
					State:       domain.APPLIED,
					AppliedAt:   time.Date(2025, 12, 06, 19, 28, 44, 05, time.UTC),
					CreatedAt:   time.Time{},
					UpdatedAt:   time.Time{},
				},
				err: nil,
			},
			expectedBody:       `{"id":"72325c96-bdea-4d60-a296-f84f2a5e3ef9","title":"Lead Developer | Java","description":"We're looking for a lead developer for a new Java application.","url":"https://freelancermap.de/projekte/lead-developer-java","state":"APPLIED","applied_at":"2025-12-06T19:28:44.000000005Z","created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z"}`,
			expectedStatusCode: http.StatusCreated,
		},
		{
			name: "Create application invalid title (empty)",
			args: args{
				payload: []byte(`
					{
						"title": "",
						"description": "We're looking for a lead developer for a new Java application.",
						"url": "https://freelancermap.de/projekte/lead-developer-java",
						"state": "APPLIED",
						"applied_at": "2025-12-06T18:28:44.0500+01:00"
					}
				`),
			},
			ucMock:             mockCreateApplication{},
			expectedBody:       `{"errors":["Title is a required field"]}`,
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "Create application invalid title (too short)",
			args: args{
				payload: []byte(`
					{
						"title": "Too short",
						"description": "We're looking for a lead developer for a new Java application.",
						"url": "https://freelancermap.de/projekte/lead-developer-java",
						"state": "APPLIED",
						"applied_at": "2025-12-06T18:28:44.0500+01:00"
					}
				`),
			},
			ucMock:             mockCreateApplication{},
			expectedBody:       `{"errors":["Title must be at least 10 characters in length"]}`,
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "Create application invalid title (too long)",
			args: args{
				payload: []byte(`
					{
						"title": "` + strings.Repeat("a", 256) + `",
						"description": "We're looking for a lead developer for a new Java application.",
						"url": "https://freelancermap.de/projekte/lead-developer-java",
						"state": "APPLIED",
						"applied_at": "2025-12-06T18:28:44.0500+01:00"
					}
				`),
			},
			ucMock:             mockCreateApplication{},
			expectedBody:       `{"errors":["Title must be a maximum of 255 characters in length"]}`,
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "Create application invalid description (empty)",
			args: args{
				payload: []byte(`
					{
						"title": "Lead Developer | Java",
						"description": "",
						"url": "https://freelancermap.de/projekte/lead-developer-java",
						"state": "APPLIED",
						"applied_at": "2025-12-06T18:28:44.0500+01:00"
					}
				`),
			},
			ucMock:             mockCreateApplication{},
			expectedBody:       `{"errors":["Description is a required field"]}`,
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "Create application invalid description (too short)",
			args: args{
				payload: []byte(`
					{
						"title": "Lead Developer | Java",
						"description": "Too short",
						"url": "https://freelancermap.de/projekte/lead-developer-java",
						"state": "APPLIED",
						"applied_at": "2025-12-06T18:28:44.0500+01:00"
					}
				`),
			},
			ucMock:             mockCreateApplication{},
			expectedBody:       `{"errors":["Description must be at least 10 characters in length"]}`,
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "Create application invalid description (too long)",
			args: args{
				payload: []byte(`
					{
						"title": "Lead Developer | Java",
						"description": "` + strings.Repeat("a", 3001) + `",
						"url": "https://freelancermap.de/projekte/lead-developer-java",
						"state": "APPLIED",
						"applied_at": "2025-12-06T18:28:44.0500+01:00"
					}
				`),
			},
			ucMock:             mockCreateApplication{},
			expectedBody:       `{"errors":["Description must be a maximum of 3,000 characters in length"]}`,
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "Create application invalid url (not a url)",
			args: args{
				payload: []byte(`
					{
						"title": "Lead Developer | Java",
						"description": "Valid description",
						"url": "not-a-url",
						"state": "APPLIED",
						"applied_at": "2025-12-06T18:28:44.0500+01:00"
					}
				`),
			},
			ucMock:             mockCreateApplication{},
			expectedBody:       `{"errors":["URL must be a valid URL"]}`,
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "Create application invalid url (empty string)",
			args: args{
				payload: []byte(`
					{
						"title": "Lead Developer | Java",
						"description": "Valid description",
						"url": "",
						"state": "APPLIED",
						"applied_at": "2025-12-06T18:28:44.0500+01:00"
					}
				`),
			},
			ucMock:             mockCreateApplication{},
			expectedBody:       `{"errors":["Url is a required field"]}`,
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "Create application valid url (https)",
			args: args{
				payload: []byte(`
					{
						"title": "Lead Developer | Java",
						"description": "Valid description",
						"url": "https://example.com/job",
						"state": "APPLIED",
						"applied_at": "2025-12-06T18:28:44.0500+01:00"
					}
				`),
			},
			ucMock: mockCreateApplication{
				result: usecase.CreateApplicationOutput{
					Id:          uuid.MustParse("72325c96-bdea-4d60-a296-f84f2a5e3ef9"),
					Title:       "Lead Developer | Java",
					Description: "Valid description",
					Url:         "https://example.com/job",
					State:       domain.APPLIED,
					AppliedAt:   time.Date(2025, 12, 06, 19, 28, 44, 05, time.UTC),
					CreatedAt:   time.Time{},
					UpdatedAt:   time.Time{},
				},
				err: nil,
			},
			expectedBody:       `{"id":"72325c96-bdea-4d60-a296-f84f2a5e3ef9","title":"Lead Developer | Java","description":"Valid description","url":"https://example.com/job","state":"APPLIED","applied_at":"2025-12-06T19:28:44.000000005Z","created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z"}`,
			expectedStatusCode: http.StatusCreated,
		},
		{
			name: "Create application invalid url (unsupported scheme)",
			args: args{
				payload: []byte(`
					{
						"title": "Lead Developer | Java",
						"description": "Valid description",
						"url": "ftp://example.com/job",
						"state": "APPLIED",
						"applied_at": "2025-12-06T18:28:44.0500+01:00"
					}
				`),
			},
			ucMock:             mockCreateApplication{},
			expectedBody:       `{"errors":["Key: 'CreateApplicationInput.Url' Error:Field validation for 'Url' failed on the 'startswith' tag"]}`,
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "Create application invalid url (too long)",
			args: args{
				payload: []byte(`
					{
						"title": "Lead Developer | Java",
						"description": "Valid description",
						"url": "https://example.com/` + strings.Repeat("a", 500) + `",
						"state": "APPLIED",
						"applied_at": "2025-12-06T18:28:44.0500+01:00"
					}
				`),
			},
			ucMock:             mockCreateApplication{},
			expectedBody:       `{"errors":["URL must be a maximum of 512 characters in length"]}`,
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "Create application invalid state (empty)",
			args: args{
				payload: []byte(`
					{
						"title": "Lead Developer | Java",
						"description": "Valid description",
						"url": "https://example.com/job",
						"applied_at": "2025-12-06T18:28:44.0500+01:00"
					}
				`),
			},
			ucMock:             mockCreateApplication{},
			expectedBody:       `{"errors":["State is a required field"]}`,
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "Create application invalid state (random value)",
			args: args{
				payload: []byte(`
					{
						"title": "Lead Developer | Java",
						"description": "Valid description",
						"url": "https://example.com/job",
						"state": "RANDOM_STATE",
						"applied_at": "2025-12-06T18:28:44.0500+01:00"
					}
				`),
			},
			ucMock:             mockCreateApplication{},
			expectedBody:       `{"errors":["invalid application state"]}`,
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "Create application invalid application date (empty)",
			args: args{
				payload: []byte(`
					{
						"title": "Lead Developer | Java",
						"description": "Valid description",
						"url": "https://example.com/job",
						"state": "APPLIED"
					}
				`),
			},
			ucMock:             mockCreateApplication{},
			expectedBody:       `{"errors":["AppliedAt is a required field"]}`,
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "Create application invalid application date (future)",
			args: args{
				payload: []byte(`
					{
						"title": "Lead Developer | Java",
						"description": "Valid description",
						"url": "https://example.com/job",
						"state": "APPLIED",
						"applied_at": "2999-12-06T18:28:44.0500+01:00"
					}
				`),
			},
			ucMock:             mockCreateApplication{},
			expectedBody:       `{"errors":["AppliedAt must be less than the current Date \u0026 Time"]}`,
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "Create application invalid application date (wrong format)",
			args: args{
				payload: []byte(`
					{
						"title": "Lead Developer | Java",
						"description": "Valid description",
						"url": "https://example.com/job",
						"state": "APPLIED",
						"applied_at": "not-a-date"
					}
				`),
			},
			ucMock:             mockCreateApplication{},
			expectedBody:       `{"errors":["parsing time \"not-a-date\" as \"2006-01-02T15:04:05Z07:00\": cannot parse \"not-a-date\" as \"2006\""]}`,
			expectedStatusCode: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest(
				http.MethodPost,
				createApplicationEndpoint,
				bytes.NewReader(tt.args.payload),
			)

			var (
				w      = httptest.NewRecorder()
				action = NewCreateApplicationAction(tt.ucMock, validator, log.LoggerMock{})
			)

			action.Execute(w, req)

			if w.Code != tt.expectedStatusCode {
				t.Errorf(
					"[Test '%s'] The API endpoint returned an unexpected HTTP status code: got '%v', expected '%v'",
					tt.name,
					w.Code,
					tt.expectedStatusCode,
				)
			}

			var result = strings.TrimSpace(w.Body.String())
			if !strings.EqualFold(result, tt.expectedBody) {
				t.Errorf(
					"[Test '%s'] Result: '%v' | Expected: '%v'",
					tt.name,
					result,
					tt.expectedBody,
				)
			}
		})
	}
}
