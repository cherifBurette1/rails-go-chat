package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/cherifBurette1/rails-go-chat/tree/master/go-chat-creation-api/configs"
	"github.com/cherifBurette1/rails-go-chat/tree/master/go-chat-creation-api/pkg/network"
	"github.com/cherifBurette1/rails-go-chat/tree/master/go-chat-creation-api/pkg/redis"
	"github.com/cherifBurette1/rails-go-chat/tree/master/go-chat-creation-api/pkg/sidekiq"
	"github.com/gorilla/mux"
	"github.com/imroc/req"
)

type messageRequest struct {
	Body string `json:"body"`
}

type messageResponse struct {
	Number      int64  `json:"number"`
	ChatNumber  int64  `json:"chat_number"`
	AccessToken string `json:"access_token"`
}

type chatApiResponse struct {
	Number       int64  `json:"number"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
	MessageCount int64  `json:"message_count"`
}

func CreateMessage(w http.ResponseWriter, r *http.Request) {
	// Read in request
	accessToken := mux.Vars(r)["access_token"]
	chatNumber := mux.Vars(r)["chat_number"]

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		network.RespondErr(w, err)
		return
	}

	var req messageRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		network.RespondErr(w, err)
		return
	}

	// Get next number
	redisClient, _ := redis.GetRedis()
	key := configs.MsgRedisKeyPrefix + accessToken + "_" + chatNumber

	exists, err := redisClient.Exists(ctx, key).Result()
	if err != nil {
		network.RespondErr(w, err)
		return
	} else if exists == 0 {
		log.Println("Key " + key + " not found in Redis, requsting message count from Rails instead")
		chatsResp, err := RequestMessages(accessToken, chatNumber)
		if err != nil {
			network.RespondErr(w, err)
			return
		}
		redisClient.Set(ctx, key, chatsResp.MessageCount, 1)
	}

	nextNum, err := redisClient.Incr(ctx, key).Result()
	if err != nil {
		network.RespondErr(w, err)
		return
	}

	// Push to Sidekiq
	err = sidekiq.Push(configs.MsgRedisQueue, configs.MsgWorkerClass, accessToken, chatNumber, strconv.FormatInt(nextNum, 10), req.Body)
	if err != nil {
		network.RespondErr(w, err)
		return
	}

	// Send response
	chatNumInt64, _ := strconv.ParseInt(chatNumber, 10, 64)
	resp := messageResponse{Number: nextNum, ChatNumber: chatNumInt64, AccessToken: accessToken}
	respBytes, err := json.Marshal(resp)
	if err != nil {
		network.RespondErr(w, err)
		return
	}

	network.Respond(w, respBytes, http.StatusCreated)
}

func RequestMessages(accessToken string, chatNumber string) (chatApiResponse, error) {
	var resp chatApiResponse

	url := strings.Replace(configs.AppEndpoint+configs.MessagesRoute, "{access_token}", accessToken, 1)
	url = strings.Replace(url, "{chat_number}", chatNumber, 1)

	r, err := req.Get(url)
	if err != nil {
		return resp, err
	}

	r.ToJSON(&resp)
	return resp, nil
}
