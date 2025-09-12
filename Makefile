.PHONY: frontend-dev backend-dev

# Default target
help:
	@echo "Available commands:"
	@echo "  frontend-dev      - Start frontend development server"
	@echo "  backend-dev       - Start backend development server"

frontend-dev:
	cd apps/frontend && npm run dev 

backend-dev:
	cd apps/backend && . ./env.sh && go run cmd/main.go