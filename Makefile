#出力先ディレクトリ
BINDIR := Tools
SRC := main.go
GOOS := linux
GOARCH := amd64
#ビルドする対象のディレクトリ
TARGET := filerename
#FileName := tools
#画像からExif情報を取り出してCsｖで出力するツール
.PHONY: mac
mac:
	$(eval GOOS := darwin)

.PHONY: windows
windows:
	$(eval GOOS=windows )

.PHONY: linux
linux:
	$(eval GOOS=linux )

.PHONY: exif_csv
exif_csv:
	$(eval TARGET=exif_csv )

.PHONY: file_move
file_move:
	$(eval TARGET=file_move )

.PHONY: filerename
filerename:
	$(eval TARGET=filerename )

.PHONY: spec
spec:
	$(eval TARGET=spec)


build_exe:
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o $(BINDIR)/$(TARGET).exe $(TARGET)/$(SRC)

build:
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o $(BINDIR)/$(TARGET) $(TARGET)/$(SRC)
clean:
	@rm -rf ./$(BINDIR)/*

task:
	echo "fdsajljk" >> $(BINDIR)/a.txt
	echo "fdsajljk" >> $(BINDIR)/a.txt
	export A="afdas"
	echo $(AA)
	echo $(GOOS)
task2:
	echo "task2" >> $(BINDIR)/a.txt
test:
	@echo ${GOOS}
	@echo $(GOARCH)
	@echo ${TARGET}
