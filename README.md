Project belajar backend Golang menggunakan Echo Framework.

- routing di Echo
- CRUD API
- JWT authentication
- middleware
- PostgreSQL
- struktur folder backend Go
- clean code sederhana pakai service & repository


Struktur Folder
<img width="327" height="742" alt="image" src="https://github.com/user-attachments/assets/4ffc30cc-1b91-42dd-93e8-377b3b6e926e" />


Penjelasan Folder
- config/database.go merupakan file untuk konfigurasi dan koneksi database PostgreSQL.
- handler merupakan layer yang menerima request dari client, mengambil request body, memanggil service, dan mengembalikan response JSON.
- helper/response.go merupakan file reusable untuk format response API agar lebih konsisten.
- middleware/jwt_middleware.go merupakan file middleware untuk validasi token JWT pada endpoint tertentu.
- model merupakan kumpulan struct model yang digunakan untuk mapping data table, request, dan response.
- repository merupakan layer yang langsung berhubungan dengan database dan berisi query-query database.
- routes merupakan file untuk menyimpan dan mengatur seluruh endpoint API.
- service merupakan layer business logic aplikasi dan menjadi penghubung antara handler dan repository.
- utils merupakan file helper tambahan yang bersifat global seperti JWT dan hashing password.
- go.mod merupakan file dependency management pada project Go.
- main.go merupakan file utama untuk menjalankan aplikasi.


Flow Request
Request
- routes
- handler
- service
- repository
- database

Penjelasan singkat:
- Request dari client akan masuk melalui routes.
- Handler menerima request dan memproses input.
- Service menjalankan business logic aplikasi.
- Repository menjalankan query ke database.
- Database mengembalikan data ke aplikasi.
