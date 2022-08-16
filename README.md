## A chat app using go and ruby on rails
Chatting backend application using Ruby on rails and GoLang
## Requirements
docker, golang, gotools, ruby v 3.1.2, rails 7.0.0 ang golang 1.19
## Overview
The API is divided into two parts :
- Rails APIS: Performing CRUD Operations on applications, chats, and messages, with the ability to search through messages in chats using `elasticsearch`.
- GOlag APIS: APIS responsible for creating chats and messages.

## Running app
### using cmd run the following:
 - ``` docker-compose down ```
 - ``` docker-compose build ```
 - ``` docker-compose up ```
 - make sure docker is running.
 - make sure that ports `3000` and ports `8080` are available.

##  APIS contract
- you can use postman to test these APIs
- Rails
```
Verb      Endpoint
--------  -----------
HTTPGET   /applications/
HTTPPOST  /applications?name={name}
HTTPGET   /applications/{access_token}
HTTPPUT   /applications/{access_token}?name={name}
HTTPGET   /applications/{access_token}/chats
HTTPGET   /applications/{access_token}/chats/{chat_number}
HTTPGET   /applications/{access_token}/chats/{chat_number}/messages
HTTPGET   /applications/{access_token}/chats/{chat_number}/messages/{message_number}
HTTPGET   /applications/{access_token}/chats/{chat_number}/messages/search?keyword={keyword}
HTTPPUT   /applications/{access_token}/chats/{chat_number}/messages/{message_number}?body={message_body}
```
- Go
```
Verb      Endpoint
--------  -----------
HTTPPOST  /applications/{access_token}/chats/
HTTPPOST  /applications/{access_token}/chats/{chat_number}/messages?body={message_body}
```
#### Examples

##### Creating a new application
```
HTTPPost
'http://localhost:3000/applications?name=app'

{
  "name": "app",
  "access_token": "kQxu2wq98orCpL5JeS8MiVxc",
  "created_at": "2022-08-09T22:17:52.599Z",
  "updated_at": "2022-08-10T22:17:52.599Z",
  "chat_count": 0
}
```

##### Getting messages
``` 
HTTPGet
'http://localhost:3000/applications/fPrv7vr57dkUsP4KfZ4BdSmt/chats/1/messages'

[
  {
    "number": 1,
    "body": "message A",
  "created_at": "2022-08-09T22:17:52.599Z",
  "updated_at": "2022-08-10T22:17:52.599Z"
  },
  {
    "number": 2,
    "body": "message b",
  "created_at": "2022-08-09T22:17:52.599Z",
  "updated_at": "2022-08-10T22:17:52.599Z"
  }
]
```
##### Searching chats
```
HTTPGet 'http://localhost:3000/applications/fPrv7vr57dkUsP4KfZ4BdSmt/chats/1/messages/search?keyword=hi'

[
  {
    "number": 2,
    "body": "hi1",
  "created_at": "2022-08-09T22:17:52.599Z",
  "updated_at": "2022-08-10T22:17:52.599Z"
  }
]
```
##### Creating a new chat
```
HTTPPOST 'http://localhost:8080/applications/fPrv7vr57dkUsP4KfZ4BdSmt/chats'
# output
{
  "number": 1,
  "access_token": "kQxu2wq98orCpL5JeS8MiVxc"
}
```

##### Sending a new message
```
HTTPPost 'http://localhost:8080/applications/fPrv7vr57dkUsP4KfZ4BdSmt/chats/1/messages'
{"body": "Rails stuff"}'

{
  "number":1,
  "chat_number":1,
  "access_token":"kQxu2wq98orCpL5JeS8MiVxc"
 }
 ```
 
 ## Internal details
 - `Redis` is used for caching, determining the next chat number and message number to send to the user, In addition to queuing jobs to `Sidekiq` workers in Rails API.
  -`Sidekiq` workers are the background jobs schedulers to prevent timing out.
  -`elasticsearch` is used to search through messages.
  - Golang Chat/Message creation API first receives this request, gets the application token and the chat number from the request endpoint.
 - It then generates a key that refers to this application token/chat number combination
 - It then tries to get the next number using this key from `Redis` store, if it exists then it automatically gets and increments the value,
 and if it's not, then it sends a request to the main Rails API to get the current message count.
 - After getting the next number, the API queues a job to create this message in `Sidekiq`, and responds to the user with the number it created.
 - When a `Sidekiq` worker pick up a message creation job, it then adds this message to the messages table, updating a [`counter_cache`](https://guides.rubyonrails.org/association_basics.html) automatically.
 - Race conditions are handled in both sides:
  - A race condition may occur in the message creation side when two concurrent requests that use the same chat are unable to find a key/value pair in `Redis`
 at the same time, when this happens they both send a request to the main Rails API and set the count in store with the same value, leading to two responses with the same number,
 to handle this [`redis-lock`](https://github.com/bsm/redis-lock) is used to avoid this issue, requests basically "lock" the key/value pair using
 this key combination and release this lock after writing to the store.
 - Race conditions are handled at the main Rails side using `uniqueness` validations on both chat number and message number.
     

     
## Notes
- [ ] Add a `.env` file that contains all the necessary configurations.
- [ ] Handle any error in `docker` due to incompatibility.
