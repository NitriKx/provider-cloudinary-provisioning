# Stage 1: Build the wrapped Terraform provider binary from source.
# This runs on the build platform but cross-compiles for the target.
FROM --platform=$BUILDPLATFORM golang:1.26-alpine AS tf-provider-builder
ARG TARGETARCH
ARG TARGETOS

RUN apk --no-cache add git

# Cross-compile the TF provider binary.
# When GOARCH != GOHOSTARCH, go install places the binary in
# $GOPATH/bin/${GOOS}_${GOARCH}/ rather than $GOPATH/bin/.
RUN CGO_ENABLED=0 GOARCH=${TARGETARCH} GOOS=${TARGETOS} \
    go install github.com/NitriKx/terraform-provider-cloudinary-provisioning@main && \
    ( mv /go/bin/${TARGETOS}_${TARGETARCH}/terraform-provider-cloudinary-provisioning /tf-provider 2>/dev/null || \
      mv /go/bin/terraform-provider-cloudinary-provisioning /tf-provider )

# Stage 2: Runtime image (native gRPC mode — no terraform CLI required).
FROM alpine:3.23.4
RUN apk --no-cache add ca-certificates bash

ARG TARGETOS
ARG TARGETARCH
ARG TERRAFORM_NATIVE_PROVIDER_BINARY
ARG TERRAFORM_PROVIDER_VERSION

ENV USER_ID=65532

# Crossplane provider controller binary
ADD "bin/${TARGETOS}_${TARGETARCH}/provider" /usr/local/bin/provider

# Wrapped TF provider binary — the controller launches it as a native gRPC subprocess.
# TERRAFORM_NATIVE_PROVIDER_PATH tells the controller where to find it at runtime.
COPY --from=tf-provider-builder /tf-provider /usr/local/bin/${TERRAFORM_NATIVE_PROVIDER_BINARY}

ENV TERRAFORM_NATIVE_PROVIDER_PATH=/usr/local/bin/${TERRAFORM_NATIVE_PROVIDER_BINARY}
ENV TERRAFORM_PROVIDER_VERSION=${TERRAFORM_PROVIDER_VERSION}

RUN chown -R ${USER_ID}:${USER_ID} /usr/local/bin/provider \
  && chown -R ${USER_ID}:${USER_ID} /usr/local/bin/${TERRAFORM_NATIVE_PROVIDER_BINARY}

USER ${USER_ID}
EXPOSE 8080

ENTRYPOINT ["provider"]
