# provider-cloudinary-provisioning

`provider-cloudinary-provisioning` is a [Crossplane](https://crossplane.io/) provider
that manages [Cloudinary](https://cloudinary.com/) account-level resources using the
Cloudinary **Provisioning API**. It is built with
[Upjet](https://github.com/crossplane/upjet) code generation tools and exposes
XRM-conformant managed resources.

Use this provider to manage API access keys, custom permission policies, and role
assignments within a Cloudinary account. For managing resources inside a product
environment (folders, triggers) see
[provider-cloudinary](https://github.com/NitriKx/provider-cloudinary).

## Supported Resources

| Kind | API Group | Cloudinary Object |
|---|---|---|
| `AccessKey` | `accesskey.cloudinaryprovisioning.crossplane.io` | API key + secret pair for a product environment |
| `CustomPolicy` | `policy.cloudinaryprovisioning.crossplane.io` | Cedar-language permissions policy |
| `PrincipalRoleAssignment` | `iam.cloudinaryprovisioning.crossplane.io` | Role assignment to a principal (key, user, group) |

All resources are also available in a namespaced variant under the
`*.cloudinaryprovisioning.m.crossplane.io` API groups.

> **Note:** `ProductEnvironment` (sub-account) management is intentionally excluded
> from v1. Product environments are created manually in the Cloudinary console.

## Installation

Requires [Crossplane](https://docs.crossplane.io/latest/software/install/) >= v1.14.

```bash
kubectl crossplane install provider ghcr.io/nitrikx/provider-cloudinary-provisioning:latest
```

Or apply directly:

```yaml
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-cloudinary-provisioning
spec:
  package: ghcr.io/nitrikx/provider-cloudinary-provisioning:latest
```

## Credentials Setup

The Provisioning API uses **account-level** credentials (different from per-cloud
Admin API keys). Retrieve them from the Cloudinary console under
Account -> Provisioning API access.

```bash
# Option A -- account URL form
kubectl create secret generic cloudinary-provisioning-creds -n crossplane-system \
  --from-literal=credentials='{"account_url":"account://PROV_API_KEY:PROV_API_SECRET@ACCOUNT_ID"}'

# Option B -- individual fields
kubectl create secret generic cloudinary-provisioning-creds -n crossplane-system \
  --from-literal=credentials='{"account_id":"abcdef123456","api_key":"prov_key","api_secret":"prov_secret"}'
```

Then create a ProviderConfig:

```yaml
apiVersion: cloudinaryprovisioning.crossplane.io/v1beta1
kind: ProviderConfig
metadata:
  name: default
spec:
  credentials:
    source: Secret
    secretRef:
      namespace: crossplane-system
      name: cloudinary-provisioning-creds
      key: credentials
```

## AccessKey Connection Secret

When you create an AccessKey managed resource, the generated api_key and
api_secret are automatically exported to a Kubernetes Secret specified via
spec.writeConnectionSecretToRef. Applications consume this Secret directly as
Cloudinary Admin API credentials.

```yaml
apiVersion: accesskey.cloudinaryprovisioning.crossplane.io/v1alpha1
kind: AccessKey
metadata:
  name: my-app-key
spec:
  forProvider:
    productEnvironmentId: "prod-env-id"
    name: my-app-key
    enabled: true
  writeConnectionSecretToRef:
    namespace: my-app-ns
    name: my-app-cloudinary-creds
  providerConfigRef:
    name: default
```

## Quick Start

```bash
kubectl apply -f examples/cluster/providerconfig/providerconfig.yaml
kubectl apply -f examples/cluster/accesskey/v1alpha1/accesskey.yaml
kubectl get accesskey -o wide
```

## Developing

Regenerate from the local sibling Terraform provider:

```bash
make generate-local
```

Build the provider binary:

```bash
make build
```

## Report a Bug

Open an [issue](https://github.com/NitriKx/provider-cloudinary-provisioning/issues).
