<?php

namespace App\Services;

use Exception;
use App\Models\Announcement;
use App\Results\CommonErrorResult;
use App\Results\CommonSuccessResult;
use Auth;

class AnnouncementService
{
    public function getAnnouncementList($perPage = 15)
    {
        $paginateResult = Announcement::paginate($perPage);

        if ($paginateResult) {
            return new CommonSuccessResult($paginateResult);
        }

        return new CommonErrorResult('No data');
    }

    public function createAnnouncement($data)
    {
        $currentUser = Auth::guard()->user();

        if ($currentUser) {
            $announcement = new Announcement();
            $announcement->user_id = $currentUser->id;
            $announcement->title = $data['title'] ?? '';
            $announcement->content = $data['content'] ?? '';

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

    public function updateAnnouncement($id, $data)
    {
        $currentUser = Auth::guard()->user();

        $announcement = Announcement::find($id);

        if ($currentUser && $announcement) {
            $announcement->title = $data['title'] ?? '';
            $announcement->content = $data['content'] ?? '';

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