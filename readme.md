##install library
<ul>
<li>go get -u github.com/gofiber/fiber/v2</li>
<li>go get -u gorm.io/gorm</li>
<li>go get -u gorm.io/driver/postgres</li>
<li>go get -u github.com/asaskevich/govalidator</li>
<li>go get -u github.com/dgrijalva/jwt-go</li>
</ul>
<hr>

##Build docker DB
`make buildDB`
<hr>

##Migrate database
`make migreate`
<hr>

##Run server
`make dev`

<hr>

## Auth API: sign up
    url: /api/auth/signup
###Header: 
    Content-Type: application/json 
###Method:
    POST
###Body:
    `{
        "name": "",
        "lastname": "",
        "email": "",
        "password": "",
        "confirmPassword": "",
        "salary": floatType,
        "phone": ""
    }`
<hr>

## Auth API: sign in
    url: /api/auth/signin
###Header:
    Content-Type: application/json 
###Method:
    POST
###Body:
    {
        "email": "",
        "password": ""
    }
<hr>

## Auth API: sign out
    url: /api/auth/signout
###Header:
    Content-Type: application/json 
###Method:
    POST
###Body:
    {}

<hr>

## User API: show self profile
    url: /api/user
###Header:
    Content-Type: application/json 
###Method:
    GET
###Body:
    {}

<hr>

## User API: find user
    url: /api/user/{userID}
###Header:
    Content-Type: application/json 
###Method:
    GET
###Body:
    {}

<hr>

## User API: show user list
    url: /api/user/list
###Header:
    Content-Type: application/json 
###Method:
    GET
###Body:
    {}

<hr>

## User API: update user
    url: /api/user/update
###Header:
    Content-Type: application/json 
###Method:
    POST
###Body:
    {
        "name": "",
        "lastname": ""
    }

<hr>

## Admin API: create new department
    url: /api/admin/department
###Header:
    Content-Type: application/json 
###Method:
    POST
###Body:
    {
        "title": ""
    }

<hr>

## Admin API: change user department
    url: /api/admin/update_user_department
###Header:
    Content-Type: application/json 
###Method:
    POST
###Body:
    {
        "userID": intType,
        "departmentID": intType
    }
