# Central User Management

Application to make user management easier across third-party services like:

* LDAP
* AWS
* Kubernetes etc.

## Requirements

This app is written in Go 1.18 and requires access to LDAP server the very least.

## Supported backends

The app supports backends for storing User and Group details in-memory, PostgreSQL or Redis.
