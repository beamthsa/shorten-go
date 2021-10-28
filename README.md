# Gin Boilerplate

Gin Boilerplate is a back-end template for standard building
work on OPN environment local, staging, DevOps and adaptable web service

This project is base on [Gin Web Framework](https://github.com/gin-gonic/gin).
It does not impose a specific development
philosophy or framework, so you're free to architect your code in the
way that you want.

## Suggests

- using unix timestamp instead string datetime format because
  - prevent parsing format error
  - code is easy to adaptable
- set datetime at layer docker and layer code

## Guidelines

- [Git commit message format](https://www.conventionalcommits.org/en/v1.0.0/)
- [Database](https://github.com/pthom/northwind_psql)
- every function should have return
- if don't have return, should be `singleton`


## Directory Structure
- if you see `init.go` any directory, please read first.
- directory `config` contains all of your application's configuration files, such as `env` `database` `log`
- directory `app`
  - directory `model` contains all the database query
  - directory `controller` contains `filter` and `validate` the http `(request/response)` from gin.Context before send to the `model`
- directory `resource` contains `docker` and `DevOps`
