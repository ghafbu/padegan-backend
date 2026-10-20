package request

import (
	"fmt"
	"testing"

	"github.com/ghafbu/padegan-backend/lib/methods"
	"github.com/ghafbu/padegan-backend/pkg/defaulttestpkg/dto/defaulttestpkgdtocode"
	"github.com/ghafbu/padegan-backend/pkg/defaulttestpkg/dto/defaulttestpkgdtomobile"
	"github.com/ghafbu/padegan-backend/pkg/ghafpkg"
	"github.com/stretchr/testify/assert"
)

func TestUsersDTOResetPasswordVerifyMobile(t *testing.T) {
	// Disable Log
	methods.LogActiveStatus(false)

	// Variables
	thisTestPrefixName := "adminmembers_dto_request_reset-password-verify_mobile"

	// init server
	thisServer := fiberServerSetup(t)

	// Init Request Struct Data
	thisRequestData := UsersResetPasswordVerify{
		// Mobile:      "", TEST
		Code: "12345",
	}

	// Init Default Test Slice
	var thisDefaultTest = defaulttestpkgdtomobile.NewDefaultTestDTOMobile()

	// Start t.Run > Loop...
	for _, CurrentContent := range thisDefaultTest {
		thisTestName := fmt.Sprintf("%s_%s_%s%d", thisTestPrefixName, CurrentContent.Category, CurrentContent.Provider, CurrentContent.Id)

		t.Run(thisTestName, func(t *testing.T) {
			// t.Parallel()

			// Log Test Name
			ghafpkg.GHPrintLnWithColor(ghafpkg.GHPrintColorPaletteWarning, thisTestName)

			// Insert Different Data
			thisRequestData.Mobile = CurrentContent.Content

			// Send To Server
			// _ = thisResponseServerBody
			thisResponseServer, thisResponseServerBody, thisResponseController, err := methods.TestDTOSendRequestToFiberServer(
				t,
				thisServer,
				&UsersResetPasswordVerify{},
				thisRequestData,
				"/reset-password/verify",
				"POST",
				nil,
				nil,
			)
			if err != nil {
				t.Fatalf("[%s] DTO Send Request To Fiber Server FAILD: %v", thisTestPrefixName, err)
			}

			// Print Details
			ghafpkg.GHPrintJsonLnWithColor(ghafpkg.GHPrintColorPaletteDebug, "thisResponseServerBody:", thisResponseServerBody)
			ghafpkg.GHPrintLnWithColor(ghafpkg.GHPrintColorPaletteDebug, "scenario:", CurrentContent.Scenario)

			// Expected
			if CurrentContent.Expected.Content != "" {
				assert.Equal(t, CurrentContent.Expected.Content, thisResponseController.Mobile)
			}
			assert.Equal(t, CurrentContent.Expected.StatusCode, thisResponseServer.StatusCode)
		})
	}

}

func TestUsersDTOResetPasswordVerifyCode(t *testing.T) {
	// Disable Log
	methods.LogActiveStatus(false)

	// Variables
	thisTestPrefixName := "adminmembers_dto_request_reset-password-verify_code"

	// init server
	thisServer := fiberServerSetup(t)

	// Init Request Struct Data
	thisRequestData := UsersResetPasswordVerify{
		Mobile: "+989112253390",
		// Code: "12345", TEST
	}

	// Init Default Test Slice
	var thisDefaultTest = defaulttestpkgdtocode.NewDefaultTestDTOCode()

	// Start t.Run > Loop...
	for _, CurrentContent := range thisDefaultTest {
		thisTestName := fmt.Sprintf("%s_%s_%s%d", thisTestPrefixName, CurrentContent.Category, CurrentContent.Provider, CurrentContent.Id)

		t.Run(thisTestName, func(t *testing.T) {
			// t.Parallel()

			// Log Test Name
			ghafpkg.GHPrintLnWithColor(ghafpkg.GHPrintColorPaletteWarning, thisTestName)

			// Insert Different Data
			thisRequestData.Code = CurrentContent.Content

			// Send To Server
			thisResponseServer, thisResponseServerBody, thisResponseController, err := methods.TestDTOSendRequestToFiberServer(
				t,
				thisServer,
				&UsersResetPasswordVerify{},
				thisRequestData,
				"/reset-password/verify",
				"POST",
				nil,
				nil,
			)
			if err != nil {
				t.Fatalf("[%s] DTO Send Request To Fiber Server FAILD: %v", thisTestPrefixName, err)
			}

			// Print Details
			ghafpkg.GHPrintJsonLnWithColor(ghafpkg.GHPrintColorPaletteDebug, "thisResponseServerBody:", thisResponseServerBody)
			ghafpkg.GHPrintLnWithColor(ghafpkg.GHPrintColorPaletteDebug, "scenario:", CurrentContent.Scenario)

			// Expected
			if CurrentContent.Expected.Content != "" {
				assert.Equal(t, CurrentContent.Expected.Content, thisResponseController.Code)
			}
			assert.Equal(t, CurrentContent.Expected.StatusCode, thisResponseServer.StatusCode)
		})
	}

}
