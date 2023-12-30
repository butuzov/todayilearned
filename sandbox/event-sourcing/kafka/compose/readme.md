<!-- menu: Local Testing -->
# Different Kafka Setups

## Regular Setup

- 2 different clients
- Zookeper & Kafka Cluster

{{% list "docker-compose.yaml" %}}

## No Zookeeper's Cluster

- akhq client
- no zookipers!
- 3 repeated kafka insances
- schmea registry

{{% list "docker-compose-nozoo.yaml" %}}
