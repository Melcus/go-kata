package user

import (
	"eon/kata/mike/pkg/kernel"
	"net/http"
)

func loadRoutes(app *kernel.Application, provider *provider) {
	app.Router.HandleFunc("GET /users", provider.handleGetUsers(app))
	app.Router.HandleFunc("POST /users", provider.handleCreateUser(app))
	app.Router.HandleFunc("GET /users/{email}", provider.handleFindUser(app))
}

func (provider *provider) handleGetUsers(app *kernel.Application) http.HandlerFunc {
	manager := newUserManager(provider.repo)

	return func(writer http.ResponseWriter, request *http.Request) {
		publicUsers, err := manager.listUsers()

		if err != nil {
			app.Logger.Error("Error while loading users: %v", err)
			app.Respond(writer, request, nil, http.StatusBadRequest)
		}

		app.Respond(writer, request, publicUsers, http.StatusOK)
	}
}

func (provider *provider) handleCreateUser(app *kernel.Application) http.HandlerFunc {
	manager := newUserManager(provider.repo)

	return func(writer http.ResponseWriter, request *http.Request) {
		parseFormError := request.ParseForm()

		if parseFormError != nil {
			app.Logger.Error("Error while parsing form: %v", parseFormError)
			app.Respond(writer, request, nil, http.StatusBadRequest)
		}

		// TODO: validations here
		// valid & unique email
		// password validation

		user := newUserBuilder().
			setName(request.Form.Get("name")).
			setEmail(request.Form.Get("email")).
			setPassword(request.Form.Get("password")). // TODO: encrypt password
			build()

		publicUser, createUserErr := manager.createUser(*user)

		if createUserErr != nil {
			app.Logger.Error("Error while creating user: %v", createUserErr)
			app.Respond(writer, request, createUserErr.Error(), http.StatusBadRequest)
		}

		app.Respond(writer, request, publicUser, http.StatusOK)
	}
}

func (provider *provider) handleFindUser(app *kernel.Application) http.HandlerFunc {
	manager := newUserManager(provider.repo)

	return func(writer http.ResponseWriter, request *http.Request) {
		email := request.PathValue("email")

		if email == "" {
			app.Logger.Error("Email is required")
			app.Respond(writer, request, nil, http.StatusBadRequest)
		}

		publicUser, err := manager.findUserByEmail(email)
		if err != nil {
			app.Logger.Error("Error while loading user: %v", err)
			app.Respond(writer, request, nil, http.StatusNotFound)
		}

		app.Respond(writer, request, publicUser, http.StatusOK)
	}
}
