package poet

import "encoding/json"

func ToJson(obj interface{}) string {
	bts, _ := json.Marshal(obj)
	return string(bts)
}
