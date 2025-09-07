.PHONY: frontend-dev 

# Default target
help:
	@echo "Available commands:"
	@echo "  frontend-dev      - Start frontend development server"
frontend-dev:
	cd apps/frontend && npm run dev