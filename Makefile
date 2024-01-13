test:
	@echo 'Мы сделали Makefile'

up:
	sudo docker-compose up --build Home

stop:
	sudo docker-compose stop