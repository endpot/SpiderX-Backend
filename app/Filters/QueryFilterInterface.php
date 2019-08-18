<?php

namespace App\Filters;

interface QueryFilterInterface
{
    public function get($key, $default = null);

    public function set($key, $value);
}