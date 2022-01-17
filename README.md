# golang-fiber-boilerplate

Golang fiber boilerplate

This the directory structure which I follow in my Golang Fiber project.

Directory Explanation in relation to MVC

Controller: All business logic should go inside this directory. You can add as many as files as required.

Model: All data model should be in this directory. The files in this directory represent the tables in database.

Database: File inside this directory returns the pointer to the database connection.
Routes: Bootstraping the routes in the application.

.env file holds sensitive variables like auth secret, DSN for database connection.

run go mod tidy command and it will download all dependencies.
