FROM golang:1.11.5

RUN mkdir -p /go/src/web_db

WORKDIR /go/src/ToDoList

COPY . /go/src/ToDoList

RUN go get github.com/astaxie/beego && go get github.com/beego/bee && go get gopkg.in/mgo.v2

EXPOSE 8080

CMD ["bee", "run"]