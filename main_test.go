package time_iso
import (
	"testing"
	"time"
	"fmt"
)
var (
	formats = map[string]string {
		"ANSIC"      : "Mon Jan _2 15:04:05 2006",
		"UnixDate"   : "Mon Jan _2 15:04:05 MST 2006",
		"RubyDate"   : "Mon Jan 02 15:04:05 -0700 2006",
		"RFC822"     : "02 Jan 06 15:04 MST",
		"RFC822Z"    : "02 Jan 06 15:04 -0700",
		"RFC850"     : "Monday, 02-Jan-06 15:04:05 MST",
		"RFC1123"    : "Mon, 02 Jan 2006 15:04:05 MST",
		"RFC1123Z"   : "Mon, 02 Jan 2006 15:04:05 -0700",
		"Stamp"      : "Jan _2 15:04:05",
		"StampMilli" : "Jan _2 15:04:05.000",
		"DateTime"   : "2006-01-02 15:04:05",
		"DateOnly"   : "2006-01-02",
		"TimeOnly"   : "15:04:05",
	}
	formats_iso = map[string]string {
		"ANSIC"      : "DDD MMM _D hh:mm:ss YYYY",
		"UnixDate"   : "DDD MMM _D hh:mm:ss ZZZ YYYY",
		"RubyDate"   : "DDD MMM DD hh:mm:ss Z YYYY",
		"RFC822"     : "DD MMM YY hh:mm ZZZ",
		"RFC822Z"    : "DD MMM YY hh:mm Z",
		"RFC850"     : "DDDD, DD-MMM-YY hh:mm:ss ZZZ",
		"RFC1123"    : "DDD, DD MMM YYYY hh:mm:ss ZZZ",
		"RFC1123Z"   : "DDD, DD MMM YYYY hh:mm:ss Z",
		"Stamp"      : "MMM _D hh:mm:ss",
		"StampMilli" : "MMM _D hh:mm:ss.sss",
		"DateTime"   : "YYYY-MM-DD hh:mm:ss",
		"DateOnly"   : "YYYY-MM-DD",
		"TimeOnly"   : "hh:mm:ss",
	}
)
func TestTime(t *testing.T) {
	code := false
	now_old := time.Now()
	now := Time{now_old}
	for key, val := range formats {
		message := "\r\t\t\t\t\t\t"
		if now_old.Format(val) == now.Format(formats_iso[key]) {
			message = "succeeded" + message + "✅"
		} else {
			code = true
			message = "failed" + message + "❌"
		}
		fmt.Printf("The test for the %s format %s\n", key, message)
	}
	if code {
		t.Error("One or more of the tests failed")
	}
}