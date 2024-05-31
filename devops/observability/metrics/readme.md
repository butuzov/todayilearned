# Metrics

## Monitoring Methods

### Golden Signals

- https://sre.google/sre-book/monitoring-distributed-systems/

For every resource, monitor:

- Latency: the time it takes to serve a request.
- Traffic: the total number of requests across the network.
- Errors: the number of requests that fail.
- Saturation: the load on your network and servers.

### The USE Method

- https://www.brendangregg.com/usemethod.html

For every resource, monitor:

- Utilization (% time that the resource was busy)
- Saturation (amount of work resource has to do, often queue length)
- Errors (count of error events)

### The RED Method
- [The RED Method](https://www.youtube.com/watch?v=zk77VS98Em8) + [Tom_Wilkie_GrafanaCon_EU_2018.pdf](https://grafana.com/files/grafanacon_eu_2018/Tom_Wilkie_GrafanaCon_EU_2018.pdf)
- https://grafana.com/blog/2018/08/02/the-red-method-how-to-instrument-your-services/

For every resource, monitor:

- Rate (the number of requests per second)
- Errors (the number of those requests that are failing)
- Duration (the amount of time those requests take)


## Actual Metrics

### Latency (Request Service Time)

Latency is the time it takes a system to respond to a request. Both successful and failed requests have latency and it’s vital to differentiate between the latency of successful and failed requests. For example, an HTTP 500 error, triggered because of a connection loss to the database might be served fast. Although, since HTTP 500 is an error indicating failed request, factoring it into the overall latency will lead to misleading calculations. Alternatively, a slow error can be even worse as it factors in even more latency. Therefore, instead of filtering out errors altogether, keep track of the error latency. Define a target for a good latency rate and monitor the latency of successful requests against failed ones to track the system’s health.

### Traffic (User Demand)

Traffic is the measure of how much your service is in demand among users. How this is determined varies depending on the type of business you have. For a web service, traffic measurement is generally HTTP requests per second, while. In a storage system, traffic might be transactions per second or retrievals per second.

By monitoring user interaction and traffic in the service, SRE teams can usually figure out the user experience with the service and how it’s affected by shifts in the service's demand.

### Errors (Rate of Failed Requests)

Error is the rate of requests that fail in any of the following ways:

- Explicitly: for example, HTTP 500 internal server error.
- Implicitly: for example, HTTP 200 success response coupled with inaccurate content.
- By policy: for example, as your response time is set to one second, any request that takes over one second is considered an error.


### Saturation (Overall Capacity of the System)

Saturation refers to the overall capacity of the service or how “full” the service is at a given time. It signifies how much memory or CPU resources your system is utilizing. Many systems start underperforming before they reach 100% utilization. Therefore, setting a utilization target is critical as it will help ensure the service performance and availability to the users.
