<?php

namespace App\Http\Controllers\Api\V1;

use App\Carriers\DataCarrier;
use App\Facades\ForumService;
use App\Filters\QueryFilter;
use App\Http\Controllers\Api\Controller;
use App\Http\Requests\Api\V1\Forum\CreateForumRequest;
use App\Http\Requests\Api\V1\Forum\QueryForumRequest;
use App\Http\Requests\Api\V1\Forum\UpdateForumRequest;
use App\Transformers\BasicTransformer;
use Illuminate\Contracts\Pagination\Paginator;

class ForumController extends Controller
{
    public function getForumList(QueryForumRequest $request)
    {
        $result = ForumService::getForumList(new QueryFilter($request->all()));

        if ($result->getStatus() && $result->getData() instanceof Paginator) {
            return $this->response->paginator($result->getData(), new BasicTransformer());
        }

        return $this->response->errorBadRequest() ?? false;
    }

    public function createForum(CreateForumRequest $request)
    {
        $result = ForumService::createForum(
            new DataCarrier($request->all())
        );

        if ($result->getStatus()) {
            return $this->response->item($result->getData(), new BasicTransformer());
        }

        return $this->response->errorBadRequest() ?? false;
    }

    public function getForum($id)
    {
        $result = ForumService::getForumById($id);

        if ($result->getStatus()) {
            return $this->response->item($result->getData(), new BasicTransformer());
        }

        return $this->response->errorNotFound() ?? false;
    }

    public function updateForum(UpdateForumRequest $request, $id)
    {
        $result = ForumService::updateForum(
            $id,
            new DataCarrier($request->all())
        );

        if ($result->getStatus()) {
            return $this->response->item($result->getData(), new BasicTransformer());
        }

        return $this->response->errorBadRequest() ?? false;
    }

    public function deleteForum($id)
    {
        $result = ForumService::deleteForum($id);

        if ($result->getStatus()) {
            return $this->response->noContent();
        }

        return $this->response->errorBadRequest() ?? false;
    }
}