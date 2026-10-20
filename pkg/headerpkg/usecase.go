package headerpkg

func New(header map[string][]string) Header {
	return Header{
		Host:            ExistHeader(header, "ServiceHost"),
		UserAgent:       ExistHeader(header, "User-Agent"),
		Accept:          ExistHeader(header, "Accept"),
		AcceptEncoding:  ExistHeader(header, "Accept-Encoding"),
		AcceptLanguage:  ExistHeader(header, "Accept-Language"),
		Connection:      ExistHeader(header, "Connection"),
		CacheControl:    ExistHeader(header, "Redis-Control"),
		Referer:         ExistHeader(header, "Referer"),
		ContentType:     ExistHeader(header, "Content-Type"),
		ContentLength:   ExistHeader(header, "Content-Length"),
		Cookie:          ExistHeader(header, "Cookie"),
		SecChUa:         ExistHeader(header, "Sec-Ch-Ua"),
		SecChUaMobile:   ExistHeader(header, "Sec-Ch-Ua-Mobile"),
		SecChUaPlatform: ExistHeader(header, "Sec-Ch-Ua-Platform"),
		SecFetchDest:    ExistHeader(header, "Sec-Fetch-Dest"),
		SecFetchMode:    ExistHeader(header, "Sec-Fetch-Mode"),
		SecFetchSite:    ExistHeader(header, "Sec-Fetch-Site"),
		CurrentUrl:      ExistHeader(header, "Current-Url"),
		SLanguage:       ExistHeader(header, "S-Language"),
	}
}

func ExistHeader(header map[string][]string, index string) []string {
	if header[index] != nil {
		return header[index]
	}

	return nil
}

func GetIndexAsString(field []string) string {
	if field != nil {
		return field[0]
	}

	return ""
}

func GetIndexAsStrings(field []string) []string {
	if field != nil {
		return field
	}

	return []string{""}
}
