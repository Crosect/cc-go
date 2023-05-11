package response

import (
	"github.com/crosect/cc-go/web/render"
	"net/http"
)

func Write(w http.ResponseWriter, res Response) {
	render.Render(w, res.Meta.HttpStatus(), render.JSON{Data: res})
}

func WriteError(w http.ResponseWriter, err error) {
	Write(w, Error(err))
}
