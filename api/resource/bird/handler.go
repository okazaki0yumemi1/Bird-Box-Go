package bird

import "net/http"

type API struct{}

// List godoc
//
//	@summary        List detections
//	@description    List detected birds
//	@tags           birds
//	@accept         json
//	@produce        json
//	@success        200 {array}     DTO
//	@failure        500 {object}    err.Error
//	@router         /birds [get]
func (a *API) List(w http.ResponseWriter, r *http.Request) {}

// // Create godoc
// //
// //  @summary        Create bird detection
// //  @description    Create bird detection - for debugging!
// //  @tags           birds
// //  @accept         json
// //  @produce        json
// //  @param          body    body    Form    true    "Book form"
// //  @success        201
// //  @failure        400 {object}    err.Error
// //  @failure        422 {object}    err.Errors
// //  @failure        500 {object}    err.Error
// //  @router         /birds [post]
func (a *API) Create(w http.ResponseWriter, r *http.Request) {}

// func (a *API) Update(w http.ResponseWriter, r *http.Request) {}

func (a *API) Delete(w http.ResponseWriter, r *http.Request) {}
