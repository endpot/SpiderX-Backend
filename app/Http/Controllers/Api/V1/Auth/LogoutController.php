<?php

namespace App\Http\Controllers\Api\V1\Auth;

use App\Http\Controllers\Api\Controller;
use Auth;

class LogoutController extends Controller
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
     * Log the user out (Invalidate the token)
     *
     * @return \Dingo\Api\Http\Response
     */
    public function logout()
    {
        Auth::guard()->logout();

        return $this->response->noContent()->setStatusCode(200);
    }
}
