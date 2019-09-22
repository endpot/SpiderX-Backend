<?php

namespace App\Carriers;

interface DataCarrierInterface
{
    public function get($key, $default = null);

    public function set($key, $value);

    public function all();
}
