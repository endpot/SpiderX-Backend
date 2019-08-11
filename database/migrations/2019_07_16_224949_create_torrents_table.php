<?php

use Illuminate\Support\Facades\Schema;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Database\Migrations\Migration;

class CreateTorrentsTable extends Migration
{
    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        Schema::create('torrents', function (Blueprint $table) {
            $table->increments('id');
            $table->binary('info_hash');
            $table->string('title');
            $table->string('simple_desc');
            $table->text('desc');
            $table->unsignedTinyInteger('category_id')->default(0);

            $table->unsignedBigInteger('size')->default(0);
            $table->unsignedMediumInteger('comment_count')->default(0);
            $table->unsignedInteger('view_count')->default(0);
            $table->unsignedInteger('seeder_count')->default(0);
            $table->unsignedInteger('leecher_count')->default(0);
            $table->unsignedInteger('snatcher_count')->default(0);

            $table->binary('nfo');
            $table->unsignedInteger('reward_bonus')->default(0)->comment('');
            $table->unsignedTinyInteger('is_anonymous')->default(0);
            $table->unsignedTinyInteger('position_level')->default(0);
            $table->unsignedInteger('uploader_id')->default(0)->comment('发布者ID');
            $table->unsignedInteger('owner_id')->default(0)->comment('所有者ID');
            $table->unsignedTinyInteger('is_inactive')->default(0)->comment('是否断种');

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
        Schema::dropIfExists('torrents');
    }
}
