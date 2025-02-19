
ifeq ($(OS),Windows_NT)
    DISCORD_ASAR := $(APPDATA)/discord/app.asar
else
    UNAME_S := $(shell uname -s)
    ifeq ($(UNAME_S),Linux)
        DISCORD_ASAR := $(HOME)/.config/discord/app.asar
    endif
    ifeq ($(UNAME_S),Darwin)
        DISCORD_ASAR := /Applications/Discord.app/Contents/Resources/app.asar
    endif
endif

backup:
	if [ -f ./app.asar ]; then \
		echo "Backup already exists"; \
	else \
		cp $(DISCORD_ASAR) ./app.asar; \
	fi

revert:
	if [ -f ./app.asar ]; then \
		cp ./app.asar $(DISCORD_ASAR); \
	else \
		echo "Backup does not exist"; \
	fi

extract: backup
	mkdir -p ./discord-source
	npx asar extract "$(DISCORD_ASAR)" ./discord-source

clean:
	rm -f ./app.asar
	rm -rf ./discord-source
	go mod tidy

build:
	go build -o ./bin/bouquet ./cmd/cli/main.go
