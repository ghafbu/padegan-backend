package consts

const (
	ErrorKeyModelDatabaseConnectionInit = "database_connection_is_not_initialized"

	ErrorKeyModelAnonymous            = "anonymous_error"
	ModelDeletedRelation              = "related_record_deleted"
	ErrorKeyModelRecordNotFound       = "record_not_found"
	ErrorKeyModelBirthdatePars        = "birthdate_parse"
	ErrorKeyModelPasswordHashGenerate = "password_hash_generate_error"
	ErrorKeyModelIDBadRequest         = "id_bad_request"

	UsersErrorKeyModelExistMobile       = "uni_admin_members_mobile"
	UsersErrorKeyModelExistNationalCode = "uni_admin_members_national_code"
	UsersErrorKeyModelExistEmail        = "uni_admin_members_email"
	UsersErrorKeyModelExistUsername     = "uni_admin_members_username"
)
