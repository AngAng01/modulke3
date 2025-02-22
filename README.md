# MODUL TIGA
 ## SOAL 1
   Program di atas adalah implementasi yang menghitung permutasi dan kombinasi dari dua pasang bilangan input.
   
   ## Overview
      Program ini terdiri dari satu file bernama 'main.go' dan mencakup komponen-komponen utama berikut:
      - Pernyataan 'package main', yang mendefinisikan paket untuk program yang dapat dieksekusi.
      - Pernyataan 'import', yang digunakan untuk menyertakan paket-paket yang diperlukan (dalam hal ini, 'fmt').
      - Fungsi 'main()', yang merupakan titik awal dari setiap program Go.
      - Fungsi 'factorial(n int)', yang menghitung faktorial dari bilangan 'n'
      - Fungsi 'permutation(a, c int)', yang menghitung permutasi 'P(a,c)'
      - Fungsi 'combination(a, c int)', yang menghitung kombinasi 'C(a,c)'

      
   ## Code Explanation
      ```go
            func factorial(n int) int {
                if n == 0 || n == 1 {
                    return 1
                }
                result := 1
                for i := 2; i <= n; i++ {
                    result *= i
                }
                return result
            }
      ```
   Kode di atas adalah implementasi fungsi factorial dalam bahasa Go yang digunakan untuk menghitung faktorial dari suatu bilangan bulat positif. Faktorial dari sebuah bilangan n, ditulis sebagai n!, didefinisikan sebagai hasil perkalian semua bilangan bulat positif dari 1 hingga n.
   
   ```go
      func permutation(a, c int) int {
          return factorial(a) / factorial(a-c)
      }
   ```
   Kode di atas adalah implementasi fungsi permutation yang digunakan untuk menghitung permutasi dari dua bilangan a dan c. Permutasi digunakan untuk menghitung berapa banyak cara untuk mengatur c objek yang dipilih dari a objek secara berurutan.

   ```go
      func combination(a, c int) int {
          return factorial(a) / (factorial(c) * factorial(a-c))
      }
   ```
   Kode di atas adalah implementasi fungsi combination yang digunakan untuk menghitung kombinasi dari dua bilangan a dan c. Kombinasi menghitung berapa banyak cara untuk memilih c objek dari a objek tanpa memperhatikan urutan.

   ```go
      var a, b, c, d int
   ```
   kode diatas adalah kode yang digunakan untuk mendeklarasikan 4 variabel int yang digunakan sebagai inputan.

   ```go
      fmt.Println("Masukkan 4 angka (a, b, c, d): ")
      fmt.Scan(&a, &b, &c, &d)
   ```
   kode diatas adalah kode yang digunakan untuk menampilkan pesan kepada pengguna agar pengguna menginputkan suatu nilai untuk dieksekusi.

   ```go
      permAC := permutation(a, c)
      combAC := combination(a, c)
  
      permBD := permutation(b, d)
      combBD := combination(b, d)
   ```
   Kode di atas adalah bagian kode yang digunakan untuk menghitung permutasi dan kombinasi untuk dua pasang bilangan input, yaitu (a, c) dan (b, d), dengan memanfaatkan fungsi permutation dan combination.

   ```go
      fmt.Printf("P(%d, %d) = %d\n", a, c, permAC)
      fmt.Printf("C(%d, %d) = %d\n", a, c, combAC)
      fmt.Printf("P(%d, %d) = %d\n", b, d, permBD)
      fmt.Printf("C(%d, %d) = %d\n", b, d, combBD)
   ```
   Kode di atas digunakan untuk menampilkan hasil perhitungan permutasi dan kombinasi ke layar dengan format tertentu. Fungsi fmt.Printf digunakan untuk mencetak string yang diformat.


   ## SOAL 2
   Program di atas adalah sebuah program yang mendemonstrasikan penggunaan fungsi komposisi dengan tiga fungsi dasar f(x), g(x), dan h(x). Program ini meminta pengguna untuk memasukkan tiga angka dan kemudian menampilkan hasil dari beberapa komposisi fungsi yang berbeda: fogoh(a), gohof(b), dan hofog(c).
   
   ## Overview
      Program ini terdiri dari satu file bernama 'main.go' dan mencakup komponen-komponen utama berikut:
      - Pernyataan 'package main', yang mendefinisikan paket untuk program yang dapat dieksekusi.
      - Pernyataan 'import', yang digunakan untuk menyertakan paket-paket yang diperlukan (dalam hal ini, 'fmt').
      - Fungsi 'main()', yang merupakan titik awal dari setiap program Go.
      - Fungsi f yang menerima parameter integer x dan mengembalikan hasil kuadrat dari x.
      - Fungsi g yang menerima parameter integer x dan mengembalikan hasil dari x dikurangi 2.
      - Fungsi h yang menerima parameter integer x dan mengembalikan hasil dari x ditambah 1.
      - Fungsi komposisi fogoh yang menggabungkan fungsi f, g, dan h dalam urutan f(g(h(x))).
      - Fungsi komposisi gohof yang menggabungkan fungsi g, h, dan f dalam urutan g(h(f(x))).
      - Fungsi komposisi hofog yang menggabungkan fungsi h, f, dan g dalam urutan h(f(g(x))).

      
   ## Code Explanation
      ```go
          func f(x int) int {
              return x * x
          }
      ```
   Kode di atas adalah fungsi yang digunakan untuk menghitung kuadrat dari bilangan bulat yang diberikan sebagai parameter.
   
   ```go
      func g(x int) int {
          return x - 2
      }
   ```
   Kode di atas adalah fungsi yang digunakan untuk mengurangi dua dari bilangan bulat yang diberikan sebagai parameter.

   ```go
      func h(x int) int {
          return x + 1
      }
   ```
   Kode di atas adalah fungsi yang digunakan untuk menambahkan satu ke bilangan bulat yang diberikan sebagai parameter.

   ```go
      func fogoh(x int) int {
          return f(g(h(x)))
      }
   ```
   Kode di atas adalah fungsi yang mendefinisikan komposisi tiga fungsi: f, g, dan h. Fungsi ini bernama fogoh.

   ```go
      func gohof(x int) int{
          return g(h(f(x)))
      }
   ```
   Kode di atas adalah fungsi yang mendefinisikan komposisi tiga fungsi: g, h, dan f. Fungsi ini bernama gohof.
   
   ```go
      func hofog(x int) int {
          return h(f(g(x)))
      }
   ```
   Kode di atas adalah fungsi yang mendefinisikan komposisi tiga fungsi: h, f, dan g. Fungsi ini bernama hofog.

   ```go
      var a, b, c int
   ```
   kode diatas adalah kode yang digunakan untuk mendeklarasikan 4 variabel int yang digunakan sebagai inputan.

   ```go
      fmt.Scan(&a, &b, &c)
   ```
   kode diatas adalah kode yang meminta inputan kepada pengguna untuk dieksekusi.

   ```go
      fmt.Println(fogoh(a))
      fmt.Println(gohof(b))
      fmt.Println(hofog(c))
   ```
   Kode di atas adalah bagian dari fungsi main, yang digunakan untuk mencetak hasil dari tiga fungsi komposisi yang telah didefinisikan sebelumnya: fogoh, gohof, dan hofog.


   ## SOAL 3
   Program di atas adalah program yang digunakan untuk menentukan posisi suatu titik relatif terhadap dua lingkaran. Program ini menghitung jarak titik dari pusat masing-masing lingkaran dan menentukan apakah titik tersebut berada di dalam atau di luar lingkaran. 
   
   ## Overview
      Program ini terdiri dari satu file bernama 'main.go' dan mencakup komponen-komponen utama berikut:
      - Pernyataan 'package main', yang mendefinisikan paket untuk program yang dapat dieksekusi.
      - Pernyataan 'import', yang digunakan untuk menyertakan paket-paket yang diperlukan (dalam hal ini, 'fmt' dan 'math'), yang dimana 'math' digunakan untuk fungsi matematika, khususnya untuk menghitung akar kuadrat
      - Fungsi 'main()', yang merupakan titik awal dari setiap program Go.
      - Fungsi distance yang menghitung jarak antara dua titik dalam ruang dua dimensi menggunakan rumus jarak Euclidean. Mengembalikan nilai jarak dalam tipe float64.
      - Mendeklarasikan variabel untuk menyimpan koordinat pusat (cx1, cy1) dan jari-jari (r1) dari lingkaran pertama, serta membaca input dari pengguna.
      - Mendeklarasikan variabel untuk menyimpan koordinat pusat (cx2, cy2) dan jari-jari (r2) dari lingkaran kedua, serta membaca input dari pengguna.
      - Mendeklarasikan variabel untuk menyimpan koordinat titik (x, y) yang akan dianalisis, serta membaca input dari pengguna.

      
   ## Code Explanation
    ```go
          func distance(x1, y1, x2, y2 int) float64 {
              return math.Sqrt(float64((x2-x1)*(x2-x1) + (y2-y1)*(y2-y1)))
          }
    ```
    Kode di atas adalah fungsi yang digunakan untuk menghitung jarak antara dua titik dalam ruang dua dimensi.
   
   ```go
      var cx1, cy1, r1 int
      fmt.Scan(&cx1, &cy1, &r1)

      var cx2, cy2, r2 int
      fmt.Scan(&cx2, &cy2, &r2)
   ```
   Kode di atas adalah bagian dari program yang digunakan untuk mendeklarasikan 6 variabel int dan membaca input dari pengguna untuk dieksekusi. 
   
   ```go
      var x, y int
      fmt.Scan(&x, &y)
   ```
   Kode di atas adalah bagian dari program yang digunakan untuk mendeklarasikan variabel dan membaca input dari pengguna untuk koordinat titik.

   ```go
      d1 := distance(x, y, cx1, cy1)
      d2 := distance(x, y, cx2, cy2)
   ```
   Kode di atas adalah bagian dari program yang digunakan untuk menghitung jarak antara suatu titik dan pusat dari dua lingkaran.

   ```go
      inCircle1 := d1 <= float64(r1)
      inCircle2 := d2 <= float64(r2)
   ```
   Kode di atas adalah bagian dari program yang digunakan untuk menentukan apakah suatu titik berada di dalam lingkaran tertentu berdasarkan jarak dari titik tersebut ke pusat lingkaran dan jari-jari lingkaran itu.

   ```go
      if inCircle1 && inCircle2 {
          fmt.Println("Titik di dalam lingkaran 1 dan 2")
      } else if inCircle1 {
          fmt.Println("Titik di dalam lingkaran 1")
      } else if inCircle2 {
          fmt.Println("Titik di dalam lingkaran 2")
      } else {
          fmt.Println("Titik di luar lingkaran 1 dan 2")
      }
   ```
  Kode di atas adalah bagian dari program yang digunakan untuk menentukan dan menampilkan status posisi suatu titik relatif terhadap dua lingkaran berdasarkan hasil perhitungan sebelumnya. 

   
