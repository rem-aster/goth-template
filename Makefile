run:
	docker compose down
	sqlc generate
	templ generate
	npx tailwindcss -i styles.css -o web/static/tw.css
	docker compose up --build
gen:
	sqlc generate
	templ generate
	npx tailwindcss -i styles.css -o web/static/tw.css