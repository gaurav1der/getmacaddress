FROM golang:latest
WORKDIR /app/macadrs
RUN mkdir -p /app/bin
ENV GOPATH=/app/macadrs
ENV GOBIN=/app/bin
COPY ./ ./
RUN go install
CMD $GOBIN/macadrs
