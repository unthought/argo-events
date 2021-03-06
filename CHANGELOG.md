# Changelog

## v0.6
+ Gitlab Gateway #120
+ If sensor is repeatable then deploy it as deployment else job #109
+ Start gateway containers in correct order. Gateway transformer > gateway processor. Add readiness probe to gateway transformer #106
+ Kubernetes configmaps as artifact locations for triggers #104
+ Let user set extra environment variable for sensor pod #103 
+ Ability to provide complete Pod spec in gateway.
+ The schedule for calendar gateway is now in standard cron format   #102
+ FileWatcher as core gateway #98
+ Tutorials on setting up pipelines #105
+ Asciinema recording for setup #107
+ StorageGrid Gateway
+ Scripts to generate swagger specs from gateway structs
+ Added support to pass non JSON payload to trigger
+ Update shopify's sarama to 1.19.0 to support Kafka 2.0


## v0.5
[#92](https://github.com/argoproj/argo-events/pull/92)
+ Introduced gateways as event generators. 
+ Added multiple flavors of gateway - core gateways, gRPC gateways, HTTP gateways, custom gateways
+ Added K8 events to capture gateway configurations update and as means to update gateway resource
+ SLA violations are now reported through k8 events
+ Sensors can trigger Argo workflow, any kubernetes resource and gateway
+ Gateway can send events to other gateways and sensors
+ Added examples for gateway and sensors
+ Sensors are now repeatable and fixed all issues with signal repeatability.
+ Removed signal deployments as microservices.

## v0.5-beta1 (tbd)
+ Signals as separate deployments [#49](https://github.com/argoproj/argo-events/pull/49)
+ Fixed code-gen bug [#46](https://github.com/argoproj/argo-events/issues/46)
+ Filters for signals [#26](https://github.com/argoproj/argo-events/issues/26)
+ Inline, file and url sources for trigger workflows [#41](https://github.com/argoproj/argo-events/issues/41)

## v0.5-alpha1
+ Initial release
