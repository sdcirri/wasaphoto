package api

import "net/http"

func (rt *_router) internalServerError(err error, w http.ResponseWriter) {
	rt.baseLogger.Debug("Internal server error: " + err.Error())
	http.Error(w, "Internal server error: "+err.Error(), http.StatusInternalServerError)
}
