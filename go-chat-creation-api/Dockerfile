FROM golang:1.19
WORKDIR /go/src/github.com/cherifBurette1/rails-go-chat/tree/master
RUN go mod init github.com/cherifBurette1/rails-go-chat/tree/master
RUN go mod tidy
RUN mkdir /go/src/github.com/cherifBurette1/rails-go-chat/tree/master/go-chat-creation-api

COPY . /go/src/github.com/cherifBurette1/rails-go-chat/tree/master/go-chat-creation-api
COPY entrypoint.sh /usr/bin/entrypoint-go.sh
RUN chmod +x /usr/bin/entrypoint-go.sh
COPY wait-for-it.sh /usr/bin
RUN chmod +x /usr/bin/wait-for-it.sh
ENTRYPOINT ["entrypoint-go.sh"]
EXPOSE 8080

CMD ["/go/src/github.com/cherifBurette1/rails-go-chat/tree/master/go-chat-creation-api/cmd/go-chat-creation-api/go-chat-creation-api"]
