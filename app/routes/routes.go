// GENERATED CODE - DO NOT EDIT
// This file provides a way of creating URL's based on all the actions
// found in all the controllers.
package routes

import "github.com/revel/revel"


type tStatic struct {}
var Static tStatic


func (_ tStatic) Serve(
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.Serve", args).URL
}

func (_ tStatic) ServeDir(
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeDir", args).URL
}

func (_ tStatic) ServeModule(
		moduleName string,
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "moduleName", moduleName)
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeModule", args).URL
}

func (_ tStatic) ServeModuleDir(
		moduleName string,
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "moduleName", moduleName)
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeModuleDir", args).URL
}


type tTestRunner struct {}
var TestRunner tTestRunner


func (_ tTestRunner) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.Index", args).URL
}

func (_ tTestRunner) Suite(
		suite string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	return revel.MainRouter.Reverse("TestRunner.Suite", args).URL
}

func (_ tTestRunner) Run(
		suite string,
		test string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	revel.Unbind(args, "test", test)
	return revel.MainRouter.Reverse("TestRunner.Run", args).URL
}

func (_ tTestRunner) List(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.List", args).URL
}


type tApplication struct {}
var Application tApplication


func (_ tApplication) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Application.Index", args).URL
}


type tProcess1 struct {}
var Process1 tProcess1


func (_ tProcess1) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Process1.Index", args).URL
}

func (_ tProcess1) HandleSubmit(
		username string,
		firstname string,
		lastname string,
		age int,
		password string,
		passwordConfirm string,
		email string,
		emailConfirm string,
		termsOfUse bool,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "username", username)
	revel.Unbind(args, "firstname", firstname)
	revel.Unbind(args, "lastname", lastname)
	revel.Unbind(args, "age", age)
	revel.Unbind(args, "password", password)
	revel.Unbind(args, "passwordConfirm", passwordConfirm)
	revel.Unbind(args, "email", email)
	revel.Unbind(args, "emailConfirm", emailConfirm)
	revel.Unbind(args, "termsOfUse", termsOfUse)
	return revel.MainRouter.Reverse("Process1.HandleSubmit", args).URL
}


type tProcess2 struct {}
var Process2 tProcess2


func (_ tProcess2) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Process2.Index", args).URL
}

func (_ tProcess2) HandleSubmit(
		username string,
		firstname string,
		lastname string,
		age int,
		password string,
		passwordConfirm string,
		email string,
		emailConfirm string,
		termsOfUse bool,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "username", username)
	revel.Unbind(args, "firstname", firstname)
	revel.Unbind(args, "lastname", lastname)
	revel.Unbind(args, "age", age)
	revel.Unbind(args, "password", password)
	revel.Unbind(args, "passwordConfirm", passwordConfirm)
	revel.Unbind(args, "email", email)
	revel.Unbind(args, "emailConfirm", emailConfirm)
	revel.Unbind(args, "termsOfUse", termsOfUse)
	return revel.MainRouter.Reverse("Process2.HandleSubmit", args).URL
}


type tProcess3 struct {}
var Process3 tProcess3


func (_ tProcess3) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Process3.Index", args).URL
}

func (_ tProcess3) HandleSubmit(
		user interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "user", user)
	return revel.MainRouter.Reverse("Process3.HandleSubmit", args).URL
}


type tProcess4 struct {}
var Process4 tProcess4


func (_ tProcess4) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Process4.Index", args).URL
}

func (_ tProcess4) HandleSubmit(
		user interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "user", user)
	return revel.MainRouter.Reverse("Process4.HandleSubmit", args).URL
}


