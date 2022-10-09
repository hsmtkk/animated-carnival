package unixepoch

import (
	"fmt"
	"strconv"
	"time"
)

func StringToTime(microSecsStr string) (time.Time, error) {
	microSecs, err := strconv.ParseInt(microSecsStr, 10, 64)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to parse as int64; %s; %w", microSecsStr, err)
	}
	t := time.UnixMicro(microSecs)
	return t, nil
}
