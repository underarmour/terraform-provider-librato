Example:
```
provider "librato" {
  user = "your@email.com"
  token = "yourlibratotoken"
}

resource "librato_space" "space" {
  name = "Identity Prod"
}

resource "librato_space_chart" "chart" {
  space = "${librato_space.space.id}"
  name = "DynamoDB Read Capacity"
  min = 0
  label = "Units"

  stream {
    name = "Consumed Read Capacity"
    metric = "AWS.DynamoDB.ConsumedReadCapacityUnits"
    source = "*identity_prod_*"
    group_function = "breakout"
    summary_function = "sum"
    units_short = "U"
    units_long = "Units"
    min = 0
    transform_function = "x/p"
    period = 60
  }
}

resource "librato_service" "identity" {
  title = "PagerDuty Identity"
  type = "pagerduty"
  settings {
    service_key = "yourpagerdutytoken"
    event_type = "trigger"
    description = "PagerDuty Identity"
  }
}

resource "librato_alert" "alert" {
  name = "identity.important.alert"
  description = "The important alert"
  rearm_seconds = 60
  services = ["${librato_service.identity.id}"]
  conditions {
    condition_type = "above"
    metric_name = "AWS.DynamoDB.ConsumedReadCapacityUnits"
    threshold = 1
    source = "*"
    summary_function = "sum"
    duration = 60
  }
}
```

