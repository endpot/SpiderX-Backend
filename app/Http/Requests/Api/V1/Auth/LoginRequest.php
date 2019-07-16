<?php

namespace App\Http\Requests\Api\V1\Auth;

use Dingo\Api\Http\FormRequest;

class LoginRequest extends FormRequest
{
    public function rules()
    {
        return [
            'login_name' => 'required',
            'password' => 'required'
        ];
    }

    public function authorize()
    {
        return true;
    }
}
