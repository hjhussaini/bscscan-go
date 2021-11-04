build:
	@docker build -t bscscan-client .

run:
	@docker run --rm -d --name bscscan-client -p 8088:8080 bscscan-client

stop:
	@docker stop bscscan-client

test:
	curl http://localhost:8088/api/contracts/0x1da200f724b6e707cD8B8593f2c270771B7FC769

clean:
	@docker rmi bscscan-client
