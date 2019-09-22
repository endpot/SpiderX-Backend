<?php

namespace App\Http\Requests\Api\V1\Topic;

use Dingo\Api\Http\FormRequest;

class CreatePostRequest extends FormRequest
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