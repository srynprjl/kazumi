app_name = kazumi
version = alpha
locs = /home/${USER}/.local/bin


build:
	@echo "Building ${app_name} ${version}"
	@mkdir -p ${locs}
	@CGO_ENABLED=0 go build -ldflags="-s -w -extldflags '-static'" -trimpath -o kazumi cmd/kazumi/main.go
	@echo "Completed..."

clean:
	@echo "Removing ${app_name} ${version}"
	@mkdir -p ${locs}
	@rm kazumi
	@echo "Completed..."

install:
	@echo "Installing ${app_name} ${version}"
	@mkdir -p ${locs}
	@CGO_ENABLED=0 go build -ldflags="-s -w -extldflags '-static'" -trimpath -o ${locs}/kazumi cmd/kazumi/main.go
	@echo "Completed..."

uninstall:
	@echo "Uninstalling ${app_name} from system"
	@rm ${locs}/kazumi
	@echo "Completed..."
