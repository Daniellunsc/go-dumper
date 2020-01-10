# GO-DUMPER

Go-Dumper is a Telegram bot that dumps a database into a file and send it via message with the command: `/dump`.

# Configuration

## Via Docker
- You can simple replace the env at [Dockerfile](https://github.com/Daniellunsc/go-dumper/blob/master/Dockerfile)
- Then build the image: `$ docker build -t go-dumper .`
- Then run the container: `$ docker run go-dumper`

## Via configuration yaml file
- You can fullfill your credentials at [Config File](https://github.com/Daniellunsc/go-dumper/blob/master/config.yml)

- On project folder, run `$ go get ./...`
- Then run `$ go run config.go databasehelpers.go main.go`