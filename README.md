# Facebook API Prometheus Exporter

Prometheus metrics exporter for tracking the status of the Facebook Platform every 60 seconds from https://www.facebook.com/platform/api-status/


# Metrics Exposed 
facebook_status 1
Value will be 1 if platform is healthy or 2 if not

# Building  
install the [Go Compiler!](https://golang.org/dl)
run go build
copy the ./facebook_api_status_exporter binary to where you need it