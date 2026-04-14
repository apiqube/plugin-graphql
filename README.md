# plugin-graphql

> GraphQL executor plugin for [ApiQube](https://github.com/apiqube).

[![License](https://img.shields.io/badge/License-MIT-blue?style=flat-square)](LICENSE)
[![Status](https://img.shields.io/badge/Status-Scaffold-yellow?style=flat-square)]()

Handles GraphQL queries and mutations. Detected by presence of `query:` or `mutation:`
fields in a test case. GraphQL runs over HTTP, so the target is an `http://` URL.

## Manifest Fields

| Field            | Type   | Required | Description |
|------------------|--------|----------|-------------|
| `query`          | string | no       | GraphQL query (use this OR `mutation`) |
| `mutation`       | string | no       | GraphQL mutation (use this OR `query`) |
| `variables`      | map    | no       | Variables passed with the operation |
| `operation_name` | string | no       | Named operation when multiple exist |
| `fragments`      | array  | no       | Named fragments referenced by the query |

## Example

```yaml
target: http://localhost:4000/graphql

tests:
  - name: Fetch user
    query: |
      query GetUser($id: ID!) {
        user(id: $id) { name email }
      }
    variables: { id: "1" }
    expect:
      status: 200
      body:
        data.user.name: exists
```

## License

[MIT](LICENSE)
