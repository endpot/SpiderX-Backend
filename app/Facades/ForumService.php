<?php

namespace App\Facades;

use Illuminate\Support\Facades\Facade;

class ForumService extends Facade
{
    protected static function getFacadeAccessor()
    {
        return 'ForumService';
    }
}