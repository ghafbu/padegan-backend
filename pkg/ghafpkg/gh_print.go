package ghafpkg

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/gookit/color"
)

const (
	GHPrintColorPaletteWarning = 1
	GHPrintColorPaletteInfo    = 2
	GHPrintColorPaletteDebug   = 3
	GHPrintColorPaletteFatal   = 4
)

func findPalette(paramColorPalette int) (color.Color, color.Color) {
	switch paramColorPalette {
	case GHPrintColorPaletteWarning:
		return color.Black, color.BgYellow
	case GHPrintColorPaletteInfo:
		return color.Black, color.BgCyan
	case GHPrintColorPaletteDebug:
		return color.Black, color.BgMagenta
	case GHPrintColorPaletteFatal:
		return color.Black, color.BgRed
	default:
		return color.Black, color.BgCyan
	}
}

func GHPrintLn(args ...interface{}) {
	log.Println(args...)
}

func GHPrintLnWithColor(paramColorPalette int, paramData ...interface{}) {
	var thisColor, thisBGColor = findPalette(paramColorPalette)

	for _, d := range paramData {
		log.Println(prepareMultilineForColor(fmt.Sprintf("%v", d), thisColor, thisBGColor))
	}
}

func GHPrintF(format string, args ...interface{}) {
	log.Printf(format, args...)
}

func GHPrintFWithColor(paramColorPalette int, paramFormat string, paramData ...interface{}) {
	var thisColor, thisBGColor = findPalette(paramColorPalette)
	var thisNewString = fmt.Sprintf(paramFormat, paramData...)
	log.Println(prepareMultilineForColor(thisNewString, thisColor, thisBGColor))
}

func GHPrintJsonLn(paramData ...interface{}) {
	for _, d := range paramData {
		var formattedJSON string

		// اگر مقدار از نوع رشته باشد، بررسی کنید که آیا JSON معتبر است
		if str, ok := d.(string); ok && strings.HasPrefix(strings.TrimSpace(str), "{") {
			var jsonData map[string]interface{}
			if err := json.Unmarshal([]byte(str), &jsonData); err == nil {
				// اگر مقدار رشته، JSON معتبر باشد، آن را مرتب‌سازی کنید
				prettyJSON, _ := json.MarshalIndent(jsonData, "", "  ")
				formattedJSON = string(prettyJSON)
			} else {
				// اگر رشته JSON معتبر نبود، همان مقدار را چاپ کن
				formattedJSON = str
			}
		} else {
			// مقدار غیررشته‌ای را مستقیماً JSON مرتب کن
			jsonBytes, err := json.MarshalIndent(d, "", "  ")
			if err != nil {
				log.Printf("GHPrint Error: JSON: %v\n", err)
				return
			}
			formattedJSON = string(jsonBytes)
		}

		// نمایش خروجی نهایی
		log.Println(formattedJSON)
	}
}

func GHPrintJsonLnWithColor(paramColorPalette int, paramData ...interface{}) {
	var thisColor, thisBGColor = findPalette(paramColorPalette)

	for _, d := range paramData {
		var formattedJSON string

		// اگر مقدار از نوع رشته باشد، بررسی کنید که آیا JSON معتبر است
		if str, ok := d.(string); ok && strings.HasPrefix(strings.TrimSpace(str), "{") {
			var jsonData map[string]interface{}
			if err := json.Unmarshal([]byte(str), &jsonData); err == nil {
				// اگر مقدار رشته، JSON معتبر باشد، آن را مرتب‌سازی کنید
				prettyJSON, _ := json.MarshalIndent(jsonData, "", "  ")
				formattedJSON = string(prettyJSON)
			} else {
				// اگر رشته JSON معتبر نبود، همان مقدار را چاپ کن
				formattedJSON = str
			}
		} else {
			// مقدار غیررشته‌ای را مستقیماً JSON مرتب کن
			jsonBytes, err := json.MarshalIndent(d, "", "  ")
			if err != nil {
				log.Printf("GHPrint Error: JSON: %v\n", err)
				return
			}
			formattedJSON = string(jsonBytes)
		}

		// نمایش خروجی نهایی
		log.Println(prepareMultilineForColor(formattedJSON, thisColor, thisBGColor))
	}
}

const resetCode = "\033[0m"

func prepareMultilineForColor(input string, fg color.Color, bg ...color.Color) string {
	var style color.Style
	if len(bg) > 0 {
		style = color.New(fg, bg[0])
	} else {
		style = color.New(fg)
	}

	lines := strings.Split(input, "\n")
	for i, line := range lines {
		// اضافه کردن resetCode در پایان هر خط
		lines[i] = style.Render(line) + resetCode
	}
	return strings.Join(lines, "\n")
}
