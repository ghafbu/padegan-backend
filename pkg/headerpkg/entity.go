package headerpkg

type Header struct {
	Host            []string // EX: developer.mozilla.org
	UserAgent       []string // EX: Mozilla/5.0 (Macintosh; Intel Mac OS X 10.9; rv:50.0) Gecko/20100101 Firefox/50.0
	Accept          []string // EX: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8
	AcceptEncoding  []string // EX: gzip, deflate, br
	AcceptLanguage  []string // EX: en-US,en;q=0.5
	Connection      []string // EX: keep-alive
	CacheControl    []string // EX: max-age=0
	Referer         []string // EX: https://developer.mozilla.org/testpage.html
	ContentType     []string // EX: [application/json]
	ContentLength   []string // EX: 396
	Cookie          []string // EX: MUID=0A5B00E8417A6FA1153311A140116EC7
	SecChUa         []string // EX: "Chromium";v="124", "Google Chrome";v="124", "Not-A.Brand";v="99"
	SecChUaMobile   []string // EX: ?0
	SecChUaPlatform []string // EX: "Windows"
	SecFetchDest    []string // EX: empty
	SecFetchMode    []string // EX: cors
	SecFetchSite    []string // EX: cross-site
	CurrentUrl      []string // EX: https://www.simorgh.app/user/create
	SLanguage       []string // EX: farsi || english
}
