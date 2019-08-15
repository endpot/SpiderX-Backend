<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\SoftDeletes;

/**
 * App\Models\Announcement
 *
 * @property int $id
 * @property int $user_id
 * @property string $title 公告标题
 * @property string $content 公告内容
 * @property \Illuminate\Support\Carbon|null $created_at
 * @property \Illuminate\Support\Carbon|null $updated_at
 * @property \Illuminate\Support\Carbon|null $deleted_at
 * @method static bool|null forceDelete()
 * @method static \Illuminate\Database\Eloquent\Builder|\App\Models\Announcement newModelQuery()
 * @method static \Illuminate\Database\Eloquent\Builder|\App\Models\Announcement newQuery()
 * @method static \Illuminate\Database\Query\Builder|\App\Models\Announcement onlyTrashed()
 * @method static \Illuminate\Database\Eloquent\Builder|\App\Models\Announcement query()
 * @method static bool|null restore()
 * @method static \Illuminate\Database\Eloquent\Builder|\App\Models\Announcement whereContent($value)
 * @method static \Illuminate\Database\Eloquent\Builder|\App\Models\Announcement whereCreatedAt($value)
 * @method static \Illuminate\Database\Eloquent\Builder|\App\Models\Announcement whereDeletedAt($value)
 * @method static \Illuminate\Database\Eloquent\Builder|\App\Models\Announcement whereId($value)
 * @method static \Illuminate\Database\Eloquent\Builder|\App\Models\Announcement whereTitle($value)
 * @method static \Illuminate\Database\Eloquent\Builder|\App\Models\Announcement whereUpdatedAt($value)
 * @method static \Illuminate\Database\Eloquent\Builder|\App\Models\Announcement whereUserId($value)
 * @method static \Illuminate\Database\Query\Builder|\App\Models\Announcement withTrashed()
 * @method static \Illuminate\Database\Query\Builder|\App\Models\Announcement withoutTrashed()
 * @mixin \Eloquent
 */
class Announcement extends Model
{
    use SoftDeletes;
}