<?php

namespace App\Models;

use Hash;
use Illuminate\Notifications\Notifiable;
use Illuminate\Foundation\Auth\User as Authenticatable;
use Spatie\Permission\Traits\HasRoles;
use Tymon\JWTAuth\Contracts\JWTSubject;

/**
 * App\Models\User
 *
 * @property-read \Illuminate\Notifications\DatabaseNotificationCollection|\Illuminate\Notifications\DatabaseNotification[] $notifications
 * @property-read \Illuminate\Database\Eloquent\Collection|\Spatie\Permission\Models\Permission[] $permissions
 * @property-read \Illuminate\Database\Eloquent\Collection|\Spatie\Permission\Models\Role[] $roles
 * @property-write mixed $password
 * @method static \Illuminate\Database\Eloquent\Builder|\App\Models\User newModelQuery()
 * @method static \Illuminate\Database\Eloquent\Builder|\App\Models\User newQuery()
 * @method static \Illuminate\Database\Eloquent\Builder|\App\Models\User permission($permissions)
 * @method static \Illuminate\Database\Eloquent\Builder|\App\Models\User query()
 * @method static \Illuminate\Database\Eloquent\Builder|\App\Models\User role($roles, $guard = null)
 * @mixin \Eloquent
 * @property int $id
 * @property string $name 用户名
 * @property string $email 注册邮箱
 * @property string $passkey Tracker 通信密钥
 * @property int $state 账号状态
 * @property string|null $last_login_date 最近登录时间
 * @property string|null $last_access_date 最近访问时间
 * @property string|null $last_access_ip 最近访问IP
 * @property int $class 用户等级
 * @property int $max_class_once 曾经最高等级
 * @property string $title 头衔
 * @property string $avatar 头像
 * @property string|null $info 个人说明
 * @property int $gender 性别
 * @property int $country_id 所属国家
 * @property int $school_id 所属学校
 * @property int $inviter 邀请人
 * @property int $invite_count 邀请数量
 * @property int $up_traffic 上传量
 * @property int $down_traffic 下载量
 * @property int $seed_time 做种时间
 * @property int $leech_time 下载时间
 * @property float $seed_bonus 做种积分
 * @property int $isp 运营商
 * @property int $up_bandwidth 上传带宽
 * @property int $down_bandwidth 下载带宽
 * @property int $is_vip 是否 VIP
 * @property string|null $vip_added_at VIP 添加日期
 * @property string|null $vip_until VIP 有效期
 * @property int $is_donor 是否捐赠用户
 * @property float $total_donation 捐赠金额
 * @property int $is_warned 是否 VIP
 * @property int $warned_type 警告类型
 * @property string|null $warned_at 警告日期
 * @property string|null $warned_by 警告人
 * @property string|null $warned_until 警告截止日期
 * @property \Illuminate\Support\Carbon|null $created_at
 * @property \Illuminate\Support\Carbon|null $updated_at
 * @property string|null $deleted_at
 * @method static \Illuminate\Database\Eloquent\Builder|\App\Models\User whereAvatar($value)
 * @method static \Illuminate\Database\Eloquent\Builder|\App\Models\User whereClass($value)
 * @method static \Illuminate\Database\Eloquent\Builder|\App\Models\User whereCountryId($value)
 * @method static \Illuminate\Database\Eloquent\Builder|\App\Models\User whereCreatedAt($value)
 * @method static \Illuminate\Database\Eloquent\Builder|\App\Models\User whereDeletedAt($value)
 * @method static \Illuminate\Database\Eloquent\Builder|\App\Models\User whereDownBandwidth($value)
 * @method static \Illuminate\Database\Eloquent\Builder|\App\Models\User whereDownTraffic($value)
 * @method static \Illuminate\Database\Eloquent\Builder|\App\Models\User whereEmail($value)
 * @method static \Illuminate\Database\Eloquent\Builder|\App\Models\User whereGender($value)
 * @method static \Illuminate\Database\Eloquent\Builder|\App\Models\User whereId($value)
 * @method static \Illuminate\Database\Eloquent\Builder|\App\Models\User whereInfo($value)
 * @method static \Illuminate\Database\Eloquent\Builder|\App\Models\User whereInviteCount($value)
 * @method static \Illuminate\Database\Eloquent\Builder|\App\Models\User whereInviter($value)
 * @method static \Illuminate\Database\Eloquent\Builder|\App\Models\User whereIsDonor($value)
 * @method static \Illuminate\Database\Eloquent\Builder|\App\Models\User whereIsVip($value)
 * @method static \Illuminate\Database\Eloquent\Builder|\App\Models\User whereIsWarned($value)
 * @method static \Illuminate\Database\Eloquent\Builder|\App\Models\User whereIsp($value)
 * @method static \Illuminate\Database\Eloquent\Builder|\App\Models\User whereLastAccessDate($value)
 * @method static \Illuminate\Database\Eloquent\Builder|\App\Models\User whereLastAccessIp($value)
 * @method static \Illuminate\Database\Eloquent\Builder|\App\Models\User whereLastLoginDate($value)
 * @method static \Illuminate\Database\Eloquent\Builder|\App\Models\User whereLeechTime($value)
 * @method static \Illuminate\Database\Eloquent\Builder|\App\Models\User whereMaxClassOnce($value)
 * @method static \Illuminate\Database\Eloquent\Builder|\App\Models\User whereName($value)
 * @method static \Illuminate\Database\Eloquent\Builder|\App\Models\User wherePasskey($value)
 * @method static \Illuminate\Database\Eloquent\Builder|\App\Models\User wherePassword($value)
 * @method static \Illuminate\Database\Eloquent\Builder|\App\Models\User whereSchoolId($value)
 * @method static \Illuminate\Database\Eloquent\Builder|\App\Models\User whereSeedBonus($value)
 * @method static \Illuminate\Database\Eloquent\Builder|\App\Models\User whereSeedTime($value)
 * @method static \Illuminate\Database\Eloquent\Builder|\App\Models\User whereState($value)
 * @method static \Illuminate\Database\Eloquent\Builder|\App\Models\User whereTitle($value)
 * @method static \Illuminate\Database\Eloquent\Builder|\App\Models\User whereTotalDonation($value)
 * @method static \Illuminate\Database\Eloquent\Builder|\App\Models\User whereUpBandwidth($value)
 * @method static \Illuminate\Database\Eloquent\Builder|\App\Models\User whereUpTraffic($value)
 * @method static \Illuminate\Database\Eloquent\Builder|\App\Models\User whereUpdatedAt($value)
 * @method static \Illuminate\Database\Eloquent\Builder|\App\Models\User whereVipAddedAt($value)
 * @method static \Illuminate\Database\Eloquent\Builder|\App\Models\User whereVipUntil($value)
 * @method static \Illuminate\Database\Eloquent\Builder|\App\Models\User whereWarnedAt($value)
 * @method static \Illuminate\Database\Eloquent\Builder|\App\Models\User whereWarnedBy($value)
 * @method static \Illuminate\Database\Eloquent\Builder|\App\Models\User whereWarnedType($value)
 * @method static \Illuminate\Database\Eloquent\Builder|\App\Models\User whereWarnedUntil($value)
 */
class User extends Authenticatable implements JWTSubject
{
    use Notifiable, HasRoles;

    protected $guard_name = 'api';

    /**
     * The attributes that are mass assignable.
     *
     * @var array
     */
    protected $fillable = [
        'name',
        'email',
        'password',
    ];

    /**
     * The attributes that should be hidden for arrays.
     *
     * @var array
     */
    protected $hidden = [
        'password',
    ];

    /**
     * Automatically creates hash for the user password.
     *
     * @param  string $value
     * @return void
     */
    public function setPasswordAttribute($value)
    {
        $this->attributes['password'] = Hash::make($value);
    }

    /**
     * Get the identifier that will be stored in the subject claim of the JWT.
     *
     * @return mixed
     */
    public function getJWTIdentifier()
    {
        return $this->getKey();
    }

    /**
     * Return a key value array, containing any custom claims to be added to the JWT.
     *
     * @return array
     */
    public function getJWTCustomClaims()
    {
        return [];
    }
}
