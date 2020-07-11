package router

import (
	"github.com/endpot/SpiderX-Backend/app/controller/handler/auth"
	"github.com/endpot/SpiderX-Backend/app/controller/handler/chat"
	"github.com/endpot/SpiderX-Backend/app/controller/handler/comment"
	"github.com/endpot/SpiderX-Backend/app/controller/handler/forum"
	"github.com/endpot/SpiderX-Backend/app/controller/handler/message"
	"github.com/endpot/SpiderX-Backend/app/controller/handler/notice"
	"github.com/endpot/SpiderX-Backend/app/controller/handler/post"
	"github.com/endpot/SpiderX-Backend/app/controller/handler/subject"
	"github.com/endpot/SpiderX-Backend/app/controller/handler/subtitle"
	"github.com/endpot/SpiderX-Backend/app/controller/handler/topic"
	"github.com/endpot/SpiderX-Backend/app/controller/handler/torrent"
	"github.com/endpot/SpiderX-Backend/app/controller/handler/user"
	"github.com/endpot/SpiderX-Backend/app/controller/middleware"
	_ "github.com/endpot/SpiderX-Backend/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"log"
)

func InitRouter() *gin.Engine {
	log.Println("Router Initializing...")

	r := gin.New()

	// Global middleware
	r.Use(gin.Logger(), gin.Recovery())
	r.Use(middleware.Cors())

	// Robots.txt
	r.StaticFile("/robots.txt", "./storage/public/robots.txt")

	// Swagger Document
	r.GET("/_doc/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	{
		// Authentication related API
		r.POST("/auth.register", auth.Register)
		r.POST("/auth.login", auth.Login)
		r.POST("/auth.logout", auth.Logout)
		r.POST("/auth.reset", auth.Reset)
		r.POST("/auth.getMine", auth.GetMine)
		r.POST("/auth.sendCode", auth.SendCode)
		r.POST("/auth.refreshToken", auth.RefreshToken)
	}

	{
		// User related API
		r.GET("/users", user.GetUserList)
		r.GET("/users/:userID", user.GetUser)
		r.POST("/users", user.CreateUser)
		r.PATCH("/users/:userID", user.UpdateUser)
		r.DELETE("/users/:userID", user.DeleteUser)
	}

	{
		// Notice related API
		r.GET("/notices", notice.GetNoticeList)
		r.GET("/notices/:noticeID", notice.GetNotice)
		r.POST("/notices", notice.CreateNotice)
		r.PATCH("/notices/:noticeID", notice.UpdateNotice)
		r.DELETE("/notices/:noticeID", notice.DeleteNotice)
	}

	{
		// Forum related API
		r.GET("/forums", forum.GetForumList)
		r.GET("/forums/:forumID", forum.GetForum)
		r.POST("/forums", forum.CreateForum)
		r.PATCH("/forums/:forumID", forum.UpdateForum)
		r.DELETE("/forums/:forumID", forum.DeleteForum)
	}

	{
		// Topic related API
		r.GET("/forums/:forumID/topics", topic.GetTopicList)
		r.GET("/forums/:forumID/topics/:topicID", topic.GetTopic)
		r.POST("/forums/:forumID/topics", topic.CreateTopic)
		r.PATCH("/forums/:forumID/topics/:topicID", topic.UpdateTopic)
		r.DELETE("/forums/:forumID/topics/:topicID", topic.DeleteTopic)
	}

	{
		// Post related API
		r.GET("/topics/:topicID/posts", post.GetPostList)
		r.GET("/topics/:topicID/posts/:postID", post.GetPost)
		r.POST("/topics/:topicID/posts", post.CreatePost)
		r.PATCH("/topics/:topicID/posts/:postID", post.UpdatePost)
		r.DELETE("/topics/:topicID/posts/:postID", post.DeletePost)
	}

	{
		// Subject related API
		r.GET("/subjects", subject.GetSubjectList)
		r.GET("/subjects/:subjectID", subject.GetSubject)
		r.POST("/subjects", subject.CreateSubject)
		r.PATCH("/subjects/:subjectID", subject.UpdateSubject)
		r.DELETE("/subjects/:subjectID", subject.DeleteSubject)
	}

	{
		// Torrent related API
		r.GET("/torrents", torrent.GetTorrentList)
		r.GET("/torrents/:torrentID", torrent.GetTorrent)
		r.POST("/torrents", torrent.CreateTorrent)
		r.POST("/torrents.preUpload", torrent.PreUploadTorrent)
		r.PATCH("/torrents/:torrentID", torrent.UpdateTorrent)
		r.DELETE("/torrents/:torrentID", torrent.DeleteTorrent)
	}

	{
		// Comment related API
		r.GET("/torrents/:torrentID/comments", comment.GetCommentList)
		r.GET("/torrents/:torrentID/comments/:commentID", comment.GetComment)
		r.POST("/torrents/:torrentID/comments", comment.CreateComment)
		r.PATCH("/torrents/:torrentID/comments/:commentID", comment.UpdateComment)
		r.DELETE("/torrents/:torrentID/comments/:commentID", comment.DeleteComment)
	}

	{
		// Comment related API
		//r.GET("/comments", comment.GetCommentList)
		//r.GET("/comments/:commentID", comment.GetComment)
		//r.POST("/comments", comment.CreateComment)
		//r.PATCH("/comments/:commentID", comment.UpdateComment)
		//r.DELETE("/comments/:commentID", comment.DeleteComment)
	}

	{
		// Subtitle related API
		r.GET("/subtitles", subtitle.GetSubtitleList)
		r.GET("/subtitles/:subtitleID", subtitle.GetSubtitle)
		r.POST("/subtitles", subtitle.CreateSubtitle)
		r.PATCH("/subtitles/:subtitleID", subtitle.UpdateSubtitle)
		r.DELETE("/subtitles/:subtitleID", subtitle.DeleteSubtitle)
	}

	{
		// Message related API
		r.GET("/messages", message.GetMessageList)
		r.GET("/messages/:messageID", message.GetMessage)
		r.POST("/messages", message.CreateMessage)
		r.PATCH("/messages/:messageID", message.UpdateMessage)
		r.DELETE("/messages/:messageID", message.DeleteMessage)

		r.GET("/boxes", message.GetBoxList)
		r.GET("/boxes/:boxID", message.GetBox)
		r.POST("/boxes", message.CreateBox)
		r.PATCH("/boxes/:boxID", message.UpdateBox)
		r.DELETE("/boxes/:boxID", message.DeleteBox)
	}

	{
		// Chat related API
		r.GET("/chats", chat.GetChatList)
		r.GET("/chats/:ChatID", chat.GetChat)
		r.POST("/chats", chat.CreateChat)
		r.PATCH("/chats/:ChatID", chat.UpdateChat)
		r.DELETE("/chats/:ChatID", chat.DeleteChat)
	}

	log.Println("Router Initialized!")

	return r
}
