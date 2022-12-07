# Hacktiv8 Golang Final Project 3 Team 8

## Link Deploy API

https://golang-final-project-3-team-8-production.up.railway.app

## Link Postman Documentation

## Anggota

1. Mochamad Suhri Ainur Rifky (GLNG-KS04-001)
2. Raden Muhammad Yudie Sanjaya (GLNG-KS04-016) :x:
3. Varrel Marcellius (GLNG-KS04-021)

## Pembagian Tugas

1. Mochamad Suhri Ainur Rifky (GLNG-KS04-001)

- API Task
- API Category
- Postman (Collection, Environment, Documentation)

2. Raden Muhammad Yudie Sanjaya (GLNG-KS04-016) :x:


3. Varrel Marcellius (GLNG-KS04-021)

- Setup project
- API User
- Readme
- Deploy
- Unit Test

# Cara Install

1. run `docker compose up` untuk menjalankan database
2. run `go run main.go seeder` untuk menjalankan aplikasi dan seeder admin
3. run `go run main.go` untuk menjalankan aplikasi jika sudah pernah seeder admin

## Credential Admin
```
email: admin@gmail.com
password: admin123
```
# List Route

## Users

- `POST` - `/users/register`. Digunakan untuk registrasi pengguna
- `POST` - `/users/login`. Digunakan untuk login pengguna
- `PUT` - `/users/update-account`. Digunakan untuk mengupdate data diri pengguna
- `DELETE` - `/users/delete-account`. Digunakan untuk menghapus akunnya sendiri

## Categories

- `POST` - `/categories`. Digunakan untuk membuat kategori baru
- `GET` - `/categories`. Digunakan untuk mendapat list kategori
- `PATCH` - `/categories/:categoryId`. Digunakan untuk mengupdate data kategori
- `DELETE` - `/categories/:categoryId`. Digunakan untuk menghapus sebuah kategori

## Tasks

- `POST` - `/tasks`. Digunakan untuk membuat sebuah task
- `GET` - `/tasks`. Digunakan untuk mendapat list task
- `PUT` - `/tasks/:taskId`. Digunakan untuk mengupdate data task
- `PATCH` - `/tasks/update-status/:taskId`. Digunakan untuk mengupdate status task
- `PATCH` - `/tasks/update-category/:taskId`. Digunakan untuk mengupdate category task
- `DELETE` - `/tasks/:taskId`. Digunakan untuk menghapus sebuah task