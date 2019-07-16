<?php

namespace App\Http\Requests\Api\V1\Auth;

use Dingo\Api\Http\FormRequest;

class ResetPasswordRequest extends FormRequest
{
    public function rules()
    {
        return [
            'token' => 'required',
            'email' => 'required|email',
            'password' => 'required|confirmed'
        ];
    }

    public function authorize()
    {
        return true;
    }
}
