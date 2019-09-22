<?php

namespace App\Services;

use App\Carriers\DataCarrierInterface;
use App\Filters\QueryFilterInterface;
use App\Models\Post;
use App\Models\Topic;
use App\Results\CommonErrorResult;
use App\Results\CommonSuccessResult;
use Auth;
use Exception;

class TopicService
{
    public function getTopicList(QueryFilterInterface $filter)
    {
        $paginateQueryBuilder = Topic::query();

        if ($filter->get('forum_id')) {
            $paginateQueryBuilder = $paginateQueryBuilder->where('forum_id', $filter->get('forum_id'));
        }

        $paginateQueryBuilder = $paginateQueryBuilder->orderBy('position', 'desc')->orderBy('id', 'asc');

        $paginateResult = $paginateQueryBuilder->paginate($filter->get('per_page', 15));

        if ($paginateResult) {
            return new CommonSuccessResult($paginateResult);
        }

        return new CommonErrorResult('No data');
    }

    public function createTopic(DataCarrierInterface $carrier)
    {
        $currentUser = Auth::guard()->user();

        if ($currentUser) {
            $topic = new Topic();
            $topic->user_id = $currentUser->getAuthIdentifier();
            $topic->forum_id = $carrier->get('forum_id', 0);
            $topic->subject = $carrier->get('subject', '');
            $topic->content = $carrier->get('content', '');
            $topic->position = $carrier->get('position', 0);

            if ($topic->save()) {
                return new CommonSuccessResult($topic);
            }
        }

        return new CommonErrorResult('Error Data');
    }
    
    public function getTopicById($id)
    {
        $topic = Topic::find($id);

        if ($topic) {
            return new CommonSuccessResult($topic);
        }

        return new CommonErrorResult('Not found');
    }

    public function updateTopic($id, DataCarrierInterface $carrier)
    {
        $currentUser = Auth::guard()->user();

        $topic = Topic::find($id);

        if ($currentUser && $topic) {
            $topic->user_id = $currentUser->getAuthIdentifier();
            $topic->forum_id = $carrier->get('forum_id', 0);
            $topic->subject = $carrier->get('subject', '');
            $topic->content = $carrier->get('content', '');
            $topic->position = $carrier->get('position', 0);

            if ($topic->save()) {
                return new CommonSuccessResult($topic);
            }
        }

        return new CommonErrorResult('Error Data');
    }

    public function deleteTopic($id)
    {
        $topic = Topic::find($id);

        if ($topic) {
            try {
                if ($topic->delete()) {
                    return new CommonSuccessResult();
                }
            } catch (Exception $exception) {
                //
            }

            return new CommonErrorResult('Something wrong');
        }

        return new CommonSuccessResult();
    }

    public function getPostListOfTopic($id, QueryFilterInterface $filter)
    {
        $paginateQueryBuilder = Post::query();

        $paginateQueryBuilder = $paginateQueryBuilder->where('topic_id', $id);

        $paginateResult = $paginateQueryBuilder->paginate($filter->get('per_page', 15));

        if ($paginateResult) {
            return new CommonSuccessResult($paginateResult);
        }

        return new CommonErrorResult('No data');
    }

    public function createPostOfTopic($id, DataCarrierInterface $carrier)
    {
        $currentUser = Auth::guard()->user();

        if ($currentUser) {
            $post = new Post();
            $post->topic_id = $id;
            $post->user_id = $currentUser->getAuthIdentifier();
            $post->content = $carrier->get('content', '');

            if ($post->save()) {
                return new CommonSuccessResult($post);
            }
        }

        return new CommonErrorResult('Error Data');
    }
}
