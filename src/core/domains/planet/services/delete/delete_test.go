package delete

import (
	"api-sw/mocks"
	"api-sw/src/shared/tools/communication"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestService(t *testing.T) {
	comm := communication.New()

	useCases := map[string]struct {
		expectedResponse communication.Response
		documentID       string
		prepare          func(repositoryMock *mocks.MockIPlanetRepository, loggerMock *mocks.MockILoggerProvider)
	}{
		"success": {
			expectedResponse: communication.Response{
				Status:  200,
				Code:    comm.Mapping["success"].Code,
				Message: comm.Mapping["success"].Message,
			},
			prepare: func(repositoryMock *mocks.MockIPlanetRepository, loggerMock *mocks.MockILoggerProvider) {

				repositoryMock.EXPECT().Delete(gomock.Any()).Return(nil)

				loggerMock.EXPECT().Info(gomock.Any(), gomock.Any())
				loggerMock.EXPECT().Error(gomock.Any(), gomock.Any()).AnyTimes()
			},
		},
		"error : should return status 400 if error to delete document in database": {
			expectedResponse: communication.Response{
				Status:  400,
				Code:    comm.Mapping["error_delete"].Code,
				Message: comm.Mapping["error_delete"].Message,
			},
			prepare: func(repositoryMock *mocks.MockIPlanetRepository, loggerMock *mocks.MockILoggerProvider) {

				repositoryMock.EXPECT().Delete(gomock.Any()).Return(errors.New("error delete"))

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

			Service := Service{
				Logger:     logger,
				Repository: repository,
			}

			response := Service.Execute(useCase.documentID)

			if response.Status != useCase.expectedResponse.Status {
				t.Errorf("Expected %d, but got %d", useCase.expectedResponse.Status, response.Status)
			}

			if response.Message != useCase.expectedResponse.Message {
				t.Errorf("Expected %s, but got %s", useCase.expectedResponse.Message, response.Message)
			}

		})
	}
}
