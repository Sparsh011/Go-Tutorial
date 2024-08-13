# Makefile

commit-all:
	@echo "Enter commit message: "
	@read -r msg; \
	git add .; \
	git commit -m "$$msg"

# Makefile

new:
	@read -p "Enter directory name: " dir_name; \
	read -p "Enter package name: " pkg_name; \
	mkdir $$dir_name; \
	cd $$dir_name; \
	go mod init $$pkg_name