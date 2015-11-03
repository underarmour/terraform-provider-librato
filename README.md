
provider "librato" {
  user = "pbeam%40underarmour.com"
  token = "e9d083aac74151d4ae23fa4a009921ee9a64686b7dc3412076ded9a3552a4521"
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

