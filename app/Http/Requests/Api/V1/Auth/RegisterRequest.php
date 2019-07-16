<?php

namespace App\Http\Requests\Api\V1\Auth;

use Dingo\Api\Http\FormRequest;

class RegisterRequest extends FormRequest
{
    public function rules()
    {
        return [
            'name' => 'required|unique:users',
            'email' => 'required|email|unique:users',
            'password' => 'required'
        ];
    }

    public function authorize()
    {
        return true;
    }
}
