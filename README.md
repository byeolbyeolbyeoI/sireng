## Introduction
Sireng is a badass API service that is designed for people who's lost in life, can't do shit for themselves and having a hard time focusing. Sireng tracks your learning process and turns it into informative graphs for you to use it as a means of review. Sireng has a chat feature that you can use to ask for feedback from your friends or you can simply use it to wake your friend up and tell em to lock tf in.

## Features
- **User Authentication**: Secure authentication using JWT.
- **Session Management**: Tracks study sessions with detailed time logging.
- **Social Interaction**: Add friends and chat along.
- **Redis Caching**: Faster performance with Redis Caching.
- **AWS S3 Integration**: Store and retrieve profile photos using AWS S3 Bucket.

## Getting Started

## Prerequisites
Before you begin, ensure you have the following installed:
- [Go](https://golang.org/dl/) (v1.18 or higher)
- [MySQL](https://www.mysql.com/downloads/) (or compatible database)
- [Redis](https://redis.io/download)
- [Docker](https://www.docker.com/products/docker-desktop) (optional, for containerized deployment)

### Installation
1. **Clone the repository:**
```bash
git clone https://github.com/chaaaeeee/sireng.git
cd sireng
```
2. **Install dependencies:**
```bash
go mod download
```   
3. **Setup configuration:**
```yaml
server:
  port: 8080
database:
  user: root
  password: 
  protocol: tcp
  path: localhost:3306
  dbname: sireng

jwt:
  secret: kobochan
```
4. **Run the application:**
```bash
go run cmd/sireng/main.go
```
## Usage
swagger