# Go, GORM & Gin CRUD Example

## Install

1. Clone this repository to `$GOPATH/src/github.com/herusdianto` directory:

        git clone https://github.com/herusdianto/gorm_crud_example.git

2. Install `glide`:

        https://glide.sh/

3. CD to `gorm_crud_example` folder:

        cd $GOPATH/src/github.com/herusdianto/gorm_crud_example

4. Install dependencies using `glide`:

        glide install

5. Open `main.go` and modify this variable values:

        dbUser, dbPassword, dbName := "root", "root", "gorm_crud_example"

6. Login to `MySQL` and create the database:

        create database gorm_crud_example;

7. Run `main.go`:

        go run main.go

## Features

- [x] Database Migration
- [x] Create Data
- [x] Read All Data
- [x] Find One Data By ID
- [x] Update Data
- [x] Delete One Data By ID
- [x] Delete Multiple Data By IDs
- [x] Sort & Paginate Data
- [ ] Search Data

If you want to watch step by step I'm making this, you can watch this [videos](https://www.youtube.com/playlist?list=PLKmlCa2HUPq-K7hIyHGbDoYs6YZBM8yA-).

Support me with subscribe to my [channel](https://www.youtube.com/channel/UCpKERrPCRQBFaTnJu4xWn5A) on youtube, thank you.