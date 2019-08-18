<?php

namespace App\Http\Requests\Api\V1\Chat;

use Dingo\Api\Http\FormRequest;

class QueryChatRequest extends FormRequest
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