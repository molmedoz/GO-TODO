# Introduction
During the previous hackerrank exercise, you built a simple todo list application in javascript. Currently, the data is not persisted and is lost as soon as you close the page.
For this exercise, you will build a backend for that todo list application in Go (Golang)



# Instructions

You will build an HTTP server that will expose RESTful json APIs to:

* Add an item to the todo list
* Update the status of an item (to do or completed)
* Remove an item from the list
* The server will persist the data in a database of your choice.
* Your code should follow the Go language conventions
* As much as possible, you will avoid using external dependencies (do not use frameworks like gobuffalo), the Go standard library should be sufficient for the exercise, except for some database drivers.
* If you need external dependencies, you will use dep https://github.com/golang/dep for dependency management

## How to submit your code

The preferred solution to submit your code would be to create a public github repository. If this is not possible, you can provide your source code in zip file.
Make sure to include all documentation needed to run your code and if applicable, the database schema.

## Bonus question

Implement user management (with authentication) to allow each user to have their own todo list.
