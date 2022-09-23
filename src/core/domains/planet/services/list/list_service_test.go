package list

import (
	"api-sw/mocks"
	"api-sw/src/core/domains/planet/entities"
	"api-sw/src/shared/tools/communication"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestSevice(t *testing.T) {
	comm := communication.New()

	useCases := map[string]struct {
		expectedResponse communication.Response
		prepare          func(repositoryMock *mocks.MockIPlanetRepository, loggerMock *mocks.MockILoggerProvider)
	}{
		"sucess": {
			expectedResponse: communication.Response{
				Status:  200,
				Code:    comm.Mapping["success"].Code,
				Message: comm.Mapping["success"].Message,
			},
			prepare: func(repositoryMock *mocks.MockIPlanetRepository, loggerMock *mocks.MockILoggerProvider) {
				expectedData := []entities.Planet{
					{
						ID:      "1",
						Name:    "Planet 1",
						Climate: "Climante",
						Terrain: "Terrain",
					},
					{
						ID:      "2",
						Name:    "Planet 2",
						Climate: "Climante",
						Terrain: "Terrain",
					},
				}
				repositoryMock.EXPECT().FindAll(gomock.Any()).Return(expectedData, nil)

				loggerMock.EXPECT().Info(gomock.Any(), gomock.Any())
				loggerMock.EXPECT().Error(gomock.Any(), gomock.Any()).AnyTimes()
			},
		},
		"success: should return status 200 if not return documents": {
			expectedResponse: communication.Response{
				Status:  200,
				Code:    comm.Mapping["success"].Code,
				Message: comm.Mapping["success"].Message,
			},
			prepare: func(repositoryMock *mocks.MockIPlanetRepository, loggerMock *mocks.MockILoggerProvider) {
				expectedData := []entities.Planet{}

				repositoryMock.EXPECT().FindAll(gomock.Any()).Return(expectedData, nil)
				loggerMock.EXPECT().Info(gomock.Any(), gomock.Any())
				loggerMock.EXPECT().Error(gomock.Any(), gomock.Any()).AnyTimes()
			},
		},
		"error: should return status 400 if error to find document in database": {
			expectedResponse: communication.Response{
				Status:  400,
				Code:    comm.Mapping["error_list"].Code,
				Message: comm.Mapping["error_list"].Message,
			},
			prepare: func(repositoryMock *mocks.MockIPlanetRepository, loggerMock *mocks.MockILoggerProvider) {
				expectedData := []entities.Planet{}

				repositoryMock.EXPECT().FindAll(gomock.Any()).Return(expectedData, errors.New("error"))

				loggerMock.EXPECT().Info(gomock.Any(), gomock.Any())
				loggerMock.EXPECT().Error(gomock.Any(), gomock.Any()).AnyTimes()
			},
		},
	}

	for name, useCase := range useCases {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repository := mocks.NewMockIPlanetRepository(ctrl)
			logger := mocks.NewMockILoggerProvider(ctrl)

			useCase.prepare(repository, logger)

			service := Service{
				Logger:     logger,
				Repository: repository,
			}

			response := service.Execute()

			if response.Status != useCase.expectedResponse.Status {
				t.Errorf("expected %d, but got %d", useCase.expectedResponse.Status, response.Status)
			}

			if response.Message != useCase.expectedResponse.Message {
				t.Errorf("expected %s, but got %s", useCase.expectedResponse.Message, response.Message)
			}

		})
	}

}
