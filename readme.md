# Golang Rest-API

#### Instalasi

Lakukan clone repository

```bash
git clone https://github.com/akmlrhim/golang-rest-api.git
```

Install package

```bash
go mod tidy
```

Jalankan API

```bash
go run .
```

#### Dependensi Utama

```bash
github.com/gorilla/mux v1.8.1
github.com/rs/cors v1.11.1
github.com/sirupsen/logrus v1.9.3
github.com/spf13/viper v1.20.1
gorm.io/driver/mysql v1.6.0
gorm.io/gorm v1.30.1
```

##### Additional Information

```bash
Golang Version : go version go1.23.4 windows/amd64
```

```bash
Jangan lupa untuk mengimport database golang_rest_api.sql ke database terlebih dahulu
atau anda bisa membuat database dan menyesuaikan
```

```bash
BASE_URL : "http:127.0.0.1:9000"
```

##### Gunakan postman collection ini untuk melakukan testing terhadap setiap endpoint

> ```bash
> https://web.postman.co/workspace/My-Workspace~636da315-a9df-45ae-8734-660d0febe959/collection/31413674-67d50246-c5be-4d1e-bb45-3b27f2ccd890?action=share&source=copy-link&creator=31413674
> ```
