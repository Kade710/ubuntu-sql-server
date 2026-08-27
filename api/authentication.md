# API Authentication

## Status

API authentication has not been implemented yet.

This document will describe how clients authenticate with the Ubuntu SQL Server API once the API is built.

## Planned Goals

Authentication should:

- Protect API access
- Prevent unauthorized requests
- Avoid storing passwords or tokens in source code
- Use environment variables or another protected secret store
- Support role-based or permission-based access when needed

## Security Requirements

When authentication is added:

- Secrets must not be committed to Git
- Tokens should be treated as sensitive credentials
- Authentication failures should be logged without exposing secrets
- API access should follow the principle of least privilege
- HTTPS should be used if the API is exposed outside a trusted local network

## Future Options

Possible authentication methods may include:

- API keys
- Bearer tokens
- Session-based authentication
- Token-based authentication

The final method will be documented after it is implemented and tested.
