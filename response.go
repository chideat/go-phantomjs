package phantomjs

import (
	"encoding/json"
)

type Response struct {
	Status    int             `json:"status"`
	SessionId string          `json:"sessionId"`
	Value     json.RawMessage `json:"value,omitempty"`
}
