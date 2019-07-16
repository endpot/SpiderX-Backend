<?php

use Dingo\Api\Routing\Router;

/** @var Router $api */
$api = app(Router::class);

$api->version('v1', ['namespace' => 'App\Http\Controllers\Api\V1'], function (Router $api) {
    $api->group(['prefix' => 'auth', 'namespace' => 'Auth'], function (Router $api) {
        $api->post('register', 'RegisterController@register');
        $api->post('login', 'LoginController@login');

        $api->post('recovery', 'ForgotPasswordController@sendResetEmail');
        $api->post('reset', 'ResetPasswordController@resetPassword');

        $api->post('logout', 'LogoutController@logout');
        $api->post('refresh', 'RefreshController@refresh');
        $api->get('me', 'UserController@me');
    });

    $api->group(['prefix' => 'actions'], function (Router $api) {
        // 健康检测
        $api->get('ping', function () {
            return response()->noContent();
        });
    });
});
