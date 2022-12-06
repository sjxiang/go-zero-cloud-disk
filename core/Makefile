
TARGET = core
OBJ = core.go


build: $(OBJ)
	go build -o $(TARGET) $(OBJ)

run: build
	./$(TARGET)

bare:
	go run $(OBJ)


# 防止 touch clean 歧义
.PHONY: clean

clean:
	rm -f $(TARGET)