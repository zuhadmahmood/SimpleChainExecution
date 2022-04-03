package controllers

import (
	"github.com/revel/examples/validation/app/models"
	"github.com/revel/revel"
)

type Process3 struct {
	*revel.Controller
}

func (c Process3) Index() revel.Result {
	return c.Render()
}

func (c Process3) HandleSubmit(user *models.User) revel.Result {
	user.Validate(c.Validation)

	// Handle errors
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(Process3.Index)
	}

	// Ok, display the created user
	return c.Render(user)
}
