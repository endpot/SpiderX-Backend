<?php

namespace App\Services;

use App\Carriers\DataCarrierInterface;
use App\Filters\QueryFilterInterface;
use Exception;
use App\Models\Chat;
use App\Results\CommonErrorResult;
use App\Results\CommonSuccessResult;
use Auth;

class ChatService
{
    public function getChatList(QueryFilterInterface $filter)
    {
        $paginateResult = Chat::paginate($filter->get('per_page', 15));

        if ($paginateResult) {
            return new CommonSuccessResult($paginateResult);
        }

        return new CommonErrorResult('No data');
    }

    public function createChat(DataCarrierInterface $carrier)
    {
        $currentUser = Auth::guard()->user();

        if ($currentUser) {
            $chat = new Chat();
            $chat->user_id = $currentUser->getAuthIdentifier();
            $chat->content = $carrier->get('content', '');

            if ($chat->save()) {
                return new CommonSuccessResult($chat);
            }
        }

        return new CommonErrorResult('Error Data');
    }
    
    public function getChatById($id)
    {
        $chat = Chat::find($id);

        if ($chat) {
            return new CommonSuccessResult($chat);
        }

        return new CommonErrorResult('Not found');
    }

    public function deleteChat($id)
    {
        $chat = Chat::find($id);

        if ($chat) {
            try {
                if ($chat->delete()) {
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