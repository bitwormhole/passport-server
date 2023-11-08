package controllers

import (
	"github.com/bitwormhole/passport-server/app/data/credentials"
	"github.com/gin-gonic/gin"
	"github.com/starter-go/libgin"
)

////////////////////////////////////////////////////////////////////////////////

// CredentialController ...
type CredentialController struct {

	//starter:component
	_as func(libgin.Controller) //starter:as(".")

	Responder libgin.Responder    //starter:inject("#")
	Service   credentials.Service //starter:inject("#")

}

func (inst *CredentialController) _impl() libgin.Controller {
	return inst
}

// Registration ...
func (inst *CredentialController) Registration() *libgin.ControllerRegistration {
	return &libgin.ControllerRegistration{
		Route: inst.route,
	}
}

func (inst *CredentialController) route(g libgin.RouterProxy) error {

	g = g.For("credentials")

	g.POST("", inst.handle)
	g.DELETE(":id", inst.handle)
	g.PUT(":id", inst.handle)
	g.GET(":id", inst.handle)
	g.GET("", inst.handle)

	return nil
}

func (inst *CredentialController) handle(c *gin.Context) {
	task := &myCredentialTask{
		context:         c,
		parent:          inst,
		wantRequestID:   false,
		wantRequestBody: false,
	}
	task.exec(task.doAction)
}

////////////////////////////////////////////////////////////////////////////////

type myCredentialTask struct {
	context *gin.Context
	parent  *CredentialController

	wantRequestID   bool
	wantRequestBody bool

	id    credentials.ID
	body1 credentials.VO
	body2 credentials.VO
}

func (inst *myCredentialTask) open() error {

	c := inst.context

	if inst.wantRequestBody {
		err := c.BindJSON(&inst.body1)
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *myCredentialTask) send(err error) {
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

func (inst *myCredentialTask) exec(fn func() error) {
	err := inst.open()
	if err == nil {
		err = fn()
	}
	inst.send(err)
}

func (inst *myCredentialTask) doAction() error {
	return nil
}

////////////////////////////////////////////////////////////////////////////////
