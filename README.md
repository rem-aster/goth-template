# goth-template

Go-Templ-HTMX project template/example

[![xc compatible](https://xcfile.dev/badge.svg)](https://xcfile.dev)

Stack:

- Go
- Templ
- HTMX
- DaisyUI
- sqlc
- Docker

Before start install go, npm, templ and create .env like below:

```bash
POSTGRES_USER=admin
POSTGRES_PASSWORD=password
POSTGRES_DB=app
```

if you are using VSCode add `.vscode/settings.json` or change settings permanently in app

```json
{
  "[templ]": {
    "editor.formatOnSave": true,
    "editor.defaultFormatter": "a-h.templ"
  },
  "tailwindCSS.includeLanguages": {
    "templ": "html"
  },
  "emmet.includeLanguages": {
    "templ": "html"
  }
}

```

## Tasks

### run

Generate and run in docker

```bash
docker compose down
sqlc generate
templ generate
npx tailwindcss -i styles.css -o web/static/tw.css
docker compose up --build
```

### gen

Just generate all

```bash
sqlc generate
templ generate
npx tailwindcss -i styles.css -o web/static/tw.css
```
