<?php

/*
|--------------------------------------------------------------------------
| Model Factories
|--------------------------------------------------------------------------
|
| Here you may define all of your model factories. Model factories give
| you a convenient way to create models for testing and seeding your
| database. Just tell the factory how a default model should look.
|
*/

/** @var \Illuminate\Database\Eloquent\Factory $factory */

use Faker\Generator;
use App\Models\User;

// User
$factory->define(User::class, function (Generator $faker) {
    $password = '123456';

    return [
        'name' => $faker->unique()->userName,
        'email' => $faker->unique()->safeEmail,
        'password' => $password,
    ];
});
