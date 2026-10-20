package request

import (
	"fmt"
	"testing"

	"github.com/ghafbu/padegan-backend/pkg/defaulttestpkg/dto/defaulttestpkgdtousername"

	"github.com/ghafbu/padegan-backend/pkg/defaulttestpkg/dto/defaulttestpkgdtoemail"
	"github.com/ghafbu/padegan-backend/pkg/defaulttestpkg/dto/defaulttestpkgdtonationalcode"
	"github.com/ghafbu/padegan-backend/pkg/defaulttestpkg/dto/defaulttestpkgdtooneof"
	"github.com/ghafbu/padegan-backend/pkg/defaulttestpkg/dto/defaulttestpkgdtoordermethod"
	"github.com/ghafbu/padegan-backend/pkg/defaulttestpkg/dto/defaulttestpkgdtophone"

	"github.com/ghafbu/padegan-backend/lib/methods"
	"github.com/ghafbu/padegan-backend/pkg/defaulttestpkg/dto/defaulttestpkgdtomobile"
	"github.com/ghafbu/padegan-backend/pkg/ghafpkg"
	"github.com/stretchr/testify/assert"
)

var allowedFieldsOneOfFromDTOBatchGetOrderBy = []string{
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

func TestUsersDTOBatchGetUsernames(t *testing.T) {
	// Disable Log
	methods.LogActiveStatus(false)

	// Variables
	thisTestPrefixName := "adminmembers_dto_request_batch-get_usernames"

	// init server
	thisServer := fiberServerSetup(t)

	// Init Request Struct Data
	thisRequestData := UsersBatchGet{
		// Usernames:   &[]string{}, // TEST
		OrderBy:     "id",
		OrderMethod: "asc",
	}

	// Init Default Test Slice
	var thisDefaultTest = defaulttestpkgdtousername.NewDefaultTestDTOUsername()

	// Start t.Run > Loop...
	for _, CurrentContent := range thisDefaultTest {
		thisTestName := fmt.Sprintf("%s_%s_%s%d", thisTestPrefixName, CurrentContent.Category, CurrentContent.Provider, CurrentContent.Id)

		t.Run(thisTestName, func(t *testing.T) {
			// t.Parallel()

			// Log Test Name
			ghafpkg.GHPrintLnWithColor(ghafpkg.GHPrintColorPaletteWarning, thisTestName)

			// Insert Different Data
			*thisRequestData.Usernames = append(*thisRequestData.Usernames, CurrentContent.Content)

			// Send To Server
			// _ = thisResponseServerBody
			thisResponseServer, thisResponseServerBody, thisResponseController,
				err := methods.TestDTOSendRequestToFiberServer(
				t,
				thisServer,
				&UsersBatchGet{},
				thisRequestData,
				"/batch/get",
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
				assert.Equal(t, CurrentContent.Expected.Content, (*thisResponseController.Usernames)[0])
			}

			assert.Equal(t, CurrentContent.Expected.StatusCode, thisResponseServer.StatusCode)

			thisRequestData.Usernames = &[]string{}
		})
	}
}

func TestUsersDTOBatchGetMobiles(t *testing.T) {
	// Disable Log
	methods.LogActiveStatus(false)

	// Variables
	thisTestPrefixName := "adminmembers_dto_request_batch-get_mobiles"

	// init server
	thisServer := fiberServerSetup(t)

	// Init Request Struct Data
	thisRequestData := UsersBatchGet{
		Mobiles:     &[]string{},
		OrderBy:     "id",
		OrderMethod: "asc",
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
			*thisRequestData.Mobiles = append(*thisRequestData.Mobiles, CurrentContent.Content)

			// Send To Server
			// _ = thisResponseServerBody
			thisResponseServer, thisResponseServerBody, thisResponseController, err := methods.TestDTOSendRequestToFiberServer(
				t,
				thisServer,
				&UsersBatchGet{},
				thisRequestData,
				"/batch/get",
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
				assert.Equal(t, CurrentContent.Expected.Content, (*thisResponseController.Mobiles)[0])
			}
			assert.Equal(t, CurrentContent.Expected.StatusCode, thisResponseServer.StatusCode)

			thisRequestData.Mobiles = &[]string{}
		})
	}
}

func TestUsersDTOBatchGetPhones(t *testing.T) {
	// Disable Log
	methods.LogActiveStatus(false)

	// Variables
	thisTestPrefixName := "adminmembers_dto_request_batch-get_phones"

	// init server
	thisServer := fiberServerSetup(t)

	// Init Request Struct Data
	thisRequestData := UsersBatchGet{
		Phones:      &[]string{},
		OrderBy:     "id",
		OrderMethod: "asc",
	}

	// Init Default Test Slice
	var thisDefaultTest = defaulttestpkgdtophone.NewDefaultTestDTOPhone()

	// Start t.Run > Loop...
	for _, CurrentContent := range thisDefaultTest {
		thisTestName := fmt.Sprintf("%s_%s_%s%d", thisTestPrefixName, CurrentContent.Category, CurrentContent.Provider, CurrentContent.Id)

		t.Run(thisTestName, func(t *testing.T) {
			// t.Parallel()

			// Log Test Name
			ghafpkg.GHPrintLnWithColor(ghafpkg.GHPrintColorPaletteWarning, thisTestName)

			// Insert Different Data
			*thisRequestData.Phones = append(*thisRequestData.Phones, CurrentContent.Content)

			// Send To Server
			// _ = thisResponseServerBody
			thisResponseServer, thisResponseServerBody, thisResponseController, err := methods.TestDTOSendRequestToFiberServer(
				t,
				thisServer,
				&UsersBatchGet{},
				thisRequestData,
				"/batch/get",
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
				assert.Equal(t, CurrentContent.Expected.Content, (*thisResponseController.Phones)[0])
			}
			assert.Equal(t, CurrentContent.Expected.StatusCode, thisResponseServer.StatusCode)

			thisRequestData.Phones = &[]string{}
		})
	}
}

func TestUsersDTOBatchGetEmails(t *testing.T) {
	// Disable Log
	methods.LogActiveStatus(true)

	// Variables
	thisTestPrefixName := "adminmembers_dto_request_batch-get_emails"

	// init server
	thisServer := fiberServerSetup(t)

	// Init Request Struct Data
	thisRequestData := UsersBatchGet{
		Emails:      &[]string{},
		OrderBy:     "id",
		OrderMethod: "asc",
	}

	// Init Default Test Slice
	var thisDefaultTest = defaulttestpkgdtoemail.NewDefaultTestDTOEmail()

	// Start t.Run > Loop...
	for _, CurrentContent := range thisDefaultTest {
		thisTestName := fmt.Sprintf("%s_%s_%s%d", thisTestPrefixName, CurrentContent.Category, CurrentContent.Provider, CurrentContent.Id)

		t.Run(thisTestName, func(t *testing.T) {
			// t.Parallel()

			// Log Test Name
			ghafpkg.GHPrintLnWithColor(ghafpkg.GHPrintColorPaletteWarning, thisTestName)

			// Insert Different Data
			*thisRequestData.Emails = append(*thisRequestData.Emails, CurrentContent.Content)

			// Send To Server
			// _ = thisResponseServerBody
			thisResponseServer, thisResponseServerBody, thisResponseController, err := methods.TestDTOSendRequestToFiberServer(
				t,
				thisServer,
				&UsersBatchGet{},
				thisRequestData,
				"/batch/get",
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
				assert.Equal(t, CurrentContent.Expected.Content, (*thisResponseController.Emails)[0])
			}
			assert.Equal(t, CurrentContent.Expected.StatusCode, thisResponseServer.StatusCode)

			thisRequestData.Emails = &[]string{}
		})
	}
}

func TestUsersDTOBatchGetNationalCodes(t *testing.T) {
	// Disable Log
	methods.LogActiveStatus(false)

	// Variables
	thisTestPrefixName := "adminmembers_dto_request_batch-get_national_codes"

	// init server
	thisServer := fiberServerSetup(t)

	// Init Request Struct Data
	thisRequestData := UsersBatchGet{
		NationalCodes: &[]string{},
		OrderBy:       "id",
		OrderMethod:   "asc",
	}

	// Init Default Test Slice
	var thisDefaultTest = defaulttestpkgdtonationalcode.NewDefaultTestDTONationalCode()

	// Start t.Run > Loop...
	for _, CurrentContent := range thisDefaultTest {
		thisTestName := fmt.Sprintf("%s_%s_%s%d", thisTestPrefixName, CurrentContent.Category, CurrentContent.Provider, CurrentContent.Id)

		t.Run(thisTestName, func(t *testing.T) {
			// t.Parallel()

			// Log Test Name
			ghafpkg.GHPrintLnWithColor(ghafpkg.GHPrintColorPaletteWarning, thisTestName)

			// Insert Different Data
			*thisRequestData.NationalCodes = append(*thisRequestData.NationalCodes, CurrentContent.Content)

			// Send To Server
			// _ = thisResponseServerBody
			thisResponseServer, thisResponseServerBody, thisResponseController, err := methods.TestDTOSendRequestToFiberServer(
				t,
				thisServer,
				&UsersBatchGet{},
				thisRequestData,
				"/batch/get",
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
				assert.Equal(t, CurrentContent.Expected.Content, (*thisResponseController.NationalCodes)[0])
			}
			assert.Equal(t, CurrentContent.Expected.StatusCode, thisResponseServer.StatusCode)

			thisRequestData.NationalCodes = &[]string{}
		})
	}
}

func TestUsersDTOBatchGetOrderBy(t *testing.T) {
	// Disable Log
	methods.LogActiveStatus(false)

	// Variables
	thisTestPrefixName := "adminmembers_dto_request_batch-get_order-by"

	// init server
	thisServer := fiberServerSetup(t)

	// Init Request Struct Data
	thisRequestData := UsersBatchGet{
		OrderMethod: "asc",
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
			thisRequestData.OrderBy = CurrentContent.Content

			// Send To Server
			// _ = thisResponseServerBody
			thisResponseServer, thisResponseServerBody, thisResponseController, err := methods.TestDTOSendRequestToFiberServer(
				t,
				thisServer,
				&UsersBatchGet{},
				thisRequestData,
				"/batch/get",
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
				random, _ := getRandomItem(allowedFieldsOneOfFromDTOBatchGetOrderBy)
				assert.Equal(t, random, thisResponseController.OrderBy)
			}
			assert.Equal(t, CurrentContent.Expected.StatusCode, thisResponseServer.StatusCode)
		})
	}
}

func TestUsersDTOBatchGetOrderMethod(t *testing.T) {
	// Disable Log
	methods.LogActiveStatus(false)

	// Variables
	thisTestPrefixName := "adminmembers_dto_request_batch-get_order-method"

	// init server
	thisServer := fiberServerSetup(t)

	// Init Request Struct Data
	thisRequestData := UsersBatchGet{
		OrderBy: "id",
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
			thisRequestData.OrderMethod = CurrentContent.Content

			// Send To Server
			// _ = thisResponseServerBody
			thisResponseServer, thisResponseServerBody, thisResponseController, err := methods.TestDTOSendRequestToFiberServer(
				t,
				thisServer,
				&UsersBatchGet{},
				thisRequestData,
				"/batch/get",
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
				assert.Equal(t, CurrentContent.Expected.Content, thisResponseController.OrderMethod)
			}
			assert.Equal(t, CurrentContent.Expected.StatusCode, thisResponseServer.StatusCode)
		})
	}
}
