# Pulumi Jetstream Provider

Generated using the Pulumi Terraform Bridge against [terraform-provider-jetstream](https://github.com/Triple-Whale/terraform-provider-jetstream).

### Dev

```
<repo base> make install
<example dir> export TF_LOG=WARN
<example dir> rm .terraform.lock.hcl
<example dir> terraform init
<example dir> terraform apply
```

### Publish

```
commit changes
git tag v1.X
git push
git push --tags
```

## pulumi-jetstream

### Build

```
# one time
make prepare NAME=jetstream REPOSITORY=github.com/Triple-Whale/pulumi-jetstream ORG=triplewhale

in go.mod, update to v1.X
cd provider && go mod tidy && cd -
make tfgen && make provider && make build_sdks
commit
git tag v0.0.Y
git push
git push --tags
```

### Publish provider binary

```
goreleaser build --skip=validate
cd dist/pulumi-jetstream_darwin_amd64_v1/
tar -zcvf pulumi-resource-jetstream-v0.0.1-darwin-arm64.tar.gz pulumi-resource-jetstream
gsutil cp pulumi-resource-jetstream-v0.0.1-darwin-arm64.tar.gz gs://pulumi-shofifi/jetstream/
cd ../..
cd dist/pulumi-jetstream_linux_amd64_v1
tar -zcvf pulumi-resource-jetstream-v0.0.1-linux-amd64.tar.gz pulumi-resource-jetstream
gsutil cp pulumi-resource-jetstream-v0.0.1-linux-amd64.tar.gz gs://pulumi-shofifi/jetstream/
cd ../..
```

### Publish npm package

```
npm config -g set @tw:registry https://us-central1-npm.pkg.dev/shofifi/npm-packages/
cd sdk/nodejs/bin
open package.json, fix package name (@tw/pulumi-jetstream) and version (whatever you set above)
tw auth && npm publish
```

## triplewhale/backend-packages/packages/pulumi

```
tw i @tw/pulumi-jetstream@latest
...
```
