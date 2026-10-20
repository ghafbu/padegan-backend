package methods

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/ghafbu/padegan-backend/pkg/ghafpkg"
	"github.com/gofiber/fiber/v3"
)

func TestDTOSendRequestToFiberServer[T comparable](
	t *testing.T,
	paramFiberServer *fiber.App,
	paramDTOStruct T,
	thisRequestData interface{},
	paramRoute string,
	paramMethod string,
	paramQueryParams *map[string]interface{},
	paramHeaders *map[string]string,
) (*http.Response, string, T, error) {
	ghafpkg.GHPrintLn("TestDTOSendRequestToFiberServer: start")

	// Convert Request Data To JSON
	thisRequestJsonData, err := json.Marshal(thisRequestData)
	if err != nil {
		t.Fatalf("Convert Request Struct To JSON says: %v", err)
	}

	/*
	 *
	 *
	 * Add Query Parameters To Route
	 */
	if paramQueryParams != nil && len(*paramQueryParams) > 0 {
		params := url.Values{}
		for key, value := range *paramQueryParams {
			// تبدیل مقدار به string
			strValue := fmt.Sprintf("%v", value)
			params.Add(key, strValue)
		}
		paramRoute = paramRoute + "?" + params.Encode()
	}

	/*
	*
	*
	* Send To Server
	 */
	// Create HTTP Client Request
	thisRequest := httptest.NewRequest(paramMethod, paramRoute, bytes.NewReader(thisRequestJsonData))

	// Customize HTTP Client Request Header
	thisRequest.Header.Set("Content-Type", "application/json")
	// افزودن هدرهای سفارشی (اگر وجود داشته باشد)
	if paramHeaders != nil {
		for key, value := range *paramHeaders {
			thisRequest.Header.Set(key, value)
		}
	}

	// Send HTTP Client Request To Server
	thisResponseServer, err := paramFiberServer.Test(thisRequest)
	if err != nil {
		t.Fatalf("Send Request To Server Fail - Errors: %v", err)
	}

	/*
	*
	*
	* Analytic Response
	 */
	// Read the response body
	ghafpkg.GHPrintLn("TestDTOSendRequestToFiberServer: Read the response body")
	bodyBytes, err := io.ReadAll(thisResponseServer.Body)
	if err != nil {
		t.Fatalf("Error reading response body: %v", err)
	}
	bodyBytesString := string(bodyBytes)
	// Close the response body (important!)
	defer thisResponseServer.Body.Close()
	// fmt.Printf("bodyBytes:\n %v \n", bodyBytes)

	// Convert Body To Struct
	if thisResponseServer.StatusCode == http.StatusOK {
		err = json.Unmarshal(bodyBytes, paramDTOStruct)
		if err != nil {
			t.Fatalf("Error unmarshaling response body: %v", err)
		}
	}

	return thisResponseServer, bodyBytesString, paramDTOStruct, nil
}
