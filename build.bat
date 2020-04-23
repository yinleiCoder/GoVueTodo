rd /s/q release
md release
go build -o govuetodo.exe
COPY govuetodo.exe release\
COPY favicon.ico release\favicon.ico

XCOPY static\*.* release\static\  /s /e
XCOPY templates\*.* release\templates\  /s /e