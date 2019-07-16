<?php

namespace App\Http\Controllers\Api\V1\Auth;

use Tymon\JWTAuth\JWTAuth;
use App\Http\Controllers\Api\Controller;
use Illuminate\Support\Facades\Password;
use App\Http\Requests\Api\V1\Auth\ResetPasswordRequest;
use Symfony\Component\HttpKernel\Exception\HttpException;

class ResetPasswordController extends Controller
{
    public function resetPassword(ResetPasswordRequest $request, JWTAuth $JWTAuth)
    {
        $response = $this->broker()->reset(
            $this->credentials($request),
            function ($user, $password) {
                $this->reset($user, $password);
            }
        );

        if ($response !== Password::PASSWORD_RESET) {
            throw new HttpException(500);
        }

        return $this->response->noContent()->setStatusCode(200);
    }

    /**
     * Get the broker to be used during password reset.
     *
     * @return \Illuminate\Contracts\Auth\PasswordBroker
     */
    public function broker()
    {
        return Password::broker();
    }

    /**
     * Get the password reset credentials from the request.
     *
     * @param  ResetPasswordRequest $request
     * @return array
     */
    protected function credentials(ResetPasswordRequest $request)
    {
        return $request->only(
            'email',
            'password',
            'password_confirmation',
            'token'
        );
    }

    /**
     * Reset the given user's password.
     *
     * @param  \Illuminate\Contracts\Auth\CanResetPassword $user
     * @param  string $password
     * @return void
     */
    protected function reset($user, $password)
    {
        $user->password = $password;
        $user->save();
    }
}
