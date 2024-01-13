test:
	@echo 'Мы сделали Makefile'

up:
	sudo docker-compose up --build home

stop:
	sudo docker-compose stop