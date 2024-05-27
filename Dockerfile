FROM golang:1.22-alpine AS stage1
WORKDIR /project/studentapp/

COPY go.* .
RUN  go mod download

COPY . .
RUN go build -o ./cmd/studentAppExec ./cmd/main.go

FROM alpine:latest
WORKDIR /project/studentapp/

COPY --from=stage1 /project/studentapp/cmd/studentAppExec ./cmd/
COPY --from=stage1 /project/studentapp/.env ./

EXPOSE 8085
ENTRYPOINT [ "/project/studentapp/cmd/studentAppExec" ]