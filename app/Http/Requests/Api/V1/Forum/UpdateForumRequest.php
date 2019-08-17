<?php

namespace App\Http\Requests\Api\V1\Forum;

use Dingo\Api\Http\FormRequest;

class UpdateForumRequest extends FormRequest
{
    public function rules()
    {
        return [
            'name' => 'required',
        ];
    }

    public function authorize()
    {
        return true;
    }
}