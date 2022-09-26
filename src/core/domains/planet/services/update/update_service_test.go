package update

import (
	"api-sw/mocks"
	"api-sw/src/core/domains/planet/entities"
	"api-sw/src/shared/tools/communication"
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestPlanetUpdateService(t *testing.T) {
	comm := communication.New()

	useCases := map[string]struct {
		inputData        Dto
		documentID       string
		expectedResponse communication.Response
		prepare          func(repositoryMock *mocks.MockIPlanetRepository, loggerMock *mocks.MockILoggerProvider)
	}{
		"success": {
			documentID: "1",
			inputData: Dto{
				Name:    "Planet 1",
				Climate: "climate",
				Terrain: "terrain",
			},
			expectedResponse: communication.Response{
				Status:  200,
				Code:    comm.Mapping["success_update"].Code,
				Message: comm.Mapping["success_update"].Message,
			},
			prepare: func(repositoryMock *mocks.MockIPlanetRepository, loggerMock *mocks.MockILoggerProvider) {
				expectedData := entities.Planet{}
				expectedUpdate := entities.Planet{
					ID:      "1",
					Name:    "Planet 1",
					Climate: "climate",
					Terrain: "terrain",
				}
				repositoryMock.EXPECT().FindByName(gomock.Any()).Return(expectedData, nil)
				repositoryMock.EXPECT().Update(gomock.Any(), gomock.Any()).Return(expectedUpdate, nil)

				loggerMock.EXPECT().Info(gomock.Any(), gomock.Any())
				loggerMock.EXPECT().Error(gomock.Any(), gomock.Any()).AnyTimes()
			},
		},
		"error: should return status code 400 when name is empty": {
			documentID: "1",
			inputData: Dto{
				Name:    "",
				Climate: "climate",
				Terrain: "terrain",
			},
			expectedResponse: communication.Response{
				Status:  400,
				Code:    comm.Mapping["validate_failed"].Code,
				Message: comm.Mapping["validate_failed"].Message,
			},
			prepare: func(repositoryMock *mocks.MockIPlanetRepository, loggerMock *mocks.MockILoggerProvider) {
				loggerMock.EXPECT().Info(gomock.Any(), gomock.Any())
				loggerMock.EXPECT().Error(gomock.Any(), gomock.Any()).AnyTimes()
			},
		},
		"error: should return status code 400 when climate is empty": {
			documentID: "1",
			inputData: Dto{
				Name:    "Planet 1",
				Climate: "",
				Terrain: "terrain",
			},
			expectedResponse: communication.Response{
				Status:  400,
				Code:    comm.Mapping["validate_failed"].Code,
				Message: comm.Mapping["validate_failed"].Message,
			},
			prepare: func(repositoryMock *mocks.MockIPlanetRepository, loggerMock *mocks.MockILoggerProvider) {
				loggerMock.EXPECT().Info(gomock.Any(), gomock.Any())
				loggerMock.EXPECT().Error(gomock.Any(), gomock.Any()).AnyTimes()
			},
		},
		"error: should return status code 400 when terrain is empty": {
			documentID: "1",
			inputData: Dto{
				Name:    "Planet 1",
				Climate: "climate",
				Terrain: "",
			},
			expectedResponse: communication.Response{
				Status:  400,
				Code:    comm.Mapping["validate_failed"].Code,
				Message: comm.Mapping["validate_failed"].Message,
			},
			prepare: func(repositoryMock *mocks.MockIPlanetRepository, loggerMock *mocks.MockILoggerProvider) {
				loggerMock.EXPECT().Info(gomock.Any(), gomock.Any())
				loggerMock.EXPECT().Error(gomock.Any(), gomock.Any()).AnyTimes()
			},
		},
		"error: should return status code 400 when error to query name in database": {
			documentID: "1",
			inputData: Dto{
				Name:    "Planet 1",
				Climate: "climate",
				Terrain: "terrain",
			},
			expectedResponse: communication.Response{
				Status:  400,
				Code:    comm.Mapping["error_update"].Code,
				Message: comm.Mapping["error_update"].Message,
			},
			prepare: func(repositoryMock *mocks.MockIPlanetRepository, loggerMock *mocks.MockILoggerProvider) {
				expectedData := entities.Planet{}

				repositoryMock.EXPECT().FindByName(gomock.Any()).Return(expectedData, errors.New("error to query name"))

				loggerMock.EXPECT().Info(gomock.Any(), gomock.Any())
				loggerMock.EXPECT().Error(gomock.Any(), gomock.Any()).AnyTimes()
			},
		},
		"error: should return status code 400 when name already exist on database": {
			documentID: "1",
			inputData: Dto{
				Name:    "Planet 1",
				Climate: "climate",
				Terrain: "terrain",
			},
			expectedResponse: communication.Response{
				Status:  400,
				Code:    comm.Mapping["already_exist"].Code,
				Message: comm.Mapping["already_exist"].Message,
			},
			prepare: func(repositoryMock *mocks.MockIPlanetRepository, loggerMock *mocks.MockILoggerProvider) {
				expectedData := entities.Planet{
					ID:      "2",
					Name:    "Planet 1",
					Climate: "climate",
					Terrain: "terrain",
				}
				repositoryMock.EXPECT().FindByName(gomock.Any()).Return(expectedData, nil)

				loggerMock.EXPECT().Info(gomock.Any(), gomock.Any())
				loggerMock.EXPECT().Error(gomock.Any(), gomock.Any()).AnyTimes()
			},
		},
		"should return status code 400 when error to update document in database": {
			documentID: "1",
			inputData: Dto{
				Name:    "Planet 1",
				Climate: "climate",
				Terrain: "terrain",
			},
			expectedResponse: communication.Response{
				Status:  400,
				Code:    comm.Mapping["error_update"].Code,
				Message: comm.Mapping["error_update"].Message,
			},
			prepare: func(repositoryMock *mocks.MockIPlanetRepository, loggerMock *mocks.MockILoggerProvider) {
				expectedData := entities.Planet{}
				expectedDocument := entities.Planet{}

				repositoryMock.EXPECT().FindByName(gomock.Any()).Return(expectedData, nil)
				repositoryMock.EXPECT().Update(gomock.Any(), gomock.Any()).Return(expectedDocument, errors.New("error to update"))

				loggerMock.EXPECT().Info(gomock.Any(), gomock.Any())
				loggerMock.EXPECT().Error(gomock.Any(), gomock.Any()).AnyTimes()

			},
		},
	}

	for name, useCase := range useCases {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			ctx := context.Background()
			logger := mocks.NewMockILoggerProvider(ctrl)
			repository := mocks.NewMockIPlanetRepository(ctrl)
			useCase.prepare(repository, logger)

			service := Service{
				Context:    ctx,
				Logger:     logger,
				Repository: repository,
			}

			response := service.Execute(useCase.documentID, useCase.inputData)

			if response.Status != useCase.expectedResponse.Status {
				t.Errorf("expected %d, but got %d", useCase.expectedResponse.Status, response.Status)
			}

			if response.Message != useCase.expectedResponse.Message {
				t.Errorf("expected %s, but got %s", useCase.expectedResponse.Message, response.Message)
			}
		})
	}
}
