# socb4zz-analyzer-app

**socb4zz-analyzer-app** adalah aplikasi berbasis **Go (Gin)** dan **Next.js** dengan **PostgreSQL** yang berfungsi untuk menganalisis aktivitas sosial media.  
Aplikasi ini mendukung pengelolaan data pengguna, penyimpanan hasil analisis, serta integrasi AI untuk memberikan insight dan visualisasi interaktif terkait tren sosial.

## ✨ Fitur Utama
- 🔐 **Autentikasi & Manajemen Pengguna**
- 📊 **Analisis Aktivitas Sosial Media**
- 🤖 **Integrasi AI/NLP** untuk insight otomatis
- 🗄️ **Penyimpanan Hasil Analisis** ke PostgreSQL
- 📈 **Visualisasi Data** secara interaktif di frontend

---

## 📂 Struktur Proyek
```
socb4zz-analyzer-app/
│── backend/ # Backend service (Go + Gin)
│ ├── main.go
│ ├── database/
│ ├── controllers/
│ ├── models/
│ └── routes/
│
│── frontend/ # Frontend service (Next.js)
│ ├── pages/
│ ├── components/
│ ├── public/
│ └── styles/
│
└── README.md
```

---

## 🛠️ Tech Stack
- **Backend**: Go (Gin)
- **Frontend**: Next.js (React + TypeScript)
- **Database**: PostgreSQL
- **AI/ML**: NLP untuk analisis konten sosial media
- **Deployment**: Docker (opsional)

---

## 🚀 Cara Menjalankan

### 1. Persiapan
Pastikan sudah menginstal:
- [Go](https://go.dev/) (v1.20+)
- [Node.js](https://nodejs.org/) (v18+)
- [PostgreSQL](https://www.postgresql.org/) (v14+)
- [Git](https://git-scm.com/)

Clone repository:
```bash
git clone https://github.com/username/socb4zz-analyzer-app.git
cd socb4zz-analyzer-app
```

### 2. Setup Database
Buat database baru di PostgreSQL:
```sql
CREATE DATABASE socb4zz_db;
```

Atur koneksi database di backend/config/config.go atau .env:
```env
DB_HOST=localhost
DB_PORT=yourport
DB_USER=youruser
DB_PASSWORD=yourpassword
DB_NAME=socb4zz_db
```

### 3. Install Dependensi & Jalankan Aplikasi
Jalankan semua langkah di bawah ini tanpa dipisah-pisah:
```bash
# masuk folder backend dan install dependensi Go
cd backend
go mod tidy

# jalankan backend
go run main.go &
cd ..

# masuk folder frontend dan install dependensi Next.js
cd frontend
npm install

# jalankan frontend
npm run dev
```

Aplikasi backend akan jalan di [http://localhost:8080](http://localhost:8080) 
Frontend (Next.js) akan jalan di [http://localhost:3000](http://localhost:3000)

## Catatan
- Pastikan PostgreSQL aktif sebelum menjalankan backend.
- Jika ingin deploy ke server/VPS, gunakan Docker atau setup CI/CD agar otomatis update saat push ke GitHub.
- Untuk integrasi AI/NLP, service Python (FastAPI) bisa ditambahkan sebagai microservice tambahan.

## 📜 License
Distributed under Mochamad Dimas Putra Hermawan License. See `LICENSE` for more information.