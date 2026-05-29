# Release Notes

## v0.1.0 (Initial Release)

First release of provider-cloudinary-provisioning.

### New Managed Resources

- **AccessKey** (`accesskey.cloudinaryprovisioning.crossplane.io/v1alpha1`) —
  Manage Cloudinary API access keys scoped to a product environment. The generated
  `api_key` and `api_secret` are exported to a Kubernetes Secret via
  `spec.writeConnectionSecretToRef`.

- **CustomPolicy** (`policy.cloudinaryprovisioning.crossplane.io/v1alpha1`) —
  Manage Cloudinary custom permissions policies written in Cedar policy language.
  Supports account-wide and per-product-environment scope.

- **PrincipalRoleAssignment** (`iam.cloudinaryprovisioning.crossplane.io/v1alpha1`) —
  Assign a role to a principal (API key, user, user group, or provisioning key)
  within a given scope (account or product environment).
