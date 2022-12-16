FROM golang:1.20rc1-buster
WORKDIR /usr/share/ivan-api/server
COPY ./server/go.mod /usr/share/ivan-api/server
COPY ./server/go.sum /usr/share/ivan-api/server
RUN cd /usr/share/ivan-api/server && go mod download

ADD ./server /usr/share/ivan-api/server

RUN go build -o ./main ./cmd/ivan-server
EXPOSE 8080


CMD [ "./main serve" ]