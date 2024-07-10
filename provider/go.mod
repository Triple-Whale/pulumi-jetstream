module github.com/Triple-Whale/pulumi-jetstream/provider

go 1.21

replace github.com/hashicorp/terraform-plugin-sdk/v2 => github.com/pulumi/terraform-plugin-sdk/v2 v2.0.0-20240229143312-4f60ee4e2975

require (
	github.com/nats-io/terraform-provider-jetstream v0.1.1
	github.com/pulumi/pulumi-terraform-bridge/v3 v3.80.0
	github.com/pulumi/pulumi/sdk/v3 v3.113.0
)

require github.com/hashicorp/terraform-exec v0.15.0 // indirect
