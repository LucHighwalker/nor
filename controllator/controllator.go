package controllator

import (
	"fmt"
	"io/ioutil"
	"path"
	"strings"

	"github.com/urfave/cli"

	"nor/helper"
	"nor/templates"
)

func Command(wd string) *cli.Command {
	return &cli.Command{
		Name:    "controller",
		Aliases: []string{"c"},
		Usage:   "Initialize a new controller.",
		Action: func(c *cli.Context) error {
			generateController(wd, c)
			return nil
		},
	}
}

func generateController(wd string, c *cli.Context) {
	name := c.Args().First()
	actions, routes, hasActions := processArguments(name, c.Args().Tail())

	fmt.Println(routes)

	if hasActions {
		actions = fmt.Sprintf("\n%s", actions)
	}

	controller := templates.Controller(name, actions)
	router := templates.Router(name, routes)

	controllerPath := path.Join(wd, "src", name)

	helper.EnsureDirExists(controllerPath)
	ioutil.WriteFile(path.Join(controllerPath, fmt.Sprintf("%s.controller.ts", name)), controller, 0644)
	ioutil.WriteFile(path.Join(controllerPath, fmt.Sprintf("%s.routes.ts", name)), router, 0644)
}

func processArguments(name string, arguments []string) (string, string, bool) {
	actions := []string{}
	routes := []string{}
	for _, act := range arguments {
		split := strings.Split(act, ":")
		action := split[0]
		actionArgs := []string{}
		routeArgs := []string{}
		routeCall := []string{}
		access := "public"
		verb := "get"

		for i, a := range split {
			if i > 0 {
				switch true {

				case a == "private":
					access = a
					break

				case strings.Contains(a, "="):
					s := strings.Split(a, "=")
					actionArgs = append(actionArgs, fmt.Sprintf("%s: %s", s[0], helper.Capitalize(s[1])))
					routeArgs = append(routeArgs, fmt.Sprintf("/:%s", s[0]))
					routeCall = append(routeCall, s[0])
					break

				case aIsVerb(a):
					verb = a
					break
				}
			}
		}

		actions = append(actions, templates.Action(access, action, strings.Join(actionArgs, ", "), ""))

		if access == "public" {
			routes = append(routes, templates.Route(name, verb, action, routeArgs, routeCall))
		}
	}
	return strings.Join(actions, "\n"), strings.Join(routes, "\n\n\t\t"), len(actions) > 0
}

func aIsVerb(a string) bool {
	if a == "get" {
		return true
	}
	if a == "post" {
		return true
	}
	if a == "put" {
		return true
	}
	if a == "delete" {
		return true
	}
	if a == "patch" {
		return true
	}
	return false
}
