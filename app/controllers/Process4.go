package controllers

import (
	"github.com/revel/examples/validation/app/models"
	"github.com/revel/revel"
)

type Process4 struct {
	*revel.Controller
}

func (c Process4) Index() revel.Result {
	return c.Render()
}

func (c Process4) HandleSubmit(user *models.User) revel.Result {
	user.Validate(c.Validation)

	// Handle errors
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(Process4.Index)
	}

	// Ok, display the created user
	return c.Render(user)
}
