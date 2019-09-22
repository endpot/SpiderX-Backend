<?php

namespace App\Http\Controllers\Api\V1\Auth;

use App\Http\Controllers\Api\Controller;
use Auth;

/**
 * User Controller
 *
 * @Resource("", uri="")
 */
class UserController extends Controller
{
    /**
     * Create a new AuthController instance.
     *
     * @return void
     */
    public function __construct()
    {
        $this->middleware('jwt.auth', []);
    }

    /**
     * Get the authenticated User
     *
     * Get the authenticated user
     *
     * @return \Illuminate\Http\JsonResponse
     *
     * @Get("/auth/me")
     * @Versions({"v1"})
     * @Transaction({
     *      @Response(200, body={
     *          "data": {
     *              "id": "id",
     *              "name": "name",
     *              "email": "email",
     *              "title": "title",
     *              "avatar": "avatar",
     *              "gender": "gender",
     *              "roles": "array",
     *          }
     *     }),
     * })
     */
    public function me()
    {
        $user = Auth::guard()->user();
        return $this->response->array([
            'data' => [
                'id' => $user->id,
                'name' => $user->name,
                'email' => $user->email,
                'title' => $user->title,
                'avatar' => $user->avatar,
                'gender' => $user->gender,
                'roles' => $user->getRoleNames(),
            ]
        ]);
    }
}
