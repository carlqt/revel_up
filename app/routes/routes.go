// GENERATED CODE - DO NOT EDIT
package routes

import "github.com/revel/revel"


type tApp struct {}
var App tApp


func (_ tApp) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Index", args).Url
}


type tSessions struct {}
var Sessions tSessions


func (_ tSessions) New(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Sessions.New", args).Url
}

func (_ tSessions) Create(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Sessions.Create", args).Url
}

func (_ tSessions) Register(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Sessions.Register", args).Url
}

func (_ tSessions) Destroy(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Sessions.Destroy", args).Url
}


type tUser struct {}
var User tUser


func (_ tUser) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("User.Index", args).Url
}

func (_ tUser) New(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("User.New", args).Url
}

func (_ tUser) Create(
		username string,
		email string,
		password string,
		passwordConfirmation string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "username", username)
	revel.Unbind(args, "email", email)
	revel.Unbind(args, "password", password)
	revel.Unbind(args, "passwordConfirmation", passwordConfirmation)
	return revel.MainRouter.Reverse("User.Create", args).Url
}


type tStatic struct {}
var Static tStatic


func (_ tStatic) Serve(
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.Serve", args).Url
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
	return revel.MainRouter.Reverse("Static.ServeModule", args).Url
}


type tTestRunner struct {}
var TestRunner tTestRunner


func (_ tTestRunner) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.Index", args).Url
}

func (_ tTestRunner) Run(
		suite string,
		test string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	revel.Unbind(args, "test", test)
	return revel.MainRouter.Reverse("TestRunner.Run", args).Url
}

func (_ tTestRunner) List(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.List", args).Url
}


