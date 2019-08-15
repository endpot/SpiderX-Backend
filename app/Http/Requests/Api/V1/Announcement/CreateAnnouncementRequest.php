<?php

namespace App\Http\Requests\Api\V1\Announcement;

use Dingo\Api\Http\FormRequest;

class CreateAnnouncementRequest extends FormRequest
{
    public function rules()
    {
        return [
            'title' => 'required',
            'content' => 'required',
        ];
    }

    public function authorize()
    {
        return true;
    }
}