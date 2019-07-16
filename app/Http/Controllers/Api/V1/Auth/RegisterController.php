<?php

namespace App\Http\Controllers\Api\V1\Auth;

use App\Models\User;
use Tymon\JWTAuth\JWTAuth;
use App\Http\Controllers\Api\Controller;
use App\Http\Requests\Api\V1\Auth\RegisterRequest;
use Symfony\Component\HttpKernel\Exception\HttpException;

class RegisterController extends Controller
{
    public function __construct()
    {
        //
    }

    public function register(RegisterRequest $request, JWTAuth $JWTAuth)
    {
        $user = new User();
        $user->name = $request->name;
        $user->email = $request->email;
        $user->password = $request->password;

        if (!$user->save()) {
            throw new HttpException(500);
        }

        return $this->response->noContent()->setStatusCode(201);
    }
}
