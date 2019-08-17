<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\SoftDeletes;

/**
 * App\Models\Forum
 *
 * @property int $id
 * @property string $name 版块名称
 * @property string|null $desc 版块说明
 * @property int $sort 版块排序
 * @property \Illuminate\Support\Carbon|null $created_at
 * @property \Illuminate\Support\Carbon|null $updated_at
 * @property \Illuminate\Support\Carbon|null $deleted_at
 * @method static bool|null forceDelete()
 * @method static \Illuminate\Database\Eloquent\Builder|\App\Models\Forum newModelQuery()
 * @method static \Illuminate\Database\Eloquent\Builder|\App\Models\Forum newQuery()
 * @method static \Illuminate\Database\Query\Builder|\App\Models\Forum onlyTrashed()
 * @method static \Illuminate\Database\Eloquent\Builder|\App\Models\Forum query()
 * @method static bool|null restore()
 * @method static \Illuminate\Database\Eloquent\Builder|\App\Models\Forum whereCreatedAt($value)
 * @method static \Illuminate\Database\Eloquent\Builder|\App\Models\Forum whereDeletedAt($value)
 * @method static \Illuminate\Database\Eloquent\Builder|\App\Models\Forum whereDesc($value)
 * @method static \Illuminate\Database\Eloquent\Builder|\App\Models\Forum whereId($value)
 * @method static \Illuminate\Database\Eloquent\Builder|\App\Models\Forum whereName($value)
 * @method static \Illuminate\Database\Eloquent\Builder|\App\Models\Forum whereSort($value)
 * @method static \Illuminate\Database\Eloquent\Builder|\App\Models\Forum whereUpdatedAt($value)
 * @method static \Illuminate\Database\Query\Builder|\App\Models\Forum withTrashed()
 * @method static \Illuminate\Database\Query\Builder|\App\Models\Forum withoutTrashed()
 * @mixin \Eloquent
 */
class Forum extends Model
{
    use SoftDeletes;
}