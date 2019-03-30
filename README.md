### golstatsd

Statsd **log**ger written in Golang that reads the udp packets sent from a statsd client
and write them to stdout, by default it listens to port 8125 on localhost.

to install it:

`go get github.com/shadi/golstatsd`

and run it using:
`golstatsd`

or if the client is sending the metrics to a different port:
`golstatsd -port $PORT`

to write metrics to a file and run the process in the background:
`golstatsd &> $LOG_FILE &`
