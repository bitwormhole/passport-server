package controllers

import (
	"github.com/bitwormhole/passport-server/app/data/publickeys"
	"github.com/gin-gonic/gin"
	"github.com/starter-go/libgin"
)

////////////////////////////////////////////////////////////////////////////////

// PublicKeyController ...
type PublicKeyController struct {

	//starter:component
	_as func(libgin.Controller) //starter:as(".")

	Responder libgin.Responder   //starter:inject("#")
	Service   publickeys.Service //starter:inject("#")

}

func (inst *PublicKeyController) _impl() libgin.Controller {
	return inst
}

// Registration ...
func (inst *PublicKeyController) Registration() *libgin.ControllerRegistration {
	return &libgin.ControllerRegistration{
		Route: inst.route,
	}
}

func (inst *PublicKeyController) route(g libgin.RouterProxy) error {

	g = g.For("public-keys")

	g.POST("", inst.handle)
	g.DELETE(":id", inst.handle)
	g.PUT(":id", inst.handle)
	g.GET(":id", inst.handle)
	g.GET("", inst.handle)

	return nil
}

func (inst *PublicKeyController) handle(c *gin.Context) {
	task := &myPublicKeyTask{
		context:         c,
		parent:          inst,
		wantRequestID:   false,
		wantRequestBody: false,
	}
	task.exec(task.doAction)
}

////////////////////////////////////////////////////////////////////////////////

type myPublicKeyTask struct {
	context *gin.Context
	parent  *PublicKeyController

	wantRequestID   bool
	wantRequestBody bool

	id    publickeys.ID
	body1 publickeys.VO
	body2 publickeys.VO
}

func (inst *myPublicKeyTask) open() error {

	c := inst.context

	if inst.wantRequestBody {
		err := c.BindJSON(&inst.body1)
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *myPublicKeyTask) send(err error) {
	c := inst.context
	d := &inst.body2
	s := inst.body2.Status
	r := &libgin.Response{
		Context: c,
		Data:    d,
		Error:   err,
		Status:  s,
	}
	inst.parent.Responder.Send(r)
}

func (inst *myPublicKeyTask) exec(fn func() error) {
	err := inst.open()
	if err == nil {
		err = fn()
	}
	inst.send(err)
}

func (inst *myPublicKeyTask) doAction() error {
	return nil
}

////////////////////////////////////////////////////////////////////////////////
