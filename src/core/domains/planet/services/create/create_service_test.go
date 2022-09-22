package create

import (
	"api-sw/mocks"
	"api-sw/src/core/domains/planet/entities"
	"api-sw/src/shared/tools/communication"
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestService(t *testing.T) {
	comm := communication.New()
	useCases := map[string]struct {
		inputData      Dto
		expectResponse communication.Response
		prepare        func(repositoryMock *mocks.MockIPlanetRepository, loggerMock *mocks.MockILoggerProvider)
	}{
		"sucess": {
			inputData: Dto{
				Name:    "Planet 1",
				Climate: "Climate",
				Terrain: "Terrain",
			},
			expectResponse: communication.Response{
				Status:  201,
				Code:    comm.Mapping["success_create"].Code,
				Message: comm.Mapping["success_create"].Message,
			},
			prepare: func(repositoryMock *mocks.MockIPlanetRepository, loggerMock *mocks.MockILoggerProvider) {
				expectedData := entities.Planet{}
				expectedCreated := entities.Planet{
					ID:      "1",
					Name:    "Planet 1",
					Climate: "Climante",
					Terrain: "Terrain",
				}

				repositoryMock.EXPECT().FindByName(gomock.Any()).Return(expectedData, nil)
				repositoryMock.EXPECT().Create(gomock.Any()).Return(expectedCreated, nil)
				loggerMock.EXPECT().Info(gomock.Any(), gomock.Any())
				loggerMock.EXPECT().Error(gomock.Any(), gomock.Any()).AnyTimes()
			},
		},
		"error : should return status code 400 when name is empty": {
			inputData: Dto{
				Name:    "",
				Climate: "Climante",
				Terrain: "Terrain",
			},
			expectResponse: communication.Response{
				Status:  400,
				Code:    comm.Mapping["validate_failed"].Code,
				Message: comm.Mapping["validate_failed"].Message,
			},
			prepare: func(repositoryMock *mocks.MockIPlanetRepository, loggerMock *mocks.MockILoggerProvider) {
				loggerMock.EXPECT().Info(gomock.Any(), gomock.Any())
				loggerMock.EXPECT().Error(gomock.Any(), gomock.Any()).AnyTimes()
			},
		},
		"error : should return status code 400 when Climate is empty": {
			inputData: Dto{
				Name:    "Planet 1",
				Climate: "",
				Terrain: "Terrain",
			},
			expectResponse: communication.Response{
				Status:  400,
				Code:    comm.Mapping["validate_failed"].Code,
				Message: comm.Mapping["validate_failed"].Message,
			},
			prepare: func(repositoryMock *mocks.MockIPlanetRepository, loggerMock *mocks.MockILoggerProvider) {
				loggerMock.EXPECT().Info(gomock.Any(), gomock.Any())
				loggerMock.EXPECT().Error(gomock.Any(), gomock.Any()).AnyTimes()
			},
		},
		"error : should return status code 400 when Terrain is empty": {
			inputData: Dto{
				Name:    "Planet 1",
				Climate: "Climate",
				Terrain: "",
			},
			expectResponse: communication.Response{
				Status:  400,
				Code:    comm.Mapping["validate_failed"].Code,
				Message: comm.Mapping["validate_failed"].Message,
			},
			prepare: func(repositoryMock *mocks.MockIPlanetRepository, loggerMock *mocks.MockILoggerProvider) {
				loggerMock.EXPECT().Info(gomock.Any(), gomock.Any())
				loggerMock.EXPECT().Error(gomock.Any(), gomock.Any()).AnyTimes()
			},
		},

		"error : should return status code 400 when error to query name in database": {
			inputData: Dto{
				Name:    "planet 1",
				Climate: "Climante",
				Terrain: "Terrain",
			},
			expectResponse: communication.Response{
				Status:  400,
				Code:    comm.Mapping["error_create"].Code,
				Message: comm.Mapping["error_create"].Message,
			},
			prepare: func(repositoryMock *mocks.MockIPlanetRepository, loggerMock *mocks.MockILoggerProvider) {
				exepctedData := entities.Planet{}

				repositoryMock.EXPECT().FindByName(gomock.Any()).Return(exepctedData, errors.New("erro to query name"))

				loggerMock.EXPECT().Info(gomock.Any(), gomock.Any())
				loggerMock.EXPECT().Error(gomock.Any(), gomock.Any()).AnyTimes()
			},
		},
		"error : should return status code 400 when name already exist on datebase": {
			inputData: Dto{
				Name:    "planet 1",
				Climate: "Climante",
				Terrain: "Terrain",
			},
			expectResponse: communication.Response{
				Status:  400,
				Code:    comm.Mapping["already_exists"].Code,
				Message: comm.Mapping["alread_exits"].Message,
			},
			prepare: func(repositoryMock *mocks.MockIPlanetRepository, loggerMock *mocks.MockILoggerProvider) {
				expectedData := entities.Planet{
					ID:      "1",
					Name:    "Planet 1",
					Climate: "Climante",
					Terrain: "Terrain",
				}
				repositoryMock.EXPECT().FindByName(gomock.Any()).Return(expectedData, nil)

				loggerMock.EXPECT().Info(gomock.Any(), gomock.Any())
				loggerMock.EXPECT().Error(gomock.Any(), gomock.Any()).AnyTimes()
			},
		},
		"error : should return status code 400 when error to create document in database": {
			inputData: Dto{
				Name:    "Planet 1",
				Climate: "Climante",
				Terrain: "Terrain",
			},
			expectResponse: communication.Response{
				Status:  400,
				Code:    comm.Mapping["error_create"].Code,
				Message: comm.Mapping["error_create"].Message,
			},
			prepare: func(repositoryMock *mocks.MockIPlanetRepository, loggerMock *mocks.MockILoggerProvider) {
				expectedData := entities.Planet{}
				expectedCreated := entities.Planet{}

				repositoryMock.EXPECT().FindByName(gomock.Any()).Return(expectedData, nil)
				repositoryMock.EXPECT().Create(gomock.Any()).Return(expectedCreated, errors.New("error to create"))

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

			response := service.Execute(useCase.inputData)

			if response.Status != useCase.expectResponse.Status {
				t.Errorf("expected %d, but got %d", useCase.expectResponse.Status, response.Status)
			}
			if response.Message != useCase.expectResponse.Message {
				t.Errorf("expected %s, but got %s", useCase.expectResponse.Message, response.Message)

			}
		})
	}
}
