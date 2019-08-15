<?php

namespace App\Results;

class CommonSuccessResult extends ResultAbstract
{
    public function __construct($data = null, $message = '')
    {
        $this->setStatus(true);
        $this->setData($data);
        $this->setMessage($message);
    }
}
