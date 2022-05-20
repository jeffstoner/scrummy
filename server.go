package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func getRoot(c echo.Context) error {
	// this is not browsable site
	return c.String(http.StatusOK, "Nothing to see here")
}

func authenticateUser(c echo.Context) error {
	// insert logic for actually authenticating a user
	return c.String(http.StatusOK, "617c6694529c465ca9f224d12513db39")
}

func getBlogTitles(c echo.Context) error {
	// return a list of blog entries, possibly with pagination
	return c.String(http.StatusOK, "Normally, there would be a list of blog titles but not today.")
}

func getBlogEntry(c echo.Context) error {
	// return a single blog entry
	return c.String(http.StatusOK, "This is a blog entry.")
}

func saveBlog(c echo.Context) error {
	// insert logic to save a blog entry to a datastore
	return c.String(http.StatusCreated, "")
}

func main() {
	// create a new instance of Echo
	e := echo.New()

	// define router paths
	// our the root of the site
	e.GET("/", getRoot)

	// log in to site. returns a token for API use
	e.POST("/login", authenticateUser)

	// get all blog entries
	e.GET("/blog", getBlogTitles)

	// get a specific blog entry
	e.GET("/blog/:id", getBlogEntry)

	// create a new blog entry
	e.POST("/blog", saveBlog)

	// start serving requests
	e.Logger.Fatal(e.Start(":3000"))
}
