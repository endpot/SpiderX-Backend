<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\SoftDeletes;

/**
 * App\Models\Topic
 *
 * @property int $id
 * @property int $user_id
 * @property int $forum_id
 * @property string $subject 主题标题
 * @property string $content 主题内容
 * @property int $view_count 查看次数
 * @property int $position 置顶位置
 * @property int $is_hot 是否热门
 * @property \Illuminate\Support\Carbon|null $created_at
 * @property \Illuminate\Support\Carbon|null $updated_at
 * @property \Illuminate\Support\Carbon|null $deleted_at
 * @method static bool|null forceDelete()
 * @method static \Illuminate\Database\Eloquent\Builder|\App\Models\Topic newModelQuery()
 * @method static \Illuminate\Database\Eloquent\Builder|\App\Models\Topic newQuery()
 * @method static \Illuminate\Database\Query\Builder|\App\Models\Topic onlyTrashed()
 * @method static \Illuminate\Database\Eloquent\Builder|\App\Models\Topic query()
 * @method static bool|null restore()
 * @method static \Illuminate\Database\Eloquent\Builder|\App\Models\Topic whereContent($value)
 * @method static \Illuminate\Database\Eloquent\Builder|\App\Models\Topic whereCreatedAt($value)
 * @method static \Illuminate\Database\Eloquent\Builder|\App\Models\Topic whereDeletedAt($value)
 * @method static \Illuminate\Database\Eloquent\Builder|\App\Models\Topic whereForumId($value)
 * @method static \Illuminate\Database\Eloquent\Builder|\App\Models\Topic whereId($value)
 * @method static \Illuminate\Database\Eloquent\Builder|\App\Models\Topic whereIsHot($value)
 * @method static \Illuminate\Database\Eloquent\Builder|\App\Models\Topic wherePosition($value)
 * @method static \Illuminate\Database\Eloquent\Builder|\App\Models\Topic whereSubject($value)
 * @method static \Illuminate\Database\Eloquent\Builder|\App\Models\Topic whereUpdatedAt($value)
 * @method static \Illuminate\Database\Eloquent\Builder|\App\Models\Topic whereUserId($value)
 * @method static \Illuminate\Database\Eloquent\Builder|\App\Models\Topic whereViewCount($value)
 * @method static \Illuminate\Database\Query\Builder|\App\Models\Topic withTrashed()
 * @method static \Illuminate\Database\Query\Builder|\App\Models\Topic withoutTrashed()
 * @mixin \Eloquent
 */
class Topic extends Model
{
    use SoftDeletes;
}