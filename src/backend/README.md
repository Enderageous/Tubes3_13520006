# Backend

## Cara menjalankan di localhost

1. Pastikan sudah menginstal Golang.
2. Copy file `.env.example` di folder `/src` ke file baru `.env` di folder yang sama.
3. Jika sudah ada, isi `.env` dengan nilai yang sesuai.
4. Jalankan perintah berikut. 
	```
	go run .
	```
	Perintah ini akan menjalankan backend API dari `localhost:8080`.
5. Jika sudah berhasil, akan muncul route apa saja yang dapat digunakan

## Daftar route

	GET		/api

	GET		/api/disease
	GET		/api/disease/:id
	POST	/api/disease
	DELETE	/api/disease/:id

	GET		/api/prediction
	GET		/api/prediction/:id
	POST	/api/prediction
	DELETE	/api/prediction/:id