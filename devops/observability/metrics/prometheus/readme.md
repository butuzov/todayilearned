# Prometheus

**Features**

- Multi-dementional.
- Simple metrics.
- Pulls metrics over HTTP(S)?.
- TSDB.
- Supports service discovery (via `dns`) and static configuration (`hardcoded targets`).
- Has rich api libraries (e.g. for python, nodejs or golang).
- Primary data source for Grafana.

## Reading

- https://prometheus.io/docs/
- [Prometheus Up and Running](https://www.oreilly.com/library/view/prometheus-up/9781492034131/)
- https://iximiuz.com/en/posts/prometheus-metrics-labels-time-series/
- https://www.youtube.com/watch?v=h4Sl21AKiDg
- https://www.robustperception.io/how-to-unit-test-prometheus-instrumentation/


## Local Testing Stack

{{% list "docker-compose.yaml,prometheus.yml" %}}

## Metrics Format

```
# Prometheus metrics format
<metric name>{<label name>=<label value>, ...} <value>
# Example
http_requests_total{job="apiserver", handler="/api/comments"} 10
```

- Metric - `http_requests_total`
- Meta / Labeles - `job="apiserver", handler="/api/comments"`
- Value - `10`

## Types

Metric types are usually in the metric # TYPE annotation

- `counter` represents any value that goes only up (http_request_total, etc). `counters` are useful to monitor rate of an event, like rps.
- `histogram` measures a frequency of an event that falls into specific predefined buckets, like request counts by response status codes (`request_duration{code="200"} 32`)
- `summary` is similar to `histogram`, but don't require pre-defined buckets, that might be useful if you want to use quantiles, but not sure in ranges you need to use.
- `gauge` is similar to `summary` and provides a bucket of metric values. Gauges are useful to observe some metrics can go up and down, like memory allocations (`go_memstats_alloc_bytes`).

### Labels

Label names starting with `__` (such as `__name__`) are reserved for internal usage.

Labels in prometheus are divided into 2 types:



- __Instrumentation labels__ come from an instrumented application (type of HTTP request, etc).

- __Target labels__ identify a specific monitoring target, and related more to an architecture and infrastructure. For example, different teams may have different vision of what a "team", "region", or "service" is, so instrumented app shouldn't expose this labels itself rather than leave it to relabeling feature. Labels most likely used as target labels: env, cluster, service, team, zone, and region.



### Annotations

```
# HELP latency_seconds Latency in seconds.
# TYPE latency_seconds summary
latency_seconds_sum{path="/foo"} 1.0
```

## Naming

- https://prometheus.io/docs/practices/naming/
- https://www.robustperception.io/on-the-naming-of-things/

### Sufixes

- `_totalis` a counter
- `_count` is a counter
- `_sum` stands for summary
- `_bucket` is for a histogram


## PromQL


```shell
# selecting by metric name
> nifi_stor_used_bytes
nifi_stor_used_bytes{env="testing", instance="nifi-export:9103", job="nifi-export", location="default", node_id="aggregate", type="content"}
	37968252928
nifi_stor_used_bytes{env="testing", instance="nifi-export:9103", job="nifi-export", location="default", node_id="aggregate", type="flow_file"}
	37968252928
nifi_stor_used_bytes{env="testing", instance="nifi-export:9103", job="nifi-export", location="default", node_id="aggregate", type="provenance"}

# aggregate by ignore type
> sum without(type)(nifi_amount_bytes_transferreds)
{env="testing", instance="nifi-export:9103", job="nifi-export", location="default", node_id="aggregate"} 113900236800
> avg without(type)(nifi_stor_used_bytes)
{env="testing", instance="nifi-export:9103", job="nifi-export", location="default", node_id="aggregate"} 	37967511552


# average per second for lat 5 minutes
> rate(nifi_amount_bytes_transferred[5m])
{component_id="1e8d6df2-d134-37e2-817a-88dc311642d3", component_name="default", component_type="ProcessGroup", exported_instance="nifi-node1", instance="nifi-node1:9092", job="nifi-analitics", parent_id="457d5c1e-018d-1000-1c9f-1eac4f5e00d1"}
	17180.74216684518
{component_id="329b6cc7-bdcf-3041-ba4b-795d6edda52e", component_name="blackbox", component_type="ProcessGroup", exported_instance="nifi-node1", instance="nifi-node1:9092", job="nifi-analitics", parent_id="1e8d6df2-d134-37e2-817a-88dc311642d3"}
	6872.296866738071
{component_id="457d5c1e-018d-1000-1c9f-1eac4f5e00d1", component_name="NiFi Flow", component_type="RootProcessGroup", exported_instance="nifi-node1", instance="nifi-node1:9092", job="nifi-analitics"}
	17180.74216684518
```


PromQL Matcher Operators

- `=`  - equality matcher
- `!=` - negative equality matcher
- `=~` - regexp matcher
- `!~` - negative regexp matcher

Notes:
- Matchers are fully anchored.
- Regexp matchers smell. If you're using them a lot, for example to match HTTP response codes like code~="4..", consider to combine this labels into 4xx to use with equality matcher `code="4xx"
- At least one matcher must not match the empty string, so {}, {foo=""}, {foo!=""}, and {foo=~".*"} returns an error.



- https://prometheus.io/docs/prometheus/latest/querying/examples/
- [regular expressions 2 format](https://github.com/google/re2/wiki/Syntax)

```promql
# Return all TS with the metric http_requests_total:
http_requests_total
# Return all TS with the metric http_requests_total and the given job and labels:
http_requests_total{job="apiserver", handler="/api/comments"}
# 5 minutes vector
http_requests_total{job="apiserver", handler="/api/comments"}[5m]
# Using regular expressions, you could select time series only for
# jobs whose name match a certain pattern, in this case, all jobs that end with server:
http_requests_total{job=~".*server"}
# To select all HTTP status codes except 4xx ones, you could run:
http_requests_total{status!~"4.."}
#
```

### Subqueries

```promql
# Return the 5-minute rate of the http_requests_total metric for the past 30 minutes,
# with a resolution of 1 minute.
rate(http_requests_total[5m])[30m:1m]
# The subquery for the deriv function uses the default resolution.
# Note that using subqueries unnecessarily is unwise.
max_over_time(deriv(rate(distance_covered_total[5s])[30s:5s])[10m:])
```

### Using functions, operators, etc.

```promql
# Return the per-second rate for all time series with the http_requests_total
# metric name, as measured over the last 5 minutes:
rate(http_requests_total[5m])
```

Assuming that the `http_requests_total`` time series all have the labels job (fanout by `job`` name) and instance (fanout by instance of the job), we might want to sum over the rate of all instances, so we get fewer output time series, but still preserve the job dimension:

```promql
sum by (job) (
  rate(http_requests_total[5m])
)
```

### PromQL Aggregation Operators


- `sum` (metric) returns a sum aggregated by same labels.
- `max` (metric) returns max value from a gauge.
- `avg` (metric) returns average value for a metric
- `/`   the division operator matches the timeseries with the same labels.

There is a without function in case you need to omit specific labels from an aggregation (sum without(label, label)).

## Push vs Pull

### `pushgateway`

You can use [Pushgateway](https://prometheus.io/docs/practices/pushing/) if you have reasons to use it.

- When monitoring multiple instances through a single Pushgateway, the Pushgateway becomes both a single point of failure and a potential bottleneck.
- You lose Prometheus's automatic instance health monitoring via the up metric (generated on every scrape).
- The Pushgateway never forgets series pushed to it and will expose them to Prometheus forever unless those series are manually deleted via the Pushgateway's API.

### Exporters

- https://github.com/prometheus/node_exporter
- https://github.com/google/cadvisor
- https://exporterhub.io/
- https://prometheus.io/docs/instrumenting/exporters/
- some applications will export metrics from the box

## Service Discovery

### Consul

```yaml
global:
  scrape_interval: 15s
  scrape_timeout: 15s
scrape_configs:
- job_name: 'consul-prometheus'
  consul_sd_configs:
    - server: 'consul:8500'
      services: ['foobar-service']
```

### DNS

```yaml
global:
  scrape_interval:     5s

scrape_configs:
  - job_name: 'promtail'
    dns_sd_configs:
      - names: [ promtail ]
        type: A
        port: 9080
```

### Using Files

[`file_sd_config`](https://prometheus.io/docs/prometheus/latest/configuration/configuration/#file_sd_config)

```yaml
global:
  evaluation_interval: 1m
  scrape_interval: 30s
  scrape_timeout: 10s
remote_write:
  - url: http://localhost:8080/workspaces/WORKSPACE/api/v1/remote_write
scrape_configs:
  - job_name: ecs_services
    file_sd_configs:
      - files:
          - /etc/config/ecs-services.json
        refresh_interval: 30s
```


