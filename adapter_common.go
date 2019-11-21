package sheru

import (
	"fmt"
)

type adapterCommon struct{}

func (adapterCommon) quote(key string) string {
	return fmt.Sprintf(`"%s"`, key)
}
