<?php

use Illuminate\Support\Facades\Schema;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Database\Migrations\Migration;

class CreateUsersTable extends Migration
{
    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        Schema::create('users', function (Blueprint $table) {
            $table->increments('id');
            $table->string('name')->unique()->comment('用户名');
            $table->string('email')->unique()->comment('注册邮箱');
            $table->string('password')->comment('密码');
            $table->string('passkey', '32')->default('')->comment('Tracker 通信密钥');
            $table->tinyInteger('state')->default(0)->comment('账号状态');

            $table->timestamp('last_login_date')->comment('最近登录时间');
            $table->timestamp('last_access_date')->comment('最近访问时间');
            $table->ipAddress('last_access_ip')->comment('最近访问IP');

            $table->unsignedTinyInteger('class')->default(1)->comment('用户等级');
            $table->unsignedTinyInteger('max_class_once')->default(1)->comment('曾经最高等级');

            $table->string('title')->comment('头衔');
            $table->string('avatar')->default('')->comment('头像');
            $table->text('info')->comment('个人说明');
            $table->unsignedTinyInteger('gender')->default(0)->comment('性别');
            $table->unsignedInteger('country_id')->default(0)->comment('所属国家');
            $table->unsignedInteger('school_id')->default(0)->comment('所属学校');

            $table->unsignedInteger('inviter')->default(0)->comment('邀请人');
            $table->unsignedTinyInteger('invite_count')->default(0)->comment('邀请数量');

            $table->unsignedBigInteger('up_traffic')->default(0)->comment('上传量');
            $table->unsignedBigInteger('down_traffic')->default(0)->comment('下载量');
            $table->unsignedBigInteger('seed_time')->default(0)->comment('做种时间');
            $table->unsignedBigInteger('leech_time')->default(0)->comment('下载时间');
            $table->decimal('seed_bonus', 10, 1)->default(0)->comment('做种积分');

            $table->unsignedTinyInteger('isp')->default(0)->comment('运营商');
            $table->unsignedTinyInteger('up_bandwidth')->default(0)->comment('上传带宽');
            $table->unsignedTinyInteger('down_bandwidth')->default(0)->comment('下载带宽');

            $table->unsignedTinyInteger('is_vip')->default(0)->comment('是否 VIP');
            $table->timestamp('vip_added_at')->comment('VIP 添加日期');
            $table->timestamp('vip_until')->comment('VIP 有效期');

            $table->unsignedTinyInteger('is_donor')->default(0)->comment('是否捐赠用户');
            $table->decimal('total_donation')->default(0)->comment('捐赠金额');

            $table->unsignedTinyInteger('is_warned')->default(0)->comment('是否 VIP');
            $table->unsignedTinyInteger('warned_type')->default(0)->comment('警告类型');
            $table->timestamp('warned_at')->comment('警告日期');
            $table->timestamp('warned_by')->comment('警告人');
            $table->timestamp('warned_until')->comment('警告截止日期');

            $table->timestamps();
            $table->softDeletes();
        });
    }

    /**
     * Reverse the migrations.
     *
     * @return void
     */
    public function down()
    {
        Schema::dropIfExists('users');
    }
}
