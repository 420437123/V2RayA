package logs

import (
	"time"
)

func format(s string) string {
	return time.Now().Format("2006-01-02 15:04:05") + " " + s
}
//
//func Print(s string) {
//	_, _ = persistence.DoAndSave("rpush", "V2RayA.log", format(s))
//}
