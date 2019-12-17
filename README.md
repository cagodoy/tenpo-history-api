# History-API

Microservice implemented in Golang that stores history information into postgres DB.

## Table

```
   Column   |           Type           | Collation | Nullable |      Default
------------+--------------------------+-----------+----------+-------------------
 id         | uuid                     |           | not null | gen_random_uuid()
 user_id    | character varying(255)   |           | not null |
 latitude   | character varying(255)   |           | not null |
 longitude  | character varying(255)   |           | not null |
 created_at | timestamp with time zone |           |          | now()
 updated_at | timestamp with time zone |           |          | now()
 deleted_at | timestamp with time zone |           |          |
Indexes:
    "history_pkey" PRIMARY KEY, btree (id)
Triggers:
    update_history_update_at BEFORE UPDATE ON history FOR EACH ROW EXECUTE PROCEDURE update_updated_at_column()
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
