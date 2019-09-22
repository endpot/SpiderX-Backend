<?php

namespace App\Http\Requests\Api\V1\Topic;

use Dingo\Api\Http\FormRequest;

class QueryPostRequest extends FormRequest
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