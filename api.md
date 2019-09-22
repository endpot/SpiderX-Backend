FORMAT: 1A

# SpiderX-Backend API

# 
Register Controller

## Register User [GET /auth/register]
Register one user

+ Request (application/json)
    + Body

            {
                "name": "name",
                "email": "email",
                "password": "password"
            }

+ Response 201 (application/json)

# 
Login Controller

## Log the user in [GET /auth/login]
Log the user in

+ Request (application/json)
    + Body

            {
                "name": "name",
                "password": "password"
            }

+ Response 200 (application/json)
    + Body

            {
                "data": {
                    "token": "token",
                    "expires_in": 3600
                }
            }

+ Response 422 (application/json)
    + Body

            {
                "error": {
                    "username": [
                        "Username is already taken."
                    ]
                }
            }

# 
User Controller

## Get the authenticated User [GET /auth/me]
Get the authenticated user

+ Response 200 (application/json)
    + Body

            {
                "data": {
                    "id": "id",
                    "name": "name",
                    "email": "email",
                    "title": "title",
                    "avatar": "avatar",
                    "gender": "gender",
                    "roles": "array"
                }
            }