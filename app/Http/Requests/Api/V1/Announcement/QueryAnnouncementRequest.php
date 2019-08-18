<?php

namespace App\Http\Requests\Api\V1\Announcement;

use Dingo\Api\Http\FormRequest;

class QueryAnnouncementRequest extends FormRequest
{
    public function rules()
    {
        return [];
    }

    public function authorize()
    {
        return true;
    }
}