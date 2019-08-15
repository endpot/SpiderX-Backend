<?php

namespace App\Results;

interface ResultInterface
{
    public function setData($data);

    public function setStatus($status);

    public function setMessage($message);

    public function getData();

    public function getStatus();

    public function getMessage();
}
