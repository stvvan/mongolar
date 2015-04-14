package redirect

import (
	"github.com/jasonrichardsmith/mongolar/wrapper"
)

func set(r string, w *wrapper.Wrapper) {
	v := make(map[string]interface{})
	v["value"] = r
	w.SetPayload("redirect", v)
}
