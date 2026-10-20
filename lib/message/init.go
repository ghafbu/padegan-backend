package messages

import "github.com/ghafbu/padegan-backend/pkg/messagepkg"

func Init() {
	newKeywords := messagepkg.KeywordStruct{
		Field: map[string]messagepkg.Language{
			// "": {
			//	Farsi:   "",
			//	English: "",
			// },
			"admin_member": {
				Farsi:   "عضو مدیریت",
				English: "admin member",
			},
			"admin_members": {
				Farsi:   "اعضای مدیریت",
				English: "admin members",
			},
		},
		Section: map[string]messagepkg.Language{
			"admin_member": {
				Farsi:   "عضو مدیریت",
				English: "admin member",
			},
			"admin_members": {
				Farsi:   "اعضای مدیریت",
				English: "admin members",
			},
		},
	}
	messagepkg.UpdateKeywords(&newKeywords)
}
