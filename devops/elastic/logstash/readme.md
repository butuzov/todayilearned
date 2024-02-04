# Logstash

Logstash is an open source data collection engine with real-time pipelining capabilities. Logstash can dynamically unify data from disparate sources and normalize the data into destinations of your choice. Cleanse and democratize all your data for diverse advanced downstream analytics and visualization use cases.



## `logstash.yml`

[reference](https://www.elastic.co/guide/en/logstash/current/logstash-settings-file.html)

```yaml
http.host: "0.0.0.0"

```

## Configuration

### grok

- https://logz.io/blog/logstash-grok/
- https://github.com/hpcugent/logstash-patterns/blob/master/files/grok-patterns
- https://grokdebugger.com/


## `pipelines.yml`

We can have [multiple pipelines](https://www.elastic.co/guide/en/logstash/current/multiple-pipelines.html)

```yaml
# Syntax & Values
# https://github.com/elastic/logstash/blob/main/config/pipelines.yml
- pipeline.id: nifi-app-logs
  config.string: |
    # Input Path dependds on nifi logback.xml
    # Pattern depends on <encoder><pattern>%date %level [%thread] %logger{40} %msg%n</pattern></encoder> @ logback.xml
    input {
      file {
        path => "/logs/nifi-{app,bootstrap}*.log"
        start_position => "beginning"
        codec => multiline {
          pattern => "^\d{4}"
          negate => "true"
          what => "previous"
        }
      }
    }
    # Main Filter
    filter{
      # append, convert, date, geoip, grok, gsub, json, lowercase, rename, set
      grok{
        match => {
          "message" => "(?<datetime>\d{4}-\d{2}-\d{2} %{TIME}) %{LOGLEVEL:level} \[%{DATA:thread}\] %{DATA:logger} %{GREEDYDATA:eventText}"
        }
      }
      mutate{
        rename => {
          "[host][name]" => "host_name"
         "[log][file][path]" => "log_file_path"
         "eventText"         => "event"
        }
        remove_field => [ "@version", "log", "host" ]
      }

      date {
        match => [ "datetime", "yyyy-MM-dd HH:mm:ss,SSS" ]
      }
    }

    # Sends to Elastic Search
    output {
      stdout {
        codec => rubydebug
      }
      elasticsearch {
        index => "nifi-app-%{+YYYY.MM.dd}"
        hosts => "${ELASTICSEARCH_HOSTS}"
      }
    }
```

### Debuging

1. Store `rubydebug`

  ```ruby
    stdout {
      codec => rubydebug
      path => "/tmp/foo.txt"
    }
  ```

2. Ruby Code debug outside of pipeline.
3. Test with

```shell
docker-compose run logstash logstash < data/Current_FY_Cases.csv
```
```
input {
  stdin {}
}
output {
  stdout {}
}
```
