NAME=gomoku
NAME_WINDOWS=gomoku.exe

SRC_PATH=src/*.go
PPROF_PATH=profile
PATH_RELEASE=./gomoku-1.0.0
ICON_PATH_FROM_RELEASE=../assets/blackstone.ico

all: $(NAME)

windows: $(NAME_WINDOWS)

$(NAME_WINDOWS): $(SRC_PATH) create-folder
	@go build -o $(NAME_WINDOWS) $(SRC_PATH)
	@rm -rf $(PATH_RELEASE)
	@mkdir -p $(PATH_RELEASE)
	@cp -rf ./save ./assets ./sounds ./fonts ./$(NAME_WINDOWS) $(PATH_RELEASE)/
	@echo "Compil Windows OK"
	@echo "Editing ico using Resource Hacker"
	@echo "INFO: ResourceHacker must be installed and PATH set up correctly, but it's not mandatory."
	@cd ./$(PATH_RELEASE); ResourceHacker -open $(NAME_WINDOWS) -save $(NAME_WINDOWS) -action addskip -res $(ICON_PATH_FROM_RELEASE) -mask ICONGROUP,MAIN,
	@echo "Success: Icon has been modified."

$(NAME): $(SRC_PATH)
	@go build -o $(NAME) $(SRC_PATH)
	@echo "Exec $(NAME) OK"

deps: linux-deps goget 

fclean:
	@rm -rf $(NAME)
	@echo "Exec $(NAME) removed"

create-folder:
	@mkdir -p save

clean:
	@rm -rf $(NAME)
	@echo "Exec $(NAME) removed"

windows-clean:
	@rm -rf $(NAME_WINDOWS)
	@echo "Exec for windows rm"

re: clean all

windows-re: windows-clean windows

goget:
	go get -u github.com/faiface/pixel
	go get -u github.com/faiface/glhf
	go get -u github.com/faiface/beep
	go get -u github.com/hajimehoshi/oto
	go get -u github.com/hajimehoshi/go-mp3
	go get -u github.com/go-gl/glfw/v3.3/glfw
	go get -u github.com/golang/freetype/truetype

linux-deps:
	@sudo apt update
	@sudo apt install libasound2-dev libgl1-mesa-dev xorg-dev

.PHONY: all fclean clean re goget
