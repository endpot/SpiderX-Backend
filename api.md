FORMAT: 1A

# SpiderX-Backend API

# Auth [/auth]
Authentication related api

## Register [GET /auth/register]
Register one user

+ Request (application/json)
    + Body

            {
                "name": "HunterX",
                "email": "endpot@gmail.com",
                "password": "password"
            }

+ Response 201 (application/json)

## Login [GET /auth/login]
Login using name (or email) and password

+ Request (application/json)
    + Body

            {
                "login_name": "HunterX",
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
                    "login_name": [
                        "Invalid login_name"
                    ]
                }
            }

## Me [GET /auth/me]
Get the authenticated user

+ Response 200 (application/json)
    + Body

            {
                "data": {
                    "id": 1,
                    "name": "HunterX",
                    "email": "endpot@gmail.com",
                    "title": "Admin",
                    "avatar": "https://i.endpot.com/image/avatar/avatar.jpg",
                    "gender": "M",
                    "roles": [
                        "user",
                        "admin"
                    ]
                }
            }

## Logout [POST /auth/logout]
Log the user out (Invalidate the token)

+ Response 200 (application/json)

## Refresh Token [POST /auth/refresh]
Refresh a token

+ Request (application/json)
    + Headers

            Authorization: Bearer bearTokenExample

+ Response 200 (application/json)
    + Body

            {
                "data": {
                    "token": "token",
                    "expires_in": 3600
                }
            }

## Forget Password [POST /auth/recovery]
Send reset password email

+ Request (application/json)
    + Body

            {
                "email": "endpot@gmail.com"
            }

+ Response 200 (application/json)

+ Response 500 (application/json)

## Reset Password [POST /auth/reset]
Reset password

+ Request (application/json)
    + Body

            {
                "email": "endpot@gmail.com",
                "password": "password",
                "password_confirmation": "password",
                "token": "token"
            }

+ Response 200 (application/json)
    + Body

            {
                "data": {
                    "token": "token",
                    "expires_in": 3600
                }
            }

+ Response 500 (application/json)