<?php

namespace App\Http\Controllers\Api\V1;

use App\Carriers\DataCarrier;
use App\Facades\TopicService;
use App\Filters\QueryFilter;
use App\Http\Controllers\Api\Controller;
use App\Http\Requests\Api\V1\Topic\CreatePostRequest;
use App\Http\Requests\Api\V1\Topic\CreateTopicRequest;
use App\Http\Requests\Api\V1\Topic\QueryPostRequest;
use App\Http\Requests\Api\V1\Topic\QueryTopicRequest;
use App\Http\Requests\Api\V1\Topic\UpdateTopicRequest;
use App\Transformers\BasicTransformer;
use Illuminate\Contracts\Pagination\Paginator;

class TopicController extends Controller
{
    public function getTopicList(QueryTopicRequest $request)
    {
        $result = TopicService::getTopicList(new QueryFilter($request->all()));

        if ($result->getStatus() && $result->getData() instanceof Paginator) {
            return $this->response->paginator($result->getData(), new BasicTransformer());
        }

        return $this->response->errorBadRequest() ?? false;
    }

    public function createTopic(CreateTopicRequest $request)
    {
        $result = TopicService::createTopic(new DataCarrier($request->all()));

        if ($result->getStatus()) {
            return $this->response->item($result->getData(), new BasicTransformer());
        }

        return $this->response->errorBadRequest() ?? false;
    }

    public function getTopic($id)
    {
        $result = TopicService::getTopicById($id);

        if ($result->getStatus()) {
            return $this->response->item($result->getData(), new BasicTransformer());
        }

        return $this->response->errorNotFound() ?? false;
    }

    public function updateTopic(UpdateTopicRequest $request, $id)
    {
        $result = TopicService::updateTopic(
            $id,
            new DataCarrier($request->all())
        );

        if ($result->getStatus()) {
            return $this->response->item($result->getData(), new BasicTransformer());
        }

        return $this->response->errorBadRequest() ?? false;
    }

    public function deleteTopic($id)
    {
        $result = TopicService::deleteTopic($id);

        if ($result->getStatus()) {
            return $this->response->noContent();
        }

        return $this->response->errorBadRequest() ?? false;
    }

    public function getPostList(QueryPostRequest $request, $id)
    {
        $result = TopicService::getPostListOfTopic($id, new QueryFilter($request->all()));

        if ($result->getStatus() && $result->getData() instanceof Paginator) {
            return $this->response->paginator($result->getData(), new BasicTransformer());
        }

        return $this->response->errorBadRequest() ?? false;
    }

    public function createPost(CreatePostRequest $request, $id)
    {
        $result = TopicService::createPostOfTopic($id, new DataCarrier($request->all()));

        if ($result->getStatus()) {
            return $this->response->item($result->getData(), new BasicTransformer());
        }

        return $this->response->errorBadRequest() ?? false;
    }
}
