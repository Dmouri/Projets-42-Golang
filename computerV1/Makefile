NAME=computorv1

SRC_PATH=src/*.go

all: $(NAME)

$(NAME): $(SRC_PATH)
	@go build -o $(NAME) $(SRC_PATH)
	@echo "$(NAME) Built successfully"

fclean:
	@rm -rf $(NAME)
	@echo "$(NAME) removed successfully"

clean:
	@rm -rf $(NAME)
	@echo "$(NAME) removed successfully"

re: clean all

.PHONY: all fclean clean re