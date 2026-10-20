package request

import (
	"fmt"
	"testing"

	"github.com/ghafbu/padegan-backend/pkg/defaulttestpkg/dto/defaulttestpkgdtooneof"
	"github.com/ghafbu/padegan-backend/pkg/defaulttestpkg/dto/defaulttestpkgdtoordermethod"

	"github.com/ghafbu/padegan-backend/lib/methods"
	"github.com/ghafbu/padegan-backend/pkg/ghafpkg"
	"github.com/stretchr/testify/assert"
)

var allowedFieldsOneOfFromDTOBatchUpdateOrderBy = []string{
	"id",
	"family",
	"mobile",
	"mobile_country_code",
	"phone",
	"phone_country_code",
	"email",
	"birth_date",
	"national_code",
	"username",
	"org_theme_mode",
	"language",
	"profile_status",
	"status",
}

func TestUsersDTOBatchUpdateOrderBy(t *testing.T) {
	// Disable Log
	methods.LogActiveStatus(false)

	// Variables
	thisTestPrefixName := "adminmembers_dto_request_batch-update_order-by"

	// init server
	thisServer := fiberServerSetup(t)

	// Init Request Struct Data
	thisRequestData := UsersBatchUpdate{
		Ids:    []uint64{1},
		Status: 1,
	}

	// Init Default Test Slice
	var thisDefaultTest = defaulttestpkgdtooneof.NewDefaultTestDTOOneOf()

	// Start t.Run > Loop...
	for _, CurrentContent := range thisDefaultTest {
		thisTestName := fmt.Sprintf("%s_%s_%s%d", thisTestPrefixName, CurrentContent.Category, CurrentContent.Provider, CurrentContent.Id)

		t.Run(thisTestName, func(t *testing.T) {
			// t.Parallel()

			// Log Test Name
			ghafpkg.GHPrintLnWithColor(ghafpkg.GHPrintColorPaletteWarning, thisTestName)

			// Insert Different Data
			thisRequestData.OrderBy = &CurrentContent.Content

			// Send To Server
			// _ = thisResponseServerBody
			thisResponseServer, thisResponseServerBody, thisResponseController, err := methods.TestDTOSendRequestToFiberServer(
				t,
				thisServer,
				&UsersBatchUpdate{},
				thisRequestData,
				"/batch/update",
				"PATCH",
				nil,
				nil,
			)
			if err != nil {
				t.Fatalf("[%s] DTO Send Request To Fiber Server FAILD: %v", thisTestPrefixName, err)
			}

			// Print Details
			ghafpkg.GHPrintJsonLnWithColor(ghafpkg.GHPrintColorPaletteDebug, "thisResponseServerBody:", thisResponseServerBody)
			ghafpkg.GHPrintLnWithColor(ghafpkg.GHPrintColorPaletteDebug, "scenario:", CurrentContent.Scenario)

			if CurrentContent.Expected.Content != "" {
				// Expected
				random, _ := getRandomItem(allowedFieldsOneOfFromDTOBatchUpdateOrderBy)
				assert.Equal(t, random, *thisResponseController.OrderBy)
			}
			assert.Equal(t, CurrentContent.Expected.StatusCode, thisResponseServer.StatusCode)
		})
	}
}

func TestUsersDTOBatchUpdateOrderMethod(t *testing.T) {
	// Disable Log
	methods.LogActiveStatus(false)

	// Variables
	thisTestPrefixName := "adminmembers_dto_request_batch-update_order-method"

	// init server
	thisServer := fiberServerSetup(t)

	// Init Request Struct Data
	thisRequestData := UsersBatchUpdate{
		Ids:    []uint64{1},
		Status: 1,
	}

	// Init Default Test Slice
	var thisDefaultTest = defaulttestpkgdtoordermethod.NewDefaultTestDTOOrderMethod()

	// Start t.Run > Loop...
	for _, CurrentContent := range thisDefaultTest {
		thisTestName := fmt.Sprintf("%s_%s_%s%d", thisTestPrefixName, CurrentContent.Category, CurrentContent.Provider, CurrentContent.Id)

		t.Run(thisTestName, func(t *testing.T) {
			// t.Parallel()

			// Log Test Name
			ghafpkg.GHPrintLnWithColor(ghafpkg.GHPrintColorPaletteWarning, thisTestName)

			// Insert Different Data
			thisRequestData.OrderMethod = &CurrentContent.Content

			// Send To Server
			// _ = thisResponseServerBody
			thisResponseServer, thisResponseServerBody, thisResponseController, err := methods.TestDTOSendRequestToFiberServer(
				t,
				thisServer,
				&UsersBatchUpdate{},
				thisRequestData,
				"/batch/update",
				"PATCH",
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
				assert.Equal(t, CurrentContent.Expected.Content, *thisResponseController.OrderMethod)
			}
			assert.Equal(t, CurrentContent.Expected.StatusCode, thisResponseServer.StatusCode)
		})
	}
}
