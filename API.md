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

# Announcement
Announcement related api

## Announcement List [GET /announcements]
Get announcement list

+ Parameters
    + page: (string, optional) - The page of results to view.
        + Default: 1
    + per_page: (string, optional) - The amount of results per page.
        + Default: 15

+ Response 200 (application/json)
    + Body

            {
                "data": [
                    {
                        "id": 2,
                        "user_id": 1,
                        "title": "admin",
                        "content": "admin@admin.com",
                        "created_at": "2019-08-18 16:33:31",
                        "updated_at": "2019-08-18 16:33:31",
                        "deleted_at": null
                    }
                ],
                "meta": {
                    "pagination": {
                        "total": 1,
                        "count": 1,
                        "per_page": 15,
                        "current_page": 1,
                        "total_pages": 1
                    }
                }
            }

## Create Announcement [POST /announcements]
Create an announcement

+ Request (application/json)
    + Body

            {
                "title": "title",
                "content": "content"
            }

+ Response 200 (application/json)
    + Body

            {
                "data": {
                    "id": 1,
                    "title": "",
                    "content": ""
                }
            }

## Retrieve Announcement [GET /announcements/{id}]
Get an announcement

+ Response 200 (application/json)
    + Body

            {
                "data": {
                    "id": 1,
                    "title": "",
                    "content": ""
                }
            }

## Update Announcement [PUT /announcements/{id}]
Update an announcement

+ Request (application/json)
    + Body

            {
                "title": "title",
                "content": "content"
            }

+ Response 200 (application/json)
    + Body

            {
                "data": {
                    "id": 1,
                    "title": "",
                    "content": ""
                }
            }

## Delete Announcement [DELETE /announcements/{id}]
Delete announcement

+ Response 200 (application/json)