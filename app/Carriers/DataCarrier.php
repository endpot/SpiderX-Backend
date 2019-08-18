<?php

namespace App\Carriers;

use Illuminate\Support\Arr;

class DataCarrier implements DataCarrierInterface
{
    protected $data = [];

    public function __construct($data = [])
    {
        $this->data = (array)$data;
    }

    public function get($key, $default = null)
    {
        if (Arr::has($this->data, $key)) {
            return Arr::get($this->data, $key);
        }

        return $default;
    }


    public function set($key, $value)
    {
        if (Arr::set($this->data, $key, $value)) {
            return true;
        }

        return false;
    }
}