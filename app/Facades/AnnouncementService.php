<?php

namespace App\Facades;

use Illuminate\Support\Facades\Facade;

class AnnouncementService extends Facade
{
    protected static function getFacadeAccessor()
    {
        return 'AnnouncementService';
    }
}