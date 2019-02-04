dep init || dep ensure
(
  for f in $(ls client/ | sed 's/\.go//')
  do
    echo $f
    go build -o bin/$f client/$f.go
  done
)
go run server.go
