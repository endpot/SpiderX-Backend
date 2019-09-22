<?php

namespace App\Http\Controllers\Api\V1;

use App\Http\Controllers\Api\Controller;
use App\Http\Requests\Api\V1\Auth\ForgotPasswordRequest;
use App\Http\Requests\Api\V1\Auth\LoginRequest;
use App\Http\Requests\Api\V1\Auth\RegisterRequest;
use App\Http\Requests\Api\V1\Auth\ResetPasswordRequest;
use App\Models\User;
use Auth;
use Illuminate\Support\Facades\Password;
use Illuminate\Support\Str;
use Symfony\Component\HttpKernel\Exception\AccessDeniedHttpException;
use Symfony\Component\HttpKernel\Exception\HttpException;
use Symfony\Component\HttpKernel\Exception\NotFoundHttpException;
use Tymon\JWTAuth\Exceptions\JWTException;
use Tymon\JWTAuth\JWTAuth;

/**
 * Authentication related api
 *
 * @package App\Http\Controllers\Api\V1
 * @Resource("Auth", uri="/auth")
 */
class AuthController extends Controller
{
    public function __construct()
    {
        $this->middleware('jwt.auth', ['only' => ['me', 'logout']]);
    }

    /**
     * Register
     *
     * Register one user
     *
     * @param RegisterRequest $request
     * @param JWTAuth $JWTAuth
     * @return \Dingo\Api\Http\Response
     *
     * @Get("/register")
     * @Versions({"v1"})
     * @Transaction({
     *      @Request({"name": "HunterX", "email": "endpot@gmail.com", "password": "password"}),
     *      @Response(201)
     * })
     */
    public function register(RegisterRequest $request, JWTAuth $JWTAuth)
    {
        $user = new User();
        $user->name = $request->get('name', '');
        $user->email = $request->get('email', '');
        $user->password = $request->get('password', '');
        $user->passkey = md5($user->name . $user->email . Str::random(32) . time());

        if (!$user->save()) {
            throw new HttpException(500);
        }

        return $this->response->noContent()->setStatusCode(201);
    }

    /**
     * Login
     *
     * Login using name (or email) and password
     *
     * @param LoginRequest $request
     * @param JWTAuth $JWTAuth
     * @return \Illuminate\Http\JsonResponse
     *
     * @Get("/login")
     * @Versions({"v1"})
     * @Transaction({
     *      @Request({"login_name": "HunterX", "password": "password"}),
     *      @Response(200, body={"data": {"token": "token", "expires_in": 3600}}),
     *      @Response(422, body={"error": {"login_name": {"Invalid login_name"}}})
     * })
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

    /**
     * Me
     *
     * Get the authenticated user
     *
     * @return \Illuminate\Http\JsonResponse
     *
     * @Get("/me")
     * @Versions({"v1"})
     * @Transaction({
     *      @Response(200, body={
     *          "data": {
     *              "id": 1,
     *              "name": "HunterX",
     *              "email": "endpot@gmail.com",
     *              "title": "Admin",
     *              "avatar": "https://i.endpot.com/image/avatar/avatar.jpg",
     *              "gender": "M",
     *              "roles": {"user", "admin"},
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

    /**
     * Logout
     *
     * Log the user out (Invalidate the token)
     *
     * @return \Dingo\Api\Http\Response
     *
     * @Post("/logout")
     * @Versions({"v1"})
     * @Response(200)
     */
    public function logout()
    {
        Auth::guard()->logout();

        return $this->response->noContent()->setStatusCode(200);
    }

    /**
     * Refresh Token
     *
     * Refresh a token
     *
     * @return \Illuminate\Http\JsonResponse
     *
     * @Post("/refresh")
     * @Versions({"v1"})
     * @Transaction({
     *      @Request(headers={"Authorization": "Bearer bearTokenExample"}),
     *      @Response(200, body={"data": {"token": "token", "expires_in": 3600}})
     * })
     */
    public function refresh()
    {
        $token = Auth::guard()->refresh();

        return $this->response->array([
            'data' => [
                'token' => $token,
                'expires_in' => Auth::guard()->factory()->getTTL() * 60
            ]
        ]);
    }

    /**
     * Forget Password
     *
     * Send reset password email
     *
     * @param ForgotPasswordRequest $request
     * @return \Dingo\Api\Http\Response
     *
     * @Post("/recovery")
     * @Versions({"v1"})
     * @Transaction({
     *      @Request({"email": "endpot@gmail.com"}),
     *      @Response(200),
     *      @Response(500)
     * })
     */
    public function sendResetEmail(ForgotPasswordRequest $request)
    {
        $user = User::where('email', '=', $request->get('email'))->first();

        if (!$user) {
            throw new NotFoundHttpException();
        }

        $broker = $this->getPasswordBroker();
        $sendingResponse = $broker->sendResetLink($request->only('email'));

        if ($sendingResponse !== Password::RESET_LINK_SENT) {
            throw new HttpException(500);
        }

        return $this->response->noContent()->setStatusCode(200);
    }

    /**
     * Reset Password
     *
     * Reset password
     *
     * @param ResetPasswordRequest $request
     * @param JWTAuth $JWTAuth
     * @return \Dingo\Api\Http\Response
     *
     * @Post("/reset")
     * @Versions({"v1"})
     * @Transaction({
     *      @Request({
     *          "email": "endpot@gmail.com",
     *          "password": "password",
     *          "password_confirmation": "password",
     *          "token": "token"
     *      }),
     *      @Response(200, body={"data": {"token": "token", "expires_in": 3600}}),
     *      @Response(500)
     * })
     */
    public function resetPassword(ResetPasswordRequest $request, JWTAuth $JWTAuth)
    {
        $response = $this->getPasswordBroker()->reset(
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

    /**
     * Get the broker to be used during password reset.
     *
     * @return \Illuminate\Contracts\Auth\PasswordBroker
     */
    private function getPasswordBroker()
    {
        return Password::broker();
    }
}