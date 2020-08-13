# Integration Tests

This directory contains integration tests that push the asc-go library outside the scope of the unit tests or the examples projects. These integration tests will make mutating calls against a production App Store Connect team, are invoked manually, and take a while to run. The unit tests by comparison are fast to run, do not involve the network, and run automatically on every commit.

To run the integration tests, provide something similar to this command line invocation.

```shell
env \
      ASC_INTEGRATION_KID="..." \
      ASC_INTEGRATION_ISS="..." \
      ASC_INTEGRATION_PRIVATE_KEY_PATH="..." \
      go test -v -tags=integration ./test/integration
```

Much like the examples, the integration tests require the presence of at least 3 of 4 different environment variables. If you want to know what they do, please consult the repo's general documentation on [authentication](./README.md#Authentication):

- `ASC_INTEGRATION_KID` – key ID
- `ASC_INTEGRATION_ISS` – issuer ID
- `ASC_INTEGRATION_PRIVATE_KEY` – private key content
- `ASC_INTEGRATION_PRIVATE_KEY_PATH` – path to a private key

Only one of either `ASC_INTEGRATION_PRIVATE_KEY` or `ASC_INTEGRATION_PRIVATE_KEY_PATH` is required; if both are provided, `ASC_INTEGRATION_PRIVATE_KEY` will take precedence. Since the App Store Connect API requires an authenticated session, you must have valid credentials to run these tests.
