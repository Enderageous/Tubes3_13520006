# Backend

## Cara menjalankan di localhost

1. Pastikan sudah menginstal Golang.
2. Jalankan perintah berikut. 
	```
	go run .
	```
	Perintah ini akan menjalankan backend API dari `localhost:8080`.
5. Jika sudah berhasil, API dapat digunakan

## Cara menjalankan dari remote

Jalankan API dari [`enigmatic-brook-59106.herokuapp.com`](https://enigmatic-brook-59106.herokuapp.com/)

## Daftar endpoint
```
GET    /
GET    /api

GET    /api/disease
GET    /api/disease/:id
POST   /api/disease
DELETE /api/disease/:id

GET    /api/prediction
GET    /api/prediction/:id
POST   /api/prediction
DELETE /api/prediction/:id
```