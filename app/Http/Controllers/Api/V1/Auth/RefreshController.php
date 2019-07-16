<?php

namespace App\Http\Controllers\Api\V1\Auth;

use App\Http\Controllers\Api\Controller;
use Auth;

class RefreshController extends Controller
{
    /**
     * Refresh a token.
     *
     * @return \Illuminate\Http\JsonResponse
     */
    public function refresh()
    {
        $token = Auth::guard()->refresh();

        return $this->response->array([
            'data' => [
                'token' => $token,
                'expires_in' => Auth::guard()->factory()->getTTL() * 60
            ]
        ]);
    }
}
