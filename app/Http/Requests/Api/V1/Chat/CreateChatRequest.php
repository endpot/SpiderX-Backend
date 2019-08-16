<?php

namespace App\Http\Requests\Api\V1\Chat;

use Dingo\Api\Http\FormRequest;

class CreateChatRequest extends FormRequest
{
    public function rules()
    {
        return [
            'content' => 'required',
        ];
    }

    public function authorize()
    {
        return true;
    }
}