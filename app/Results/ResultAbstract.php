<?php

namespace App\Results;

class ResultAbstract implements ResultInterface
{
    protected $data;

    protected $status;

    protected $message;

    public function setData($data)
    {
        $this->data = $data;
    }

    public function setStatus($status)
    {
        $this->status = $status;
    }

    public function setMessage($message)
    {
        $this->message = $message;
    }

    public function getData()
    {
        return $this->data;
    }

    public function getStatus()
    {
        return $this->status;
    }

    public function getMessage()
    {
        return $this->message;
    }
}
