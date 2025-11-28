help:
	@echo "Avaliable targets:"
	@echo "up			- Start everything"
	@echo "down			- Stop everything"
	@echo "restart		- Restart services, but not containers"
	@echo "test			- Test everything"
	@echo "clean		- Clean everything"

up:
	docker compose up -d
	@echo "Waiting containers to be ready..."
	@sleep 5
	$(MAKE) -C backend start

down:
	$(MAKE) -C backend stop || true
	docker compose down

restart:
	$(MAKE) -C backend restart

test:
	$(MAKE) -C backend test

clean:
	$(MAKE) -C backend clean
	docker compose down --volumes
