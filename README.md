
Solution :

1. https://github.com/alfiankan/haioo-take-home/blob/main/1-5/pecahan_test.go 
2. https://github.com/alfiankan/haioo-take-home/blob/main/1-5/valid_edit_test.go
3. berikut beberapa perbaikan dan keterangan dari Dockerfile yang disediakan


```bash
FROM golang # mendefinisikan baseimage yang akan digunakan
WORKDIR /go/src/github.com/telkomdev/indihome # mendefiniskan working directory

ADD . /go/src/github.com/telkomdev/indihome/backend # menambahkan folder darihost ke dalam image/container yang akan dijaanknan, pada versi baru lebih familiar menggunakan COPY
RUN go get github.com/tools/godep # menginstall godep, pada go versi baru mempunyai dependency management sendiri yaitu go mod, lebih baik jika menggunakan standart go
RUN godep restore 
# RUN go install github.com/telkomdev/indihome # ini akan menginstall binary namun lebih baik jika dibuild terlebih dahulu maka dari itu saya comment

RUN go build -o /go/bin/indihome /go/src/github.com/telkomdev/indihome # build binary

ENTRYPOINT /go/bin/indihome # entrypointn atauu program pertama yang akan di jalankan ketika container up and running
LISTEN 80 # buka port pada 80
```

berikut versi perbaikan dengan versi standart go baru
```bash
FROM golang # mendefinisikan baseimage yang akan digunakan
WORKDIR /app # mendefiniskan working directory

COPY . . # menambahkan folder darihost ke dalam image/container yang akan dijaanknan, pada versi baru lebih familiar menggunakan COPY

RUN go mod tidy # install semua dependencies yang diperlukan

RUN go build -o /app/indihome . # build binary

ENTRYPOINT ['/app/indihome'] # entrypointn atauu program pertama yang akan di jalankan ketika container up and running
LISTEN 80 # buka port pada 80
```



4. Tujuan penggunaan microservices
  - unntuk independent deployability, ini mempermudah ketika iterasi atau menambah fitur atau maintenace fix problem, kita cukup mengerjakan sebagian kecil dari sistem tanpa mengganggu yanglain
  - independent scalability, ini unutuk skalabilitas yang fleksibel contohnya jika service product itu yang banyak trafficnya bisa di scaleup lebih tinggi dari pada misalkan service wishlist yang mungkin trafficnya lebih kecil
  - dan terakir menurut saya adalah di code management lebih mudah untuk di bagi ke team besar

5. index bekerja dengan membuat sebuah data strujtur baru dari tabel asli, misalkan kita mengindex nama barang maka akan dibuatkan data structure baru yang akan memuat index dan pointer ke data pada tabel asli, data structure yang digunakan adalah yang cepat dan bagus untuk searching dan indexing atau ordering atau operasi lainya, sehingga misalkan kita ingin mencari data berdasarkan nama produk tidak perlu mencari secara linear ke dalam tabel yang akan memakan waktu lama, cukup melalui index begitu ketemu indexnya maka akan dapat pointer yang akan menunjukan lokasi data lengkapnya

6. shoping cart service https://github.com/alfiankan/shopping-cart-service/tree/f86c61cf49fa861b9d49a4f5b6f01b2de1444f26
