FROM golang:1.21-alpine as builder
WORKDIR /
COPY . ./
RUN go mod download


RUN go build -o /vancouver-snow-conditions


FROM alpine
COPY --from=builder /vancouver-snow-conditions .


EXPOSE 80
CMD [ "/vancouver-snow-conditions" ]