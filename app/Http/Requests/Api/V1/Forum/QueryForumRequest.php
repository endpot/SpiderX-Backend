<?php

namespace App\Http\Requests\Api\V1\Forum;

use Dingo\Api\Http\FormRequest;

class QueryForumRequest extends FormRequest
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