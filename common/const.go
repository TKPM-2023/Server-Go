package common

import "TKPM-Go/pubsub"

type DbType int

const (
	DbTypeCategory      DbType = 1
	DbTypeUser          DbType = 2
	DbTypeProduct       DbType = 3
	DbTypeOrder         DbType = 4
	DbTypeOrder_Detail  DbType = 5
	DbTypeProductRating DbType = 6
	DbTypeContact		DbType = 7
)

const CurrentUser = "user"

type Requester interface {
	GetUserId() int
	GetUserEmail() string
	GetUserRole() string
}

const (
	TopicUserAddProduct          pubsub.Topic = "TopicUserAddProduct"
	TopicUserDeleteProduct       pubsub.Topic = "TopicUserDeleteProduct"
	TopicUserRatingProduct       pubsub.Topic = "TopicUserRatingProduct"
	TopicUserDeleteRatingProduct pubsub.Topic = "TopicUserDeleteRatingProduct"
)
