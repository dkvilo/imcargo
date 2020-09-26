package controller

import (
	"fmt"
	"image"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/dkvilo/imcargo/core"
	"github.com/dkvilo/imcargo/functions"
	"github.com/dkvilo/imcargo/model"
	"github.com/julienschmidt/httprouter"
)

// Upload controller
func (ctrl *Controller) Upload(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	
	w.Header().Set("Content-Type", "text/json")
	queryValuse := r.URL.Query()

	var _type string = "default"
	var blur float64 = 0.0
	var cs image.Point = image.Point{
		X: 128,
		Y: 0,
	};

	if queryValuse.Get("size") != "" {
		sizes := strings.Split(queryValuse.Get("size"), "x")
		cs.X, _ = strconv.Atoi(sizes[0])
		cs.Y, _ = strconv.Atoi(sizes[1])
	}

	if queryValuse.Get("type") != "" {
		_type = queryValuse.Get("type")
	}

	if queryValuse.Get("blur") != "" {
		blurInt, _ := strconv.Atoi(queryValuse.Get("blur"))
		blur = float64(blurInt)
	}

	mf, _, err := r.FormFile("image")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, string(
			core.Response(
				model.ImageObject{
					Success: false,
					Message: fmt.Sprintf("Unable to read Image from form: %s", err.Error()),
				},
			),
		))
		return
	}
	defer mf.Close()

	var avatar image.Image
	switch _type {
	case "default":
		avatar, err = functions.ResizeImage(mf, cs)
		break
	case "centered" :
			avatar, err = functions.CropCenterAnchor(mf, cs)
		break
	default:
		avatar, err = functions.ResizeImage(mf, cs)
		break
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, string(
			core.Response(
				model.ImageObject{
					Success: false,
					Message: fmt.Sprintf("Unable to resize the image: %s", err.Error()),
				},
			),
		))
		return
	}

	if blur > 0.0 {
		avatar, err = functions.BlurImage(avatar, blur)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			io.WriteString(w, string(
				core.Response(
					model.ImageObject{
						Success: false,
						Message: fmt.Sprintf("Unable to blur the image: %s", err.Error()),
					},
				),
			))
			return
		}
	}

	avatarPath, err := functions.SaveImage("static/avatar/", avatar);
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, string(
			core.Response(
				model.ImageObject{
					Success: false,
					Message: fmt.Sprintf("Unable to save the image: %s", err.Error()),
				},
			),
		))
		return
	}

	w.WriteHeader(http.StatusOK)
	io.WriteString(w, string(
		core.Response(
			model.ImageObject{
				Success: true,
				Message: "Avatar was uploaded successfully",
				Data: model.Data {
					Path: avatarPath,
					Size: model.Size {
						Width: avatar.Bounds().Size().X,
						Height: avatar.Bounds().Size().Y,
					},
				},
			},
		),
	))
	
	return
}
