<?php

namespace App\Services;

use App\Carriers\DataCarrierInterface;
use App\Filters\QueryFilterInterface;
use Exception;
use App\Models\Announcement;
use App\Results\CommonErrorResult;
use App\Results\CommonSuccessResult;
use Auth;

class AnnouncementService
{
    public function getAnnouncementList(QueryFilterInterface $filter)
    {
        $paginateResult = Announcement::paginate($filter->get('per_page', 15));

        if ($paginateResult) {
            return new CommonSuccessResult($paginateResult);
        }

        return new CommonErrorResult('No data');
    }

    public function createAnnouncement(DataCarrierInterface $carrier)
    {
        $currentUser = Auth::guard()->user();

        if ($currentUser) {
            $announcement = new Announcement();
            $announcement->user_id = $currentUser->getAuthIdentifier();
            $announcement->title = $carrier->get('title', '');
            $announcement->content = $carrier->get('content', '');

            if ($announcement->save()) {
                return new CommonSuccessResult($announcement);
            }
        }

        return new CommonErrorResult('Error Data');
    }
    
    public function getAnnouncementById($id)
    {
        $announcement = Announcement::find($id);

        if ($announcement) {
            return new CommonSuccessResult($announcement);
        }

        return new CommonErrorResult('Not found');
    }

    public function updateAnnouncement($id, DataCarrierInterface $carrier)
    {
        $currentUser = Auth::guard()->user();

        $announcement = Announcement::find($id);

        if ($currentUser && $announcement) {
            $announcement->title = $carrier->get('title', '');
            $announcement->content = $carrier->get('content', '');

            if ($announcement->save()) {
                return new CommonSuccessResult($announcement);
            }
        }

        return new CommonErrorResult('Error Data');
    }

    public function deleteAnnouncement($id)
    {
        $announcement = Announcement::find($id);

        if ($announcement) {
            try {
                if ($announcement->delete()) {
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