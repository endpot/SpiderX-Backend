<?php

namespace App\Results;

class CommonErrorResult extends ResultAbstract
{
    public function __construct($message = '')
    {
        $this->setStatus(false);
        $this->setMessage($message);
    }
}
