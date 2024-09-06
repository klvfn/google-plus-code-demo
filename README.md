## Given problems

- Validasi apakah titik geolocation customer masih dalam jarak 1 kabupaten, dengan asumsi ongkir akan sama
- Radius kota terluas di indonesia
    - Di Indonesia, kota-kota besar memiliki radius yang bervariasi tergantung pada ukuran dan batas administratifnya. Berikut adalah beberapa kota besar di Indonesia dengan estimasi radiusnya, menggunakan data populasi dan ukuran wilayah:
    - Estimasi Radius Kota Besar di Indonesia
        | **Kota** | **Luas (km²)** | **Radius Estimasi (km)** |
        | --- | --- | --- |
        | **Jakarta** | 662.33 | ~14.5 km |
        | **Surabaya** | 350.54 | ~10.6 km |
        | **Bandung** | 167.16 | ~7.3 km |
        | **Medan** | 265.10 | ~9.2 km |
        | **Makassar** | 175.77 | ~7.5 km |
    - Penjelasan
        - **Radius Estimasi**: Radius ini dihitung dari luas kota dengan menganggap kota sebagai lingkaran untuk memberikan gambaran kasar. Radius ini tidak memperhitungkan bentuk kota yang sebenarnya, yang mungkin tidak berbentuk lingkaran sempurna.
        - **Jakarta**: Sebagai ibu kota negara dan kota terbesar di Indonesia, Jakarta memiliki radius yang lebih besar dibandingkan kota-kota lainnya.
        - **Surabaya**: Kota terbesar kedua di Indonesia dengan luas yang cukup besar namun lebih kecil dibandingkan Jakarta.
        - **Bandung**: Terkenal dengan iklim sejuk dan sebagai pusat pendidikan serta teknologi.
        - **Medan**: Kota terbesar di Sumatera dengan luas yang relatif besar.
        - **Makassar**: Kota terbesar di Sulawesi dan pusat ekonomi penting di kawasan timur Indonesia.
    - Untuk perhitungan lebih akurat atau data lebih spesifik, terutama jika melibatkan bentuk kota yang tidak bulat, Anda bisa menggunakan alat GIS atau data administratif resmi yang disediakan oleh pemerintah daerah atau Badan Pusat Statistik (BPS).
- Radius kabupaten terluas di indonesia
    - Di Indonesia, kabupaten-kabupaten memiliki berbagai ukuran wilayah yang bervariasi. Berikut adalah beberapa kabupaten terluas di Indonesia berdasarkan luas wilayahnya, beserta estimasi radiusnya jika kota-kota tersebut dianggap sebagai lingkaran untuk tujuan perhitungan kasar:
    - Kabupaten Terluas di Indonesia
        
        
        | **Kabupaten** | **Luas (km²)** | **Radius Estimasi (km)** |
        | --- | --- | --- |
        | **Kabupaten Merauke** | 45,018 | ~119.6 km |
        | **Kabupaten Kutai Kartanegara** | 27,263 | ~93.2 km |
        | **Kabupaten Malinau** | 18,365 | ~76.4 km |
        | **Kabupaten Berau** | 34,125 | ~104.0 km |
        | **Kabupaten Puncak Jaya** | 16,846 | ~73.1 km |
    - Penjelasan
        - **Radius Estimasi**: Radius dihitung dengan asumsi kabupaten berbentuk lingkaran untuk memberikan gambaran kasar. Radius ini tidak memperhitungkan bentuk geografis sebenarnya dari kabupaten yang mungkin tidak bulat sempurna.
        - **Kabupaten Merauke**: Merupakan kabupaten terluas di Indonesia dengan luas yang sangat besar, terletak di Papua.
        - **Kabupaten Kutai Kartanegara**: Terletak di Kalimantan Timur dan merupakan salah satu kabupaten dengan luas terbesar.
        - **Kabupaten Malinau**: Juga berada di Kalimantan Utara, memiliki wilayah yang luas.
        - **Kabupaten Berau**: Terletak di Kalimantan Timur dan mencakup area yang luas.
        - **Kabupaten Puncak Jaya**: Terletak di Papua dan dikenal dengan wilayahnya yang luas dan topografi pegunungan.

## Open location code a.k.a Plus Code

### Definition

Open Location Code is a technology that gives a way of encoding location into a form that is easier to use than latitude and longitude. The codes generated are called plus codes, as their distinguishing attribute is that they include a "+" character.

The technology is designed to produce codes that can be used as a replacement for street addresses, especially in places where buildings aren't numbered or streets aren't named.

Plus codes represent an area, not a point. As digits are added to a code, the area shrinks, so a long code is more precise than a short code.

Codes that are similar are located closer together than codes that are different.

**A location can be converted into a code, and this (full) code can be converted back to a location completely offline**, without any data tables to lookup or online services required.

Codes can be shortened for easier communication, in which case they can be used regionally or in combination with a reference location that all users of this short code need to be aware of. If the reference location is given in form of a location name, use of a geocoding service might be necessary to recover the original location.

Algorithms to

- encode and decode full codes,
- shorten them relative to a reference location, and
- recover a location from a short code and a reference location given as latitude/longitude pair

are publicly available and can be used without restriction. Geocoding services are not a part of the Open Location Code technology.

### Block size

### Summary Table

| **OLC Length** | **Block Size** | **Coverage Area** |
| --- | --- | --- |
| 2 Characters | ~1,000 km x 1,000 km (1,000,000 sq km) | Extremely broad |
| 4 Characters | ~100 km x 100 km (10,000 sq km) | Very broad |
| 6 Characters | ~10 km x 10 km (100 sq km) | Large area |
| 8 Characters | ~1 km x 1 km (1 sq km) | Medium-sized area |
| 10 Characters | ~100 m x 100 m (10,000 sq m) | Small area |
| 12 Characters | ~10 m x 10 m (100 sq m) | Very precise |
| 14 Characters | < 1 m x 1 m (< 1 sq m) | Extremely precise |
- A way to categorized OLC length https://github.com/bocops/open-geotiling/blob/master/java/org/bocops/opengeotiling/OpenGeoTile.java
- Best way to check OLC adjacency according to its author https://groups.google.com/g/open-location-code/c/qmqDqwStWVw/m/Va4t7UAhAgAJ

## APIs

- HTTP API
    - Ref: https://github.com/google/open-location-code/wiki/Plus-codes-API/aef6adb82b1fe62e95f718075243809e42c5e3f0#encoding-requests
    - Without API key
        
        ```json
        https://plus.codes/api?address=-6.1932611,106.7441334
        ```
        
    - Go SDK: https://pkg.go.dev/github.com/google/open-location-code/go#CodeArea
