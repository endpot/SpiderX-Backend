<?php

namespace App\Http\Requests\Api\V1\Topic;

use Dingo\Api\Http\FormRequest;

class CreateTopicRequest extends FormRequest
{
    public function rules()
    {
        return [
            'forum_id' => 'required|numeric|exists:forums,id',
            'subject' => 'required|max:255',
            'content' => 'required',
        ];
    }

    public function authorize()
    {
        return true;
    }
}