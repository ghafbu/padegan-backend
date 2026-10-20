package types

type MediaBrowseImage struct {
	DriverId           *uint64                `json:"driver_id"`                                               // 0=external , 1,2,3,4... = internal
	ExternalDriverLink *string                `json:"external_driver_link"`                                    // https://google.com
	Path               string                 `json:"path" before_validation:"trim" validate:"gh_printable"`   // /folder1/folder2/folder3
	Name               string                 `json:"name" before_validation:"trim" validate:"gh_printable"`   // food
	Format             string                 `json:"format" before_validation:"trim" validate:"gh_printable"` // png
	Size               float64                `json:"size"`                                                    // 450 [kb]
	Width              float64                `json:"width"`                                                   // 435.36
	Height             float64                `json:"height"`                                                  // 435.36
	Rate               string                 `json:"rate"`                                                    // 9:6 , 1:1 , 4:3 , 3:4 ....
	Alt                string                 `json:"alt"`                                                     // download my brand logo [use in seo]
	Title              string                 `json:"title"`                                                   // download my brand logo [use in seo]
	Thumbnail          *MediaDetailsThumbnail `json:"thumbnail"`
}

type MediaBrowseVideo struct {
	DriverId           *uint64                         `json:"driver_id"`                                               // 0=external , 1,2,3,4... = internal
	ExternalDriverLink *string                         `json:"external_driver_link"`                                    // https://google.com
	Path               string                          `json:"path" before_validation:"trim" validate:"gh_printable"`   // /folder1/folder2/folder3
	Name               string                          `json:"name" before_validation:"trim" validate:"gh_printable"`   // food
	Format             string                          `json:"format" before_validation:"trim" validate:"gh_printable"` // mp4
	Size               float64                         `json:"size"`                                                    // 450 [kb]
	Time               float64                         `json:"width"`                                                   // 125 [second]
	Rate               string                          `json:"rate"`                                                    // 9:6 , 1:1 , 4:3 , 3:4 ....
	Thumbnail          *MediaDetailsThumbnail          `json:"thumbnail"`
	Summarization      *MediaDetailsVideoSummarization `json:"summarization"`
}

type MediaBrowseDocument struct {
	DriverId           *uint64 `json:"driver_id"`                                               // 0=external , 1,2,3,4... = internal
	ExternalDriverLink *string `json:"external_driver_link"`                                    // https://google.com
	Path               string  `json:"path" before_validation:"trim" validate:"gh_printable"`   // /folder1/folder2/folder3
	Name               string  `json:"name" before_validation:"trim" validate:"gh_printable"`   // food
	Format             string  `json:"format" before_validation:"trim" validate:"gh_printable"` // pdf,word
	Size               float64 `json:"size"`                                                    // 450 [kb]
}

type MediaDetailsThumbnail struct {
	DriverId           *uint64 `json:"driver_id"`                                               // 0=external , 1,2,3,4... = internal
	ExternalDriverLink *string `json:"external_driver_link"`                                    // https://google.com
	Path               string  `json:"path" before_validation:"trim" validate:"gh_printable"`   // /folder1/folder2/folder3
	Name               string  `json:"name" before_validation:"trim" validate:"gh_printable"`   // food-image
	Format             string  `json:"format" before_validation:"trim" validate:"gh_printable"` // png
	Width              float64 `json:"width"`                                                   // 435.36
	Height             float64 `json:"height"`                                                  // 435.36
	// DriverLink + Path + Name + Format
}

type MediaDetailsVideoSummarization struct {
	DriverId           *uint64 `json:"driver_id"`                                               // 0=external, 1,2,3,4... = internal
	ExternalDriverLink *string `json:"external_driver_link"`                                    // https://google.com
	Path               string  `json:"path" before_validation:"trim" validate:"gh_printable"`   // /folder1/folder2/folder3
	Name               string  `json:"name" before_validation:"trim" validate:"gh_printable"`   // food-video
	Format             string  `json:"format" before_validation:"trim" validate:"gh_printable"` // mp4
	// DriverLink + Path + Name + Format
}
