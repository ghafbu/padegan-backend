package request

type UserSettingsUpdate struct {
	StatusSite       bool `json:"status_site" validate:"gh_required_bool"`
	StatusOrgPanel   bool `json:"status_org_panel" validate:"gh_required_bool"`
	StatusSiteSignUp bool `json:"status_site_sign_up" validate:"gh_required_bool"`
	StatusOrgSite    bool `json:"status_org_site" validate:"gh_required_bool"`
}
