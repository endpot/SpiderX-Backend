<?php

namespace App\Http\Controllers\Api\V1\Auth;

use Symfony\Component\HttpKernel\Exception\HttpException;
use Tymon\JWTAuth\JWTAuth;
use App\Http\Controllers\Api\Controller;
use App\Http\Requests\Api\V1\Auth\LoginRequest;
use Tymon\JWTAuth\Exceptions\JWTException;
use Symfony\Component\HttpKernel\Exception\AccessDeniedHttpException;
use Auth;

class LoginController extends Controller
{
    /**
     * Log the user in
     *
     * @param LoginRequest $request
     * @param JWTAuth $JWTAuth
     * @return \Illuminate\Http\JsonResponse
     */
    public function login(LoginRequest $request, JWTAuth $JWTAuth)
    {
        $credentials = $request->only(['login_name', 'password']);
        try {
            $token = Auth::guard()->attempt([
                'name' => $credentials['login_name'],
                'password' => $credentials['password'],
            ]);

            if (!$token) {
                $token = Auth::guard()->attempt([
                    'email' => $credentials['login_name'],
                    'password' => $credentials['password'],
                ]);
            }
            if (!$token) {
                throw new AccessDeniedHttpException();
            }

        } catch (JWTException $e) {
            throw new HttpException(500);
        }

        return $this->response->array([
            'data' => [
                'token' => $token,
                'expires_in' => Auth::guard()->factory()->getTTL() * 60
            ]
        ])->setStatusCode(200);
    }
}
