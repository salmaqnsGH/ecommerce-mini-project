CREATE TABLE users (
    id                  INT             NOT NULL AUTO_INCREMENT,
    nama                VARCHAR(255),
    kata_sandi          VARCHAR(255),
    no_telp             VARCHAR(255),
    tanggal_lahir       DATETIME,
    jenis_kelamin       VARCHAR(255),
    tentang             TEXT,
    pekerjaan           VARCHAR(255),
    email               VARCHAR(255),
    id_provinsi         VARCHAR(255),
    id_kota             VARCHAR(255),
    is_admin            BOOLEAN,
    created_at          DATETIME,
    updated_at          DATETIME,
    PRIMARY KEY (id)
)ENGINE = InnoDB;

CREATE TABLE categories (
    id                  INT             NOT NULL AUTO_INCREMENT,
    nama_category       VARCHAR(255),
    created_at          DATETIME,
    updated_at          DATETIME,
    PRIMARY KEY (id)
)ENGINE = InnoDB;

CREATE TABLE tokos (
    id                  INT             NOT NULL AUTO_INCREMENT,
    id_user             INT NOT NULL,
    nama_toko           VARCHAR(255),
    url_foto            VARCHAR(255),
    created_at          DATETIME,
    updated_at          DATETIME,
    PRIMARY KEY (id)
)ENGINE = InnoDB;

ALTER TABLE tokos ADD FOREIGN KEY (id_user) REFERENCES users(id);

CREATE TABLE alamats (
    id                  INT             NOT NULL AUTO_INCREMENT,
    id_user             INT NOT NULL,
    judul_alamat        VARCHAR(255),
    nama_penerima       VARCHAR(255),
    no_telp             VARCHAR(255),
    detail_alamat       VARCHAR(255),
    created_at          DATETIME,
    updated_at          DATETIME,
    PRIMARY KEY (id)
)ENGINE = InnoDB;

ALTER TABLE alamats ADD FOREIGN KEY (id_user) REFERENCES users(id);

CREATE TABLE trxs (
    id                  INT             NOT NULL AUTO_INCREMENT,
    id_user             INT NOT NULL,
    alamat_pengiriman   INT NOT NULL,
    harga_total         INT,
    kode_invoice        VARCHAR(255),
    method_bayar        VARCHAR(255),
    created_at          DATETIME,
    updated_at          DATETIME,
    PRIMARY KEY (id)
)ENGINE = InnoDB;

ALTER TABLE trxs ADD FOREIGN KEY (id_user) REFERENCES users(id);
ALTER TABLE trxs ADD FOREIGN KEY (alamat_pengiriman) REFERENCES alamats(id);

CREATE TABLE produks (
    id                  INT             NOT NULL AUTO_INCREMENT,
    id_toko             INT NOT NULL,
    id_category         INT NOT NULL,
    nama_produk         VARCHAR(255),
    slug                VARCHAR(255),
    harga_reseller      INT,
    harga_konsumen      INT,
    stok                INT,
    deskripsi           TEXT,
    created_at          DATETIME,
    updated_at          DATETIME,
    PRIMARY KEY (id)
)ENGINE = InnoDB;

ALTER TABLE produks ADD FOREIGN KEY (id_toko) REFERENCES tokos(id);
ALTER TABLE produks ADD FOREIGN KEY (id_category) REFERENCES categories(id);

CREATE TABLE foto_produks (
    id                  INT             NOT NULL AUTO_INCREMENT,
    id_produk           INT NOT NULL,
    url                 VARCHAR(255),
    created_at          DATETIME,
    updated_at          DATETIME,
    PRIMARY KEY (id)
)ENGINE = InnoDB;

ALTER TABLE foto_produks ADD FOREIGN KEY (id_produk) REFERENCES produks(id);

CREATE TABLE log_produks (
    id                  INT             NOT NULL AUTO_INCREMENT,
    id_produk           INT NOT NULL,
    id_toko             INT NOT NULL,
    id_category         INT NOT NULL,
    nama_produk         VARCHAR(255),
    slug                VARCHAR(255),
    harga_reseller      VARCHAR(255),
    harga_konsumen      VARCHAR(255),
    deskripsi           TEXT,
    created_at          DATETIME,
    updated_at          DATETIME,
    PRIMARY KEY (id)
)ENGINE = InnoDB;

ALTER TABLE log_produks ADD FOREIGN KEY (id_produk) REFERENCES produks(id);
ALTER TABLE log_produks ADD FOREIGN KEY (id_toko) REFERENCES tokos(id);
ALTER TABLE log_produks ADD FOREIGN KEY (id_category) REFERENCES categories(id);

CREATE TABLE detail_trxs (
    id                  INT             NOT NULL AUTO_INCREMENT,
    id_trx              INT NOT NULL,
    id_log_produk       INT NOT NULL,
    id_toko             INT NOT NULL,
    kuantitas           INT,
    harga_total         INT,
    created_at          DATETIME,
    updated_at          DATETIME,
    PRIMARY KEY (id)
)ENGINE = InnoDB;

ALTER TABLE detail_trxs ADD FOREIGN KEY (id_trx) REFERENCES trxs(id);
ALTER TABLE detail_trxs ADD FOREIGN KEY (id_log_produk) REFERENCES log_produks(id);
ALTER TABLE detail_trxs ADD FOREIGN KEY (id_toko) REFERENCES tokos(id);

INSERT INTO users VALUES (1, "nama", "sandi", "0085999999111", CURRENT_TIMESTAMP, "P", "tentang", "pekerjajan", "email@mail.com", "1", "1", true, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
INSERT INTO categories VALUES (1, "nama_category", CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
INSERT INTO tokos VALUES (1, 1, "nama_toko","url_foto", CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
INSERT INTO alamats VALUES (1, 1, "judul_alamat", "nama_penerima", "085225226336", "detail_alamat", CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
INSERT INTO trxs VALUES (1, 1, 1, 1000, "kode_invoice", "method_bayar", CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
INSERT INTO produks VALUES (1, 1, 1, "nama_produk", "slug",15000, 17000, 10, "desc", CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
INSERT INTO foto_produks VALUES (1, 1, "url", CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
INSERT INTO log_produks VALUES (1, 1, 1, 1, "nama_produk", "slug", 15000, 17000, "desc", CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
INSERT INTO detail_trxs VALUES (1, 1, 1, 1, 11, 2000, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);













