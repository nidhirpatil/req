# req

A curl wrapper with named profiles. Save HTTP requests by name and replay them with a single command.

## Build

```bash
go build -o req .
```

## Usage

### Save a profile

```bash
# Simple GET
./req save github-user --url https://api.github.com/users/octocat

# POST with body and headers
./req save create-post \
  --method POST \
  --url https://jsonplaceholder.typicode.com/posts \
  --header "Content-Type:application/json" \
  --body '{"title":"foo","body":"bar","userId":1}'

# With Bearer auth
./req save my-api --url https://api.example.com/data --bearer mytoken123

# With Basic auth
./req save admin --url https://api.example.com/admin --user admin --pass secret
```

### Run a profile

```bash
./req run github-user
./req run create-post
```

### Manage profiles

```bash
./req list            # list all saved profiles
./req show github-user  # show a profile's config as JSON
./req delete github-user  # delete a profile
```

## Profiles stored at

```
~/.config/req/profiles.json
```

## Flags for `req save`

| Flag | Description |
|---|---|
| `-u`, `--url` | Request URL (required) |
| `-X`, `--method` | HTTP method (default: GET) |
| `-H`, `--header` | Header in `Key:Value` format (repeatable) |
| `-d`, `--body` | Request body |
| `--bearer` | Bearer token |
| `--user` / `--pass` | Basic auth credentials |
