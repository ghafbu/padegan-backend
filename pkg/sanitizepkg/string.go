package sanitizepkg

import "html"

// string
func (r Sanitize) String(input string) string {
	//input = html.EscapeString(input)
	return r.BlueMonday.Sanitize(input)
}
func (r Sanitize) StringEscape(input string) string {
	return html.EscapeString(input)
}
func (r Sanitize) StringUnEscape(input string) string {
	return html.UnescapeString(input)
}
