<?php

namespace App\Filters;

use Illuminate\Support\Arr;

class QueryFilter implements QueryFilterInterface
{
    protected $filterParams = [];

    public function __construct($params = [])
    {
        $this->filterParams = (array)$params;
    }

    public function get($key, $default = null)
    {
        if (Arr::has($this->filterParams, $key)) {
            return Arr::get($this->filterParams, $key);
        }

        return $default;
    }


    public function set($key, $value)
    {
        if (Arr::set($this->filterParams, $key, $value)) {
            return true;
        }

        return false;
    }
}