set GOOS=linux
set GOARCH=amd64
rm -rf ./release
mkdir release
go build -o govuetodo
chmod +x ./govuetodo
cp govuetodo ./release/
cp favicon.ico ./release/
cp -arf ./static ./release/
cp -arf ./templates ./release/