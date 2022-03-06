# go-kdrama-crud
A very simple CRUD API for managing a "database" of K-Dramas.

This API allows you to create, read, update, or delete a list of K-dramas (Korean Dramas).

Each K-drama struct has an ID, ISBN, Title, Writer, and Director.

Everything is contained within one file (main.go). As the goal of this project was just to get familiar with creating and serving endpoints, 
all of the Drama structs are simply kept in a slice, instead of a database.

