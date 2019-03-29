# stats
Prometheus statistic helper lib for Go

## HTTP Server API Statistics

| name                                 | type      | labels                             |
|--------------------------------------|-----------|------------------------------------|
| http_server_requests_in_flight       | guage     | instance, job                      |
| http_server_requests_total           | counter   | instance, job, code, method, uri |
| http_server_request_duration_seconds | histogram | instance, job, method, uri |

### Labels
#### instance
Added automatically by kube and is the instance that is statting.
#### job
This is the name of the scrape job in the prometheus config and should be set the
same as the name of the service. Querying for job should provide all the stat of that type for the service. This is added automatically.
#### code
The http status code of the response. This is used for alerting on specific code counts. The instrumentation code must add this label.
#### method
This is the http method used for the request. The instrumentation code must add this label.
#### uri
This is the actual api used and is added by instrumentation code

### http_server_requests_in_flight
This stat shows the current number of requests in flight for a given job/service.

### http_server_requests_total
This stat is used to show the total number of requests.

### http_server_request_duration_seconds
This stat is used to provide a histogram of request latency. Note that the buckets need to be tuned if more precision is required from the histogram. This metric, more than the other http_server* metrics need to be researched and understood by anyone using it.
