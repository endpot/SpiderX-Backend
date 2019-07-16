<?php

namespace App\Http\Requests\Api\V1\Auth;

use Dingo\Api\Http\FormRequest;

class ForgotPasswordRequest extends FormRequest
{
    public function rules()
    {
        return [
            'email' => 'required|email'
        ];
    }

    public function authorize()
    {
        return true;
    }
}
