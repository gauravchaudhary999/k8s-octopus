package model

type OctopusEnvironment struct {
	Environment *Environment
	Account     *Account
	Machine     *Machine
}

type ActionTemplateInput struct {
	Name, Desc, Type string
}
