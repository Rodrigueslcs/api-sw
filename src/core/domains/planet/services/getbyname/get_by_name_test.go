package getbyname

import (
	"api-sw/mocks"
	"api-sw/src/core/domains/planet/entities"
	"api-sw/src/shared/tools/communication"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestPlanetGetByName(t *testing.T) {
	comm := communication.New()
	useCases := map[string]struct {
		documentName     string
		expectedResponse communication.Response
		prepare          func(repositoryMock *mocks.MockIPlanetRepository, loggerMock *mocks.MockILoggerProvider)
	}{
		"success": {
			documentName: " Planet 1",
			expectedResponse: communication.Response{
				Status:  200,
				Code:    comm.Mapping["success_search"].Code,
				Message: comm.Mapping["success_search"].Message,
			},
			prepare: func(repositoryMock *mocks.MockIPlanetRepository, loggerMock *mocks.MockILoggerProvider) {
				expectedSearch := entities.Planet{
					ID:          "1",
					Name:        "planet 1",
					Climate:     "climate",
					Terrain:     "terrian",
					Apparitions: 0,
				}

				repositoryMock.EXPECT().FindByName(gomock.Any()).Return(expectedSearch, nil)

				loggerMock.EXPECT().Info(gomock.Any(), gomock.Any())
				loggerMock.EXPECT().Error(gomock.Any(), gomock.Any()).AnyTimes()
			},
		},
		"success: should return status 200 if not return documents": {
			expectedResponse: communication.Response{
				Status:  200,
				Code:    comm.Mapping["success_search"].Code,
				Message: comm.Mapping["success_search"].Message,
			},
			prepare: func(repositoryMock *mocks.MockIPlanetRepository, loggerMock *mocks.MockILoggerProvider) {
				expectedData := entities.Planet{}

				repositoryMock.EXPECT().FindByName(gomock.Any()).Return(expectedData, nil)

				loggerMock.EXPECT().Info(gomock.Any(), gomock.Any())
				loggerMock.EXPECT().Error(gomock.Any(), gomock.Any()).AnyTimes()
			},
		},
		"error: should return status code 400 when error to query in database.": {
			documentName: "Planet 1",
			expectedResponse: communication.Response{
				Status:  400,
				Code:    comm.Mapping["planet_not_found"].Code,
				Message: comm.Mapping["planet_not_found"].Message,
			},
			prepare: func(repositoryMock *mocks.MockIPlanetRepository, loggerMock *mocks.MockILoggerProvider) {
				expectedData := entities.Planet{}

				repositoryMock.EXPECT().FindByName(gomock.Any()).Return(expectedData, errors.New("error to query name"))

				loggerMock.EXPECT().Info(gomock.Any(), gomock.Any())
				loggerMock.EXPECT().Error(gomock.Any(), gomock.Any()).AnyTimes()
			},
		},
	}
	for name, useCase := range useCases {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			// ctx := context.Background()
			logger := mocks.NewMockILoggerProvider(ctrl)
			repository := mocks.NewMockIPlanetRepository(ctrl)

			useCase.prepare(repository, logger)

			service := Service{
				Repository: repository,
				Logger:     logger,
			}
			response := service.Execute(useCase.documentName)

			if response.Status != useCase.expectedResponse.Status {
				t.Errorf("Expected %d, but got %d", useCase.expectedResponse.Status, response.Status)
			}

			if response.Message != useCase.expectedResponse.Message {
				t.Errorf("Expected %s, but got %s", useCase.expectedResponse.Message, response.Message)
			}
		})
	}
}
