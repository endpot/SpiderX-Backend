<?php

namespace App\Providers;

use App\Services\AnnouncementService;
use App\Services\ChatService;
use App\Services\ForumService;
use Illuminate\Support\ServiceProvider;

class DomainServiceProvider extends ServiceProvider
{
    /**
     * Register services.
     *
     * @return void
     */
    public function register()
    {
        //
    }

    /**
     * Bootstrap services.
     *
     * @return void
     */
    public function boot()
    {
        $this->app->bind('AnnouncementService', function () {
            return new AnnouncementService();
        });

        $this->app->bind('ChatService', function () {
            return new ChatService();
        });

        $this->app->bind('ForumService', function () {
            return new ForumService();
        });
    }
}
