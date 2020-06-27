## docker-fluentd-mongo

Why:
As a developer,
I want to stream my logs to mongodb,
In order to view the insights on metabase.

How:
- add mongod, fluentd and server
- stream logs to fluentd from server
- add plugin to write to mongo

Start:
```bash
$ docker-compose up -d
```

## Setting up metabase

1. Open metabase at `localhost:3000`.
2. Add database type `MongoDB`
	- name: mongo
	- host: mongo
	- database name: fluentd
	- database username: john
	- database password: 123456
	- authentication database: admin
3. Use x-ray insights

## Accessing mongo shell

```bash
$ docker exec -it <container_id> mongo -u john -p 123456
```

## Metabase mongodb query

```bash
[{$match: {source: {$eq: 'stderr'}}}]
```
