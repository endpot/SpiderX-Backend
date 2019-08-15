<?php

use Illuminate\Database\Seeder;

class UserTableSeeder extends Seeder
{
    /**
     * Run the database seeds.
     *
     * @return void
     */
    public function run()
    {
        $adminUser = new \App\Models\User();
        $adminUser->name = 'admin';
        $adminUser->email = 'admin@admin.com';
        $adminUser->password = '123456';
        $adminUser->real_name = '管理员';
        $adminUser->avatar = 'https://i.endpot.com/image/WPA5T/2014100110290460-easyicon-net-128.jpg';
        $adminUser->save();

        factory(App\Models\User::class, 50)->create();
    }
}