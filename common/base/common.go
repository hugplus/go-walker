package base

import "fmt"

func FmtReqId(reqId string) string {
	return fmt.Sprintf("REQID:%s", reqId)
}
