<?php

namespace App\Http\Controllers\Api\V1;

use AnnouncementService;
use App\Carriers\DataCarrier;
use App\Filters\QueryFilter;
use App\Http\Controllers\Api\Controller;
use App\Http\Requests\Api\V1\Announcement\CreateAnnouncementRequest;
use App\Http\Requests\Api\V1\Announcement\QueryAnnouncementRequest;
use App\Http\Requests\Api\V1\Announcement\UpdateAnnouncementRequest;
use App\Transformers\BasicTransformer;
use Illuminate\Contracts\Pagination\Paginator;

class AnnouncementController extends Controller
{
    public function getAnnouncementList(QueryAnnouncementRequest $request)
    {
        $result = AnnouncementService::getAnnouncementList(new QueryFilter($request->all()));

        if ($result->getStatus() && $result->getData() instanceof Paginator) {
            return $this->response->paginator($result->getData(), new BasicTransformer());
        }

        return $this->response->errorBadRequest() ?? false;
    }

    public function createAnnouncement(CreateAnnouncementRequest $request)
    {
        $result = AnnouncementService::createAnnouncement(
            new DataCarrier($request->all())
        );

        if ($result->getStatus()) {
            return $this->response->item($result->getData(), new BasicTransformer());
        }

        return $this->response->errorBadRequest() ?? false;
    }

    public function getAnnouncement($id)
    {
        $result = AnnouncementService::getAnnouncementById($id);

        if ($result->getStatus()) {
            return $this->response->item($result->getData(), new BasicTransformer());
        }

        return $this->response->errorNotFound() ?? false;
    }

    public function updateAnnouncement(UpdateAnnouncementRequest $request, $id)
    {
        $result = AnnouncementService::updateAnnouncement(
            $id,
            new DataCarrier($request->all())
        );

        if ($result->getStatus()) {
            return $this->response->item($result->getData(), new BasicTransformer());
        }

        return $this->response->errorBadRequest() ?? false;
    }

    public function deleteAnnouncement($id)
    {
        $result = AnnouncementService::deleteAnnouncement($id);

        if ($result->getStatus()) {
            return $this->response->noContent();
        }

        return $this->response->errorBadRequest() ?? false;
    }
}