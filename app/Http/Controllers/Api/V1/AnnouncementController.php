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

/**
 * Announcement related api
 *
 * @package App\Http\Controllers\Api\V1
 * @Resource("Announcement", uri="/")
 */
class AnnouncementController extends Controller
{
    /**
     * Announcement List
     *
     * Get announcement list
     *
     * @param QueryAnnouncementRequest $request
     * @return bool|\Dingo\Api\Http\Response|void
     *
     * @Get("/announcements")
     * @Versions({"v1"})
     * @Parameters({
     *      @Parameter("page", description="The page of results to view.", default=1),
     *      @Parameter("per_page", description="The amount of results per page.", default=15)
     * })
     * @Response(200, body={
     *          "data": {
     *              {
     *                  "id": 2,
     *                  "user_id": 1,
     *                  "title": "admin",
     *                  "content": "admin@admin.com",
     *                  "created_at": "2019-08-18 16:33:31",
     *                  "updated_at": "2019-08-18 16:33:31",
     *                  "deleted_at": null
     *              }
     *          },
     *          "meta": {"pagination": {"total": 1, "count": 1, "per_page": 15, "current_page": 1, "total_pages": 1}}
     *      })
     */
    public function getAnnouncementList(QueryAnnouncementRequest $request)
    {
        $result = AnnouncementService::getAnnouncementList(new QueryFilter($request->all()));

        if ($result->getStatus() && $result->getData() instanceof Paginator) {
            return $this->response->paginator($result->getData(), new BasicTransformer());
        }

        return $this->response->errorBadRequest() ?? false;
    }

    /**
     * Create Announcement
     *
     * Create an announcement
     *
     * @param CreateAnnouncementRequest $request
     * @return bool|\Dingo\Api\Http\Response|void
     *
     * @Post("/announcements")
     * @Versions({"v1"})
     * @Transaction({
     *      @Request({"title": "title", "content": "content"}),
     *      @Response(200, body={"data": {"id": 1, "title": "", "content": ""}})
     * })
     */
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

    /**
     * Retrieve Announcement
     *
     * Get an announcement
     *
     * @param $id
     * @return bool|\Dingo\Api\Http\Response|void
     *
     * @Get("/announcements/{id}")
     * @Versions({"v1"})
     * @Transaction({
     *      @Response(200, body={"data": {"id": 1, "title": "", "content": ""}})
     * })
     */
    public function getAnnouncement($id)
    {
        $result = AnnouncementService::getAnnouncementById($id);

        if ($result->getStatus()) {
            return $this->response->item($result->getData(), new BasicTransformer());
        }

        return $this->response->errorNotFound() ?? false;
    }

    /**
     * Update Announcement
     *
     * Update an announcement
     *
     * @param UpdateAnnouncementRequest $request
     * @param $id
     * @return bool|\Dingo\Api\Http\Response|void
     *
     * @Put("/announcements/{id}")
     * @Versions({"v1"})
     * @Transaction({
     *      @Request({"title": "title", "content": "content"}),
     *      @Response(200, body={"data": {"id": 1, "title": "", "content": ""}})
     * })
     */
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

    /**
     * Delete Announcement
     *
     * Delete announcement
     *
     * @param $id
     * @return bool|\Dingo\Api\Http\Response|void
     *
     * @Delete("/announcements/{id}")
     * @Versions({"v1"})
     * @Transaction({
     *      @Response(200)
     * })
     */
    public function deleteAnnouncement($id)
    {
        $result = AnnouncementService::deleteAnnouncement($id);

        if ($result->getStatus()) {
            return $this->response->noContent();
        }

        return $this->response->errorBadRequest() ?? false;
    }
}