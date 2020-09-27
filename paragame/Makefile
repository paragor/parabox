.PHONY: build

GOL_CMD_FILE='./cmd/game_of_life/main.go'
GOL_GUI_CMD_FILE='./cmd/gol_gui/main.go'
GOL_WITH_RAND_CMD_FILE='./cmd/gol_with_rand/main.go'

gol:
	go run ${GOL_CMD_FILE}

gol_with_rand:
	go run ${GOL_WITH_RAND_CMD_FILE}

gui:
	go run ${GOL_GUI_CMD_FILE}

run: gol
