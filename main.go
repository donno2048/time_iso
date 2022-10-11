package time_iso
import (
	"strings"
	"sort"
	"time"
)
var replacements = map[string]string {
	"DD"   : "02",
	"DDD"  : "Mon",
	"DDDD" : "Monday",
	"MM"   : "01",
	"MMM"  : "Jan",
	"YY"   : "06",
	"YYYY" : "2006",
	"Z"    : "-0700",
	"ZZZ"  : "MST",
	"_D"   : "_2",
	"hh"   : "15",
	"mm"   : "04",
	"ss"   : "05",
	"sss"  : "000",
}
type Time struct {time.Time}
func (t Time) Format(format string) string {
	layout := format
	keys := make([]string, 0, len(replacements))
	for k := range replacements {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {return len(keys[i]) > len(keys[j])})
	for _, key := range keys {
		layout = strings.Replace(layout, key, replacements[key], 2)
	}
	return t.Time.Format(layout)
}
func Now() Time {
	return Time{time.Now()}
}