<?php


namespace App\Http\Controllers\Api\V1;

use App\Http\Requests\Api\V1\Chat\CreateChatRequest;
use ChatService;
use App\Http\Controllers\Api\Controller;
use App\Transformers\BasicTransformer;
use Illuminate\Contracts\Pagination\Paginator;

class ChatController extends Controller
{
    public function getChatList()
    {
        $result = ChatService::getChatList();

        if ($result->getStatus() && $result->getData() instanceof Paginator) {
            return $this->response->paginator($result->getData(), new BasicTransformer());
        }

        return $this->response->errorBadRequest() ?? false;
    }

    public function createChat(CreateChatRequest $request)
    {
        $result = ChatService::createChat($request->all());

        if ($result->getStatus()) {
            return $this->response->item($result->getData(), new BasicTransformer());
        }

        return $this->response->errorBadRequest() ?? false;
    }

    public function getChat($id)
    {
        $result = ChatService::getChatById($id);

        if ($result->getStatus()) {
            return $this->response->item($result->getData(), new BasicTransformer());
        }

        return $this->response->errorNotFound() ?? false;
    }

    public function deleteChat($id)
    {
        $result = ChatService::deleteChat($id);

        if ($result->getStatus()) {
            return $this->response->noContent();
        }

        return $this->response->errorBadRequest() ?? false;
    }
}