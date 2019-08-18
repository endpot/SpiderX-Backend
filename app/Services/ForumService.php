<?php

namespace App\Services;

use App\Carriers\DataCarrierInterface;
use App\Filters\QueryFilterInterface;
use App\Models\Forum;
use App\Results\CommonErrorResult;
use App\Results\CommonSuccessResult;
use Exception;

class ForumService
{
    public function getForumList(QueryFilterInterface $filter)
    {
        $paginateResult = Forum::paginate($filter->get('per_page', 15));

        if ($paginateResult) {
            return new CommonSuccessResult($paginateResult);
        }

        return new CommonErrorResult('No data');
    }

    public function createForum(DataCarrierInterface $carrier)
    {
        $forum = new Forum();
        $forum->name = $carrier->get('name', '');
        $forum->desc = $carrier->get('desc', '');
        $forum->sort = $carrier->get('sort', 0);

        if ($forum->save()) {
            return new CommonSuccessResult($forum);
        }

        return new CommonErrorResult('Error Data');
    }
    
    public function getForumById($id)
    {
        $chat = Forum::find($id);

        if ($chat) {
            return new CommonSuccessResult($chat);
        }

        return new CommonErrorResult('Not found');
    }

    public function updateForum($id, DataCarrierInterface $carrier)
    {
        $forum = Forum::find($id);

        if ($forum) {
            $forum->name = $carrier->get('name', '');
            $forum->desc = $carrier->get('desc', '');
            $forum->sort = $carrier->get('sort', 0);

            if ($forum->save()) {
                return new CommonSuccessResult($forum);
            }
        }

        return new CommonErrorResult('Error Data');
    }

    public function deleteForum($id)
    {
        $forum = Forum::find($id);

        if ($forum) {
            try {
                if ($forum->delete()) {
                    return new CommonSuccessResult();
                }
            } catch (Exception $exception) {
                //
            }

            return new CommonErrorResult('Something wrong');
        }

        return new CommonSuccessResult();
    }
}
