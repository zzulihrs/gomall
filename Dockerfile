FROM docker.1ms.run/golang:1.21

WORKDIR /usr/src/gomall

ENV GOPROXY https://goproxy.io,direct

COPY app/frontend/go.mod app/frontend/go.sum ./app/frontend/
COPY rpc_gen rpc_gen
RUN cd app/frontend/ && go mod download && go mod tidy

COPY app/frontend/ app/frontend/

RUN cd app/frontend/ && go build -o /opt/gomall/frontend/server

COPY app/frontend/conf /opt/gomall/frontend/conf
COPY app/frontend/static /opt/gomall/frontend/static
COPY app/frontend/template /opt/gomall/frontend/template

WORKDIR /opt/gomall/frontend
EXPOSE 8080

CMD ["./server"]