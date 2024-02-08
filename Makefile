build:
	cd bbname-front && npm run build && rsync -rlv --delete build/ ../static
	cd ..
	docker-compose build

