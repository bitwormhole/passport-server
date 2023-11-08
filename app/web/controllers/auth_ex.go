package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/starter-go/libgin"
	"github.com/starter-go/security/auth"
	"github.com/starter-go/security/rbac"
)

////////////////////////////////////////////////////////////////////////////////

// AuthExController ...
type AuthExController struct {

	//starter:component
	_as func(libgin.Controller) //starter:as(".")

	Responder libgin.Responder //starter:inject("#")
	Service2  auth.Service     //starter:inject("#")
	Service1  rbac.AuthService //starter:inject("#")
}

func (inst *AuthExController) _impl() libgin.Controller {
	return inst
}

// Registration ...
func (inst *AuthExController) Registration() *libgin.ControllerRegistration {
	return &libgin.ControllerRegistration{
		Route: inst.route,
	}
}

func (inst *AuthExController) route(g libgin.RouterProxy) error {
	g = g.For("auth")
	g.POST("sign-up-in", inst.handleSignUpIn)
	return nil
}

func (inst *AuthExController) handle(c *gin.Context) {
	task := &myAuthExTask{
		context:         c,
		parent:          inst,
		wantRequestID:   false,
		wantRequestBody: false,
	}
	task.exec(task.doAction)
}

func (inst *AuthExController) handleSignUpIn(c *gin.Context) {
	task := &myAuthExTask{
		context:         c,
		parent:          inst,
		wantRequestID:   false,
		wantRequestBody: true,
	}
	task.exec(task.doSignUpIn)
}

////////////////////////////////////////////////////////////////////////////////

type myAuthExTask struct {
	context *gin.Context
	parent  *AuthExController

	wantRequestID   bool
	wantRequestBody bool

	body1 rbac.AuthVO
	body2 rbac.AuthVO
}

func (inst *myAuthExTask) open() error {

	c := inst.context

	if inst.wantRequestBody {
		err := c.BindJSON(&inst.body1)
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *myAuthExTask) send(err error) {
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

func (inst *myAuthExTask) exec(fn func() error) {
	err := inst.open()
	if err == nil {
		err = fn()
	}
	inst.send(err)
}

func (inst *myAuthExTask) doAction() error {
	return fmt.Errorf("no impl")
}

func (inst *myAuthExTask) doSignUpIn() error {
	ctx := inst.context
	ser := inst.parent.Service1
	return ser.Handle(ctx, "sign-up-in", inst.body1.Auth)
}

////////////////////////////////////////////////////////////////////////////////
