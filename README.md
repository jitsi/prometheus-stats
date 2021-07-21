# stats
Library of common Prometheus statistics for Jitsi Go services.

These are largely based on prometehus example code.

## HTTP Server API Statistics

These metrics are used as a wrapper by HTTP servers fielding API requests. Example usage:
```
service := stats.WrapHTTPHandler("endpoint", chain.ThenFunc(handlers.Endpoint))
```

| name                                 | type      | labels                            |
|--------------------------------------|-----------|-----------------------------------|
| http_server_requests_in_flight       | guage     | instance, job                     |
| http_server_requests_total           | counter   | instance, job, code, method, uri  |
| http_server_request_duration_seconds | histogram | instance, job, method, uri        |
### metrics
* `http_server_requests_in_flight`: current number of requests being handled by the service
* `http_server_requests_total`: total request counter
* `http_server_request_duration_seconds`: histogram of request latencies

Note that the buckets for `http_server_request_duration_seconds need to be tuned
for precision if a particular issue is being analyzed.

### labels
* `instance`: the instance that is collecting stats; added by k8s
* `job`: the name of the service; added by k8s based on the name of the scrape job
* `code`: HTTP status code of the response; added by instrumentation
* `method`: HTTP method used for the request; added by instrumentation
* `uri`: actual API call used by caller; added by instrumentation

## HTTP Client RoundTripper Statistics

These metrics are used by HTTP Clients that are executing API calls. Example:
```
HTTPClient: &http.Client{
	Transport: stats.RoundTripper(),
},
```

client.HTTPClient.Transport = stats.RoundTripper()

| name                            | type      | labels        |
|---------------------------------|-----------|---------------|
| client_in_flight_requests       | guage     |               | 
| client_requests_total           | counter   | code, method  |
| client_request_duration_seconds | histogram |               | 
| dns_duration_seconds            | histogram | event         |
| tls_duration_seconds 		  | histogram | event         |

### metrics
* `client_in_flight_requests`: current number of requests being handled by this client
* `client_requests_total`: total requests made by client
* `client_request_duration_seconds`: histogram of durations of successful requests
* `dns_duration_seconds`: histogram of durations of successful DNS queries
* `tls_duration_seconds`: histogram of durations of successful TLS handshakes

### labels
* `code`: HTTP status code of the response; added by instrumentation
* `method`: HTTP method used for the request; added by instrumentation
* `event`: start/done trace state of DNS or TLS
