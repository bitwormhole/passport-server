package controllers

import (
	"github.com/bitwormhole/passport-server/app/data/secretkeys"
	"github.com/gin-gonic/gin"
	"github.com/starter-go/libgin"
)

////////////////////////////////////////////////////////////////////////////////

// SecretKeyController ...
type SecretKeyController struct {

	//starter:component
	_as func(libgin.Controller) //starter:as(".")

	Responder libgin.Responder   //starter:inject("#")
	Service   secretkeys.Service //starter:inject("#")

}

func (inst *SecretKeyController) _impl() libgin.Controller {
	return inst
}

// Registration ...
func (inst *SecretKeyController) Registration() *libgin.ControllerRegistration {
	return &libgin.ControllerRegistration{
		Route: inst.route,
	}
}

func (inst *SecretKeyController) route(g libgin.RouterProxy) error {

	g = g.For("secret-keys")

	g.POST("", inst.handle)
	g.DELETE(":id", inst.handle)
	g.PUT(":id", inst.handle)
	g.GET(":id", inst.handle)
	g.GET("", inst.handle)

	return nil
}

func (inst *SecretKeyController) handle(c *gin.Context) {
	task := &mySecretKeyTask{
		context:         c,
		parent:          inst,
		wantRequestID:   false,
		wantRequestBody: false,
	}
	task.exec(task.doAction)
}

////////////////////////////////////////////////////////////////////////////////

type mySecretKeyTask struct {
	context *gin.Context
	parent  *SecretKeyController

	wantRequestID   bool
	wantRequestBody bool

	id    secretkeys.ID
	body1 secretkeys.VO
	body2 secretkeys.VO
}

func (inst *mySecretKeyTask) open() error {

	c := inst.context

	if inst.wantRequestBody {
		err := c.BindJSON(&inst.body1)
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *mySecretKeyTask) send(err error) {
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

func (inst *mySecretKeyTask) exec(fn func() error) {
	err := inst.open()
	if err == nil {
		err = fn()
	}
	inst.send(err)
}

func (inst *mySecretKeyTask) doAction() error {
	return nil
}

////////////////////////////////////////////////////////////////////////////////
