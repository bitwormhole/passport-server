package controllers

import (
	"github.com/bitwormhole/passport-server/app/data/terminals"
	"github.com/gin-gonic/gin"
	"github.com/starter-go/libgin"
)

////////////////////////////////////////////////////////////////////////////////

// TerminalController ...
type TerminalController struct {

	//starter:component
	_as func(libgin.Controller) //starter:as(".")

	Responder libgin.Responder  //starter:inject("#")
	Service   terminals.Service //starter:inject("#")

}

func (inst *TerminalController) _impl() libgin.Controller {
	return inst
}

// Registration ...
func (inst *TerminalController) Registration() *libgin.ControllerRegistration {
	return &libgin.ControllerRegistration{
		Route: inst.route,
	}
}

func (inst *TerminalController) route(g libgin.RouterProxy) error {

	g = g.For("terminals")

	g.POST("", inst.handle)
	g.DELETE(":id", inst.handle)
	g.PUT(":id", inst.handle)
	g.GET(":id", inst.handle)
	g.GET("", inst.handle)

	return nil
}

func (inst *TerminalController) handle(c *gin.Context) {
	task := &myTerminalTask{
		context:         c,
		parent:          inst,
		wantRequestID:   false,
		wantRequestBody: false,
	}
	task.exec(task.doAction)
}

////////////////////////////////////////////////////////////////////////////////

type myTerminalTask struct {
	context *gin.Context
	parent  *TerminalController

	wantRequestID   bool
	wantRequestBody bool

	id    terminals.ID
	body1 terminals.VO
	body2 terminals.VO
}

func (inst *myTerminalTask) open() error {

	c := inst.context

	if inst.wantRequestBody {
		err := c.BindJSON(&inst.body1)
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *myTerminalTask) send(err error) {
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

func (inst *myTerminalTask) exec(fn func() error) {
	err := inst.open()
	if err == nil {
		err = fn()
	}
	inst.send(err)
}

func (inst *myTerminalTask) doAction() error {
	return nil
}

////////////////////////////////////////////////////////////////////////////////
