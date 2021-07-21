package models

import "casorder/utils/types"

type Health struct {
	StatusCode int
	Message    string
}

func (h *Health) Serialize() types.JSON {
	return types.JSON{
		"status":  h.StatusCode,
		"message": h.Message,
	}
}

func (h *Health) Read(m types.JSON) {
	h.StatusCode = int(m["status_code"].(int))
	h.Message = m["message"].(string)
}
