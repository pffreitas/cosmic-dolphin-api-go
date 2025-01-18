/*
Cosmic Dolphin

Testing PipelinesAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package cosmicdolphinapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/pffreitas/cosmic-dolphin-api-go"
)

func Test_cosmicdolphinapi_PipelinesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test PipelinesAPIService PipelinesFindByRefId", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var refId int32

		resp, httpRes, err := apiClient.PipelinesAPI.PipelinesFindByRefId(context.Background(), refId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
