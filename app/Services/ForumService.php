<?php

namespace App\Services;

use App\Models\Forum;
use App\Results\CommonErrorResult;
use App\Results\CommonSuccessResult;
use Exception;

class ForumService
{
    public function getForumList($perPage = 15)
    {
        $paginateResult = Forum::paginate($perPage);

        if ($paginateResult) {
            return new CommonSuccessResult($paginateResult);
        }

        return new CommonErrorResult('No data');
    }

    public function createForum($data)
    {
        $forum = new Forum();
        $forum->name = $data['name'] ?? '';
        $forum->desc = $data['desc'] ?? '';
        $forum->sort = $data['sort'] ?? 0;

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

    public function updateForum($id, $data)
    {
        $forum = Forum::find($id);

        if ($forum) {
            $forum->name = $data['name'] ?? '';
            $forum->desc = $data['desc'] ?? '';
            $forum->sort = $data['sort'] ?? 0;

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
