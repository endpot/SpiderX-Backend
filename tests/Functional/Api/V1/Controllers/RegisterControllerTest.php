<?php

namespace App\Functional\Api\V1\Controllers;

use Config;
use App\TestCase;
use Illuminate\Foundation\Testing\DatabaseMigrations;

class RegisterControllerTest extends TestCase
{
    use DatabaseMigrations;

    public function testRegisterSuccessfully()
    {
        $this->post('api/auth/register', [
            'name' => 'Test User',
            'email' => 'test@email.com',
            'password' => '123456'
        ])->assertJson([
            'status' => 'ok'
        ])->assertStatus(201);
    }

    public function testRegisterSuccessfullyWithTokenRelease()
    {
        Config::set('boilerplate.register.release_token', true);

        $this->post('api/auth/register', [
            'name' => 'Test User',
            'email' => 'test@email.com',
            'password' => '123456'
        ])->assertJsonStructure([
            'status', 'token'
        ])->assertJson([
            'status' => 'ok'
        ])->assertStatus(201);
    }

    public function testRegisterReturnsValidationError()
    {
        $this->post('api/auth/register', [
            'name' => 'Test User',
            'email' => 'test@email.com'
        ])->assertJsonStructure([
            'error'
        ])->assertStatus(422);
    }
}
