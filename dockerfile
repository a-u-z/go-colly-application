# base go image
FROM golang:1.19-rc-alpine

# 創建一個叫做 app 的資料夾
RUN mkdir /app

# 將 dockerfile 這層的所有東西，複製到 images 裡面的這個路徑下
COPY . /app

# 切換到在此 images 空間中的資料夾
WORKDIR /app

# RUN apk add --no-cache go

# 打包成二進制檔案 : build -o <檔名> <位置>
RUN go build -o /app/colly-application ./...