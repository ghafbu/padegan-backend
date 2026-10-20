package request

import (
	"fmt"
	"testing"

	"github.com/ghafbu/padegan-backend/pkg/defaulttestpkg/dto/defaulttestpkgdtousername"

	"github.com/ghafbu/padegan-backend/pkg/defaulttestpkg/dto/defaulttestpkgdtobirthdate"
	"github.com/ghafbu/padegan-backend/pkg/defaulttestpkg/dto/defaulttestpkgdtoemail"
	"github.com/ghafbu/padegan-backend/pkg/defaulttestpkg/dto/defaulttestpkgdtonationalcode"
	"github.com/ghafbu/padegan-backend/pkg/defaulttestpkg/dto/defaulttestpkgdtophone"
	"github.com/ghafbu/padegan-backend/pkg/defaulttestpkg/dto/defaulttestpkgdtourl"

	"github.com/ghafbu/padegan-backend/lib/methods"
	"github.com/ghafbu/padegan-backend/pkg/defaulttestpkg/dto/defaulttestpkgdtolikename"
	"github.com/ghafbu/padegan-backend/pkg/ghafpkg"
	"github.com/stretchr/testify/assert"
)

func TestUsersDTOUpdateNameAndFamily(t *testing.T) {
	// Disable Log
	methods.LogActiveStatus(false)

	// Variables
	thisTestPrefixName := "adminmembers_dto_request_update_name_and_family"

	// init server
	thisServer := fiberServerSetup(t)

	// Init Request Struct Data
	thisRequestData := UsersUpdate{}

	// Init Default Test Slice
	var thisDefaultTest = defaulttestpkgdtolikename.NewDefaultTestDTOLikeName()

	// Start t.Run > Loop...
	for _, CurrentContent := range thisDefaultTest {
		thisTestName := fmt.Sprintf("%s_%s_%s%d", thisTestPrefixName, CurrentContent.Category, CurrentContent.Provider, CurrentContent.Id)

		t.Run(thisTestName, func(t *testing.T) {
			// t.Parallel()

			// Log Test Name
			ghafpkg.GHPrintLnWithColor(ghafpkg.GHPrintColorPaletteWarning, thisTestName)

			// Insert Different Data
			thisRequestData.Name = CurrentContent.Content
			thisRequestData.Family = CurrentContent.Content

			// Send To Server
			// _ = thisResponseServerBody
			thisResponseServer, thisResponseServerBody, thisResponseController,
				err := methods.TestDTOSendRequestToFiberServer(
				t,
				thisServer,
				&UsersUpdate{},
				thisRequestData,
				"/1",
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
				assert.Equal(t, CurrentContent.Expected.Content, thisResponseController.Name)
				assert.Equal(t, CurrentContent.Expected.Content, thisResponseController.Family)
			}

			assert.Equal(t, CurrentContent.Expected.StatusCode, thisResponseServer.StatusCode)
		})
	}
}

func TestUsersDTOUpdateUsername(t *testing.T) {
	// Disable Log
	methods.LogActiveStatus(false)

	// Variables
	thisTestPrefixName := "adminmembers_dto_request_update_username"

	// init server
	thisServer := fiberServerSetup(t)

	// Init Request Struct Data
	thisRequestData := UsersUpdate{}

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
			thisRequestData.Username = &CurrentContent.Content

			// Send To Server
			// _ = thisResponseServerBody
			thisResponseServer, thisResponseServerBody, thisResponseController,
				err := methods.TestDTOSendRequestToFiberServer(
				t,
				thisServer,
				&UsersUpdate{},
				thisRequestData,
				"/1",
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
				assert.Equal(t, CurrentContent.Expected.Content, *thisResponseController.Username)
			}

			assert.Equal(t, CurrentContent.Expected.StatusCode, thisResponseServer.StatusCode)
		})
	}
}

//
// func TestUsersDTOUpdateMobile(t *testing.T) {
// 	// Disable Log
// 	methods.LogActiveStatus(false)
//
// 	// Variables
// 	thisTestPrefixName := "adminmembers_dto_request_update_mobile"
//
// 	// init server
// 	thisServer := fiberServerSetup(t)
//
// 	// Init Request Struct Data
// 	thisRequestData := UsersUpdate{
// 		Name:   "ABCDEFG",
// 		Family: "GFEDCBA",
// 	}
//
// 	// Init Default Test Slice
// 	var thisDefaultTest = defaulttestpkgdtomobile.NewDefaultTestDTOMobile()
//
// 	// Start t.Run > Loop...
// 	for _, CurrentContent := range thisDefaultTest {
// 		thisTestName := fmt.Sprintf("%s_%s_%s%d", thisTestPrefixName, CurrentContent.Category, CurrentContent.Provider, CurrentContent.Id)
//
// 		t.Run(thisTestName, func(t *testing.T) {
// 			// t.Parallel()
//
// 			// Log Test Name
// 			ghafpkg.GHPrintLnWithColor(ghafpkg.GHPrintColorPaletteWarning, thisTestName)
//
// 			// Insert Different Data
// 			thisRequestData.Mobile = CurrentContent.Content
//
// 			// Send To Server
// 			// _ = thisResponseServerBody
// 			thisResponseServer, thisResponseServerBody, thisResponseController, err := methods.TestDTOSendRequestToFiberServer(
// 				t,
// 				thisServer,
// 				&UsersUpdate{},
// 				thisRequestData,
// 				"/1",
// 				"PATCH",
// 				nil,
// 				nil,
// 			)
// 			if err != nil {
// 				t.Fatalf("[%s] DTO Send Request To Fiber Server FAILD: %v", thisTestPrefixName, err)
// 			}
//
// 			// Print Details
// 			ghafpkg.GHPrintJsonLnWithColor(ghafpkg.GHPrintColorPaletteDebug, "thisResponseServerBody:", thisResponseServerBody)
// 			ghafpkg.GHPrintLnWithColor(ghafpkg.GHPrintColorPaletteDebug, "scenario:", CurrentContent.Scenario)
//
// 			// Expected
// 			if CurrentContent.Expected.Content != "" {
// 				assert.Equal(t, CurrentContent.Expected.Content, thisResponseController.Mobile)
// 			}
// 			assert.Equal(t, CurrentContent.Expected.StatusCode, thisResponseServer.StatusCode)
// 		})
// 	}
// }

func TestUsersDTOUpdatePhone(t *testing.T) {
	// Disable Log
	methods.LogActiveStatus(false)

	// Variables
	thisTestPrefixName := "adminmembers_dto_request_update_phone"

	// init server
	thisServer := fiberServerSetup(t)

	// Init Request Struct Data
	thisRequestData := UsersUpdate{
		Name:   "ABCDEFG",
		Family: "GFEDCBA",
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
			thisRequestData.Phone = &CurrentContent.Content

			// Send To Server
			// _ = thisResponseServerBody
			thisResponseServer, thisResponseServerBody, thisResponseController, err := methods.TestDTOSendRequestToFiberServer(
				t,
				thisServer,
				&UsersUpdate{},
				thisRequestData,
				"/1",
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
				assert.Equal(t, CurrentContent.Expected.Content, *thisResponseController.Phone)
			}
			assert.Equal(t, CurrentContent.Expected.StatusCode, thisResponseServer.StatusCode)
		})
	}
}

func TestUsersDTOUpdateEmail(t *testing.T) {
	// Disable Log
	methods.LogActiveStatus(false)

	// Variables
	thisTestPrefixName := "adminmembers_dto_request_update_email"

	// init server
	thisServer := fiberServerSetup(t)

	// Init Request Struct Data
	thisRequestData := UsersUpdate{
		Name:   "ABCDEFG",
		Family: "GFEDCBA",
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
			thisRequestData.Email = &CurrentContent.Content

			// Send To Server
			// _ = thisResponseServerBody
			thisResponseServer, thisResponseServerBody, thisResponseController, err := methods.TestDTOSendRequestToFiberServer(
				t,
				thisServer,
				&UsersUpdate{},
				thisRequestData,
				"/1",
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
				assert.Equal(t, CurrentContent.Expected.Content, *thisResponseController.Email)
			}
			assert.Equal(t, CurrentContent.Expected.StatusCode, thisResponseServer.StatusCode)
		})
	}
}

func TestUsersDTOUpdateBirthDate(t *testing.T) {
	// Disable Log
	methods.LogActiveStatus(false)

	// Variables
	thisTestPrefixName := "adminmembers_dto_request_update_birth-date"

	// init server
	thisServer := fiberServerSetup(t)

	// Init Request Struct Data
	thisRequestData := UsersUpdate{
		Name:   "ABCDEFG",
		Family: "GFEDCBA",
	}

	// Init Default Test Slice
	var thisDefaultTest = defaulttestpkgdtobirthdate.NewDefaultTestDTOBirthdate()

	// Start t.Run > Loop...
	for _, CurrentContent := range thisDefaultTest {
		thisTestName := fmt.Sprintf("%s_%s_%s%d", thisTestPrefixName, CurrentContent.Category, CurrentContent.Provider, CurrentContent.Id)

		t.Run(thisTestName, func(t *testing.T) {
			// t.Parallel()

			// Log Test Name
			ghafpkg.GHPrintLnWithColor(ghafpkg.GHPrintColorPaletteWarning, thisTestName)

			// Insert Different Data
			thisRequestData.BirthDate = &CurrentContent.Content

			// Send To Server
			// _ = thisResponseServerBody
			thisResponseServer, thisResponseServerBody, thisResponseController, err := methods.TestDTOSendRequestToFiberServer(
				t,
				thisServer,
				&UsersUpdate{},
				thisRequestData,
				"/1",
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
				assert.Equal(t, CurrentContent.Expected.Content, *thisResponseController.BirthDate)
			}
			assert.Equal(t, CurrentContent.Expected.StatusCode, thisResponseServer.StatusCode)
		})
	}
}

func TestUsersDTOUpdateNationalCode(t *testing.T) {
	// Disable Log
	methods.LogActiveStatus(false)

	// Variables
	thisTestPrefixName := "adminmembers_dto_request_update_national-code"

	// init server
	thisServer := fiberServerSetup(t)

	// Init Request Struct Data
	thisRequestData := UsersUpdate{
		Name:   "ABCDEFG",
		Family: "GFEDCBA",
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
			thisRequestData.NationalCode = &CurrentContent.Content

			// Send To Server
			// _ = thisResponseServerBody
			thisResponseServer, thisResponseServerBody, thisResponseController, err := methods.TestDTOSendRequestToFiberServer(
				t,
				thisServer,
				&UsersUpdate{},
				thisRequestData,
				"/1",
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
				assert.Equal(t, CurrentContent.Expected.Content, *thisResponseController.NationalCode)
			}
			assert.Equal(t, CurrentContent.Expected.StatusCode, thisResponseServer.StatusCode)
		})
	}
}

func TestUsersDTOUpdateProfileImage(t *testing.T) {
	// Disable Log
	methods.LogActiveStatus(false)

	// Variables
	thisTestPrefixName := "adminmembers_dto_request_update_profile-image"

	// init server
	thisServer := fiberServerSetup(t)

	// Init Request Struct Data
	thisRequestData := UsersUpdate{
		Name:   "ABCDEFG",
		Family: "GFEDCBA",
	}

	// Init Default Test Slice
	var thisDefaultTest = defaulttestpkgdtourl.NewDefaultTestDTOURL()

	// Start t.Run > Loop...
	for _, CurrentContent := range thisDefaultTest {
		thisTestName := fmt.Sprintf("%s_%s_%s%d", thisTestPrefixName, CurrentContent.Category, CurrentContent.Provider, CurrentContent.Id)

		t.Run(thisTestName, func(t *testing.T) {
			// t.Parallel()

			// Log Test Name
			ghafpkg.GHPrintLnWithColor(ghafpkg.GHPrintColorPaletteWarning, thisTestName)

			// Insert Different Data
			thisRequestData.ProfileImage = &CurrentContent.Content

			// Send To Server
			// _ = thisResponseServerBody
			thisResponseServer, thisResponseServerBody, thisResponseController, err := methods.TestDTOSendRequestToFiberServer(
				t,
				thisServer,
				&UsersUpdate{},
				thisRequestData,
				"/1",
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
				assert.Equal(t, CurrentContent.Expected.Content, *thisResponseController.ProfileImage)
			}
			assert.Equal(t, CurrentContent.Expected.StatusCode, thisResponseServer.StatusCode)
		})
	}
}
