# Makefile

commit-all:
	@echo "Enter commit message: "
	@read -r msg; \
	git add .; \
	git commit -m "$$msg"
