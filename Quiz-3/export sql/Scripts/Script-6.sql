

CREATE TABLE db_mahasiswa.mahasiswa (
    id INT auto_increment primary key,
    nama VARCHAR(255) not null,
    mataKuliah VARCHAR(255) not null,
   	indeksNilai VARCHAR(255) not null,
   	nilai INT not null
    created_at DATETIME not null,
    updated_at DATETIME not null
)