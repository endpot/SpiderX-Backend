<?php
/**
 * Created by PhpStorm.
 * User: swxuan
 * Date: 6/25/2019
 * Time: 3:47 PM
 */

namespace App\Transformers;

use Illuminate\Contracts\Support\Arrayable;
use League\Fractal\TransformerAbstract;

class BasicTransformer extends TransformerAbstract
{
    public function transform($model)
    {
        if ($model instanceof Arrayable) {
            return $model->toArray();
        } elseif ($model instanceof \stdClass) {
            return (array)$model;
        }

        return [];
    }
}