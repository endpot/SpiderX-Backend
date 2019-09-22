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

    $api->group([], function (Router $api) {
        $api->get('announcements', 'AnnouncementController@getAnnouncementList');
        $api->post('announcements', 'AnnouncementController@createAnnouncement');
        $api->get('announcements/{id}', 'AnnouncementController@getAnnouncement');
        $api->put('announcements/{id}', 'AnnouncementController@updateAnnouncement');
        $api->delete('announcements/{id}', 'AnnouncementController@deleteAnnouncement');

        $api->get('chats', 'ChatController@getChatList');
        $api->post('chats', 'ChatController@createChat');
        $api->get('chats/{id}', 'ChatController@getChat');
        $api->delete('chats/{id}', 'ChatController@deleteChat');

        $api->get('forums', 'ForumController@getForumList');
        $api->post('forums', 'ForumController@createForum');
        $api->get('forums/{id}', 'ForumController@getForum');
        $api->put('forums/{id}', 'ForumController@updateForum');
        $api->delete('forums/{id}', 'ForumController@deleteForum');

        $api->get('topics', 'TopicController@getTopicList');
        $api->post('topics', 'TopicController@createTopic');
        $api->get('topics/{id}', 'TopicController@getTopic');
        $api->put('topics/{id}', 'TopicController@updateTopic');
        $api->delete('topics/{id}', 'TopicController@deleteTopic');

        $api->get('topics/{id}/posts', 'TopicController@getPostList');
        $api->post('topics/{id}/posts', 'TopicController@createPost');
    });

    $api->group(['prefix' => 'actions'], function (Router $api) {
        // 健康检测
        $api->get('ping', function () {
            return response()->noContent();
        });
    });
});
