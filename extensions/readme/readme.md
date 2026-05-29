# provider-cloudinary-provisioning

Crossplane provider for [Cloudinary](https://cloudinary.com/) that manages
account-level resources via the Cloudinary **Provisioning API**.

Use this provider alongside
[provider-cloudinary](https://github.com/NitriKx/provider-cloudinary)
to implement a full GitOps workflow for Cloudinary resource management.

## Supported Resources

| Kind | Description |
|---|---|
| `AccessKey` | API key + secret pair for a product environment; credentials exported to a Kubernetes Secret |
| `CustomPolicy` | Cedar-language permissions policy scoped to account or product environment |
| `PrincipalRoleAssignment` | Role assignment binding a key, user, or group to a policy |

## Installation

```bash
kubectl crossplane install provider ghcr.io/nitrikx/provider-cloudinary-provisioning:latest
```

## Credentials

The Provisioning API uses account-level credentials (different from per-cloud
Admin API keys). Retrieve them from the Cloudinary console under
Account -> Provisioning API access.

```bash
kubectl create secret generic cloudinary-provisioning-creds -n crossplane-system \
  --from-literal=credentials='{"account_id":"ACCOUNT_ID","api_key":"PROV_KEY","api_secret":"PROV_SECRET"}'
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

## Example: Create an AccessKey

The AccessKey's generated api_key and api_secret are written to a
Kubernetes Secret for use by applications.

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
