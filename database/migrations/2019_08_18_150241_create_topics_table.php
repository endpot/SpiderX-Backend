<?php

use Illuminate\Support\Facades\Schema;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Database\Migrations\Migration;

class CreateTopicsTable extends Migration
{
    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        Schema::create('topics', function (Blueprint $table) {
            $table->increments('id');
            $table->unsignedInteger('user_id');
            $table->unsignedInteger('forum_id');
            $table->string('subject')->comment('主题标题');
            $table->text('content')->comment('主题内容');
            $table->unsignedInteger('view_count')->default(0)->comment('查看次数');
            $table->unsignedInteger('position')->default(0)->comment('置顶位置');
            $table->boolean('is_hot')->default(false)->comment('是否热门');
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
        Schema::dropIfExists('topics');
    }
}
