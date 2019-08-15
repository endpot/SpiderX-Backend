<?php

namespace App\Providers;

use App\Services\AnnouncementService;
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
    }
}
