#Go Web Server

A simple Go web server created for learning Go.

To install:

1. vagrant up
2. vagrant ssh
3. "go run ./src/webserver.go"
4. Open your browser and navigate to localhost:8080/index.html

You can edit the config.json file to change the webroot and server port.

Note: Ensure your GOPATH is set to /var/apps/gowebserver. You can run "echo $GOPATH" to get the current version and run "export GOPATH=/var/apps/gowebserver" to change it, if needed.

## Testing

To run the test suit:

1. From the root dir run "go test ./..."

You should get some ouput as such:

```
vagrant@vagrant-ubuntu-trusty-64:/var/apps/gowebserver$ go test ./...
?   	_/var/apps/gowebserver/src	[no test files]
ok  	util/reader	0.019s
```
