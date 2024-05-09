## Development

### Test self update program

1. compile `v1` program to `bin/v1_serve`

```
go build -o bin/v1_serve
```

2. compile `v2` program to `bin/v2_serve`

```
go build -o bin/v2_serve -ldflags "-X dorobo/web.version=v0.0.2"
```

3. execute `bin/v1_serve` and check the version and pid of stdout from the process (default http port: 8090)

```
./bin/v1_serve
```

4. invoke API `version` to check the version of program, you should get `v0.0.1` via response message

```
curl -X GET http://localhost:8090/version
```

5. invoke upgrade API via `curl`

```
curl -X POST -H 'Content-Type: application/json' -d '{"path": "bin/v2_serve"}' http://localhost:8090/upgrade
```

6. invoke API `version` again, you should get `v0.0.2` via response message

```
curl -X GET http://localhost:8090/version
```

7. kill the process (take the look the pid from stdout of the first terminal section)

```
kill <pid>
```
