# Grafana

## Setup

{{% list "docker-compose.yaml,datasources.yaml,dashboard.yaml,prometheus.yml" %}}

## Mac

```shell
brew install grafana

# then
/opt/homebrew/opt/grafana/bin/grafana server --config /opt/homebrew/etc/grafana/grafana.ini --homepath /opt/homebrew/opt/grafana/share/grafana --packaging\=brew cfg:default.paths.logs\=/opt/homebrew/var/log/grafana cfg:default.paths.data\=/opt/homebrew/var/lib/grafana cfg:default.paths.plugins\=/opt/homebrew/var/lib/grafana/plugins

# or
brew services start grafana
```

## Dashboards

### Variables

Open __Dashboard Settings__, select __Variables__, and click __Add Variable__. Select the __Query variable__ type and use these settings:

- Name:  "Var Name"
- Label: "Var"

And then templating:

__Dashboard__ `->` __Variable__ `->` __Add Variable__


## Data Sources

### Elasticsearch

- `{{metric}}` metric name (e.g. `Count` for `count`)
- `{{term fieldname}}`
- `{{field}}`
