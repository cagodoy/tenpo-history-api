# History-API

Microservice implemented in Golang that stores history information into postgres DB.

## Table

```

```

## GRPC Service

```go
service HistoryService {
  rpc ListHistoryByUserId(HistoryListByUserIdRequest) returns (HistoryListByUserIdResponse) {}
}

message History {
  string id = 1;
  string user_id = 2;
  string latitude = 3;
  string longitude = 4;

  int64 created_at = 5;
	int64 updated_at = 6;
}

message HistoryListByUserIdRequest {
  string user_id = 1;
}

message HistoryListByUserIdResponse {
  repeated History data = 1;
  Meta meta = 2;
  Error error = 3;
}
```

## Commands (Development)

`make build`: build user service for osx.

`make linux`: build user service for linux os.

`make docker`: build docker.

`docker run -it -p 5020:5020 tenpo-history`: run docker.

`PORT=<port> POSTGRES_DSN=<postgres_dsn> NATS_HOST=<nats_host> NATS_PORT=<nats_port> ./bin/tenpo-history-api`: run tenpo history service.
