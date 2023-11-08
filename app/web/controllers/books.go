package controllers

import (
	"github.com/bitwormhole/passport-server/app/data/books"
	"github.com/gin-gonic/gin"
	"github.com/starter-go/libgin"
)

////////////////////////////////////////////////////////////////////////////////

// BookController ...
type BookController struct {

	//starter:component
	_as func(libgin.Controller) //starter:as(".")

	Responder libgin.Responder //starter:inject("#")
	Service   books.Service    //starter:inject("#")
}

func (inst *BookController) _impl() libgin.Controller {
	return inst
}

// Registration ...
func (inst *BookController) Registration() *libgin.ControllerRegistration {
	return &libgin.ControllerRegistration{
		Route: inst.route,
	}
}

func (inst *BookController) route(g libgin.RouterProxy) error {

	g = g.For("books")

	g.POST("", inst.handle)
	g.DELETE(":id", inst.handle)
	g.PUT(":id", inst.handle)
	g.GET(":id", inst.handle)
	g.GET("", inst.handle)

	return nil
}

func (inst *BookController) handle(c *gin.Context) {
	task := &myBookTask{
		context:         c,
		parent:          inst,
		wantRequestID:   false,
		wantRequestBody: false,
	}
	task.exec(task.doAction)
}

////////////////////////////////////////////////////////////////////////////////

type myBookTask struct {
	context *gin.Context
	parent  *BookController

	wantRequestID   bool
	wantRequestBody bool

	body1 books.VO
	body2 books.VO
}

func (inst *myBookTask) open() error {

	c := inst.context

	if inst.wantRequestBody {
		err := c.BindJSON(&inst.body1)
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *myBookTask) send(err error) {
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

func (inst *myBookTask) exec(fn func() error) {
	err := inst.open()
	if err == nil {
		err = fn()
	}
	inst.send(err)
}

func (inst *myBookTask) doAction() error {
	return nil
}

////////////////////////////////////////////////////////////////////////////////
