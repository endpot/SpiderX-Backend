<?php

namespace App\Facades;

use Illuminate\Support\Facades\Facade;

class TopicService extends Facade
{
    protected static function getFacadeAccessor()
    {
        return 'TopicService';
    }
}