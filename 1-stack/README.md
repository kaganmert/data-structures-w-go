### Yığın/Stack

- Algoritmaların uygulanmasında kullanılan veri yapılarının en yaygın olanlardan birisi Yığın/Stack kavramıdır. Yapı sadece başlangıç ve son uç elemanları üzerinde ekleme/silme işlemlerine izin vermesi üzerine kuruludur.
- Yığında kümeden silinecek eleman daima kümeye en son eklenen elemandır. Yani "son giren,ilk çıkar(Last In First Out-LIFO)" mantığıyla çalışır. Dolayısıyla bu veri yapısında ilk erişilecek olan veri, diziye son eklenen elemandır. En üstteki eleman işlendikten sonra, örneğin yığından çekildikten sonra bir sonraki elemana erişilir. Böylece dizi üzerinde alt ve üst taşmaların kontrol mekanizması büyük önem taşır.
- Bu kontrolü sağlamak için bir gösterici(pointer) tanımlarız. Bu en üst noktadan(top) erişilir ve bu gösterici dizi boyuna ulaşmışsa taşmalar olur.
- Yığınlar aynı zamanda önceden belirlenmiş bir işlemle elemanları kümeden silinebilen dinamik kümelerdir.
- Bir yığın üzerinde EKLEME(INSERT) işlemi PUSH ve parametre almayan SİLME(DELETE) işlemi ise POP olarak adlandırılır. Bu isimler restoranlarda yer alan tabak yığınlarından esinlenilmiştir. Tabak yığınından tabakları alma sırası,tabakları koyma sırasının tersi böyle en üstteki tabak erişebilirdir.
  [![](https://raw.githubusercontent.com/kaganmert/veri-yapilari/main/assets/1-stack-files/plate.png?token=GHSAT0AAAAAABQWWIROMWIBKFM7DQPKLN4OYP2M34Q)](https://raw.githubusercontent.com/kaganmert/veri-yapilari/main/assets/1-stack-files/plate.png?token=GHSAT0AAAAAABQWWIROMWIBKFM7DQPKLN4OYP2M34Q)
- Bir yığın yapısını gerçekleştirmek için dizi ya da bağlantılı liste kullanabiliriz. Bu başlık altında dizi ile gerçekleştireceğiz.
- O halde elimizde n eleman içeren bir yığını ancak Y[1..n] dizisi ile ifade edilebiliriz.
  **Yığın işlemleri;**

| Yığın İşlemi  | Açıklama                                                                |
| ------------- | ----------------------------------------------------------------------- |
| `push(nesne)` | Yeni bir girdi nesnesini ekler.                                         |
| `pop()`       | En son ekleneni çıkarıp nesne olarak geri döndürür.                     |
| `top()`       | En son ekleneni çıkarmadan nesne olarak geri döndürür                   |
| `size()`      | Toplam yer edinen nesne sayısını tam sayı olarak geri döndürür.         |
| `isEmpty()`   | Yığında bir nesne var mı yok mu bilgisini boolean olarak geri döndürür. |

- Y.top = 0 olduğunda yığın eleman içermez ve boş olarak adlandırılır.
- Yığının boş olup olmadığını Y.isEmpty() sorgusu ile test edebiliriz. Eğer yığın boşsa alttan taşma hatası oluşur. Ayrıca n elemanlı bir yığına eleman eklenmeye çalışılırsa yığında üstten taşma hatası oluşur.

[![](https://raw.githubusercontent.com/kaganmert/veri-yapilari/main/assets/1-stack-files/stack.png?token=GHSAT0AAAAAABQWWIRP3MHOKXWGIQGA5SDOYP2M6JA)](https://raw.githubusercontent.com/kaganmert/veri-yapilari/main/assets/1-stack-files/stack.png?token=GHSAT0AAAAAABQWWIRP3MHOKXWGIQGA5SDOYP2M6JA)

**Yığın yapısı nerelerde kullanılır?**

- Uygulamalarda _Undo_ işlemleri,browser'lardaki _Back_ butonu,yazım kontrollerindeki _parantez kontrolü_ gibi yerlerde stack kullanılır.
- **Dizi Tabanlı Yığın:** Daha önceden bahsettiğimiz gibi dizi üzerinde en fazla N tane eleman tutulabilir. Dizileri soldan sağa doğru ekler,değişken ile en üstteki nesnenin indisini öğrenir ve elemanı çıkarırken bu indis değerini kullanırız.Yığın nesnelerinin saklandığı dizi dolar ise _push()_ işlemi gerçekleştiğinde _FullStackException_ mesajı alınır.
- Yığın veri yapısına matematik ifadelerin sondan gösterimi olan postfix veya polish notasyonu örnek gösterilebilir.
- Aşağıdaki örnekte matematiksel ifadelendirmelerde parantezlemeden kurtularak ifadeleri sadece operatörlerle ifade edebiliriz.

  > **(a-b)_d+(c/e)_(f-k) ifadesi postfix biçiminde ab-d*ce/fk-*+ şeklinde ifade edilir. **

  Yığın ile bu notasyonun elde edinimi şu şekildedir;
  1- Sıradaki eleman operatör ise elemanı yığına koy,
  2- Sıradaki eleman operatör ise yığından 2 eleman al ve işleme tabi tutarak sonucu tekrar yığına koy.

**Yığın işlemlerini kaba olarak şöyle gerçekleştirebiliriz;**

      YIGIN-EMPTY(Y)
    	if.Y.top == 0
    		return TRUE
    	else return FALSE

       PUSH(Y, x)
    	Y.top = Y.top + 1
    	Y[Y.top]=x

       POP(Y)
        if YIGIN-EMPTY(Y)
    	 error "alttan taşma"
    	else Y.top = Y.top-1
    	 return Y[Y.top + 1]

> **n nesne sayısı olan yığında kullanılan alan _O(n)_ iken her bir işlem _O(1_) sürede tamamlanır.**

**Büyüyebilir Dizi Tabanlı Yığın Yaklaşımı**

_push()_ işlemi kullanıldığında dizi dolu ise yığının tutulduğu dizi daha büyük bir dizi ile yer değiştiririz.

**Soru:** O halde yeni dizi büyüklüğü ne kadar olmalı?

**1-Artımlı yaklaşım:** Yığının büyüklüğü sabit bir değer kadar arttırılır.

**2-İkiye katlama yaklaşımı:** Önceki dizi boyutunu iki kat arttırmak.

**Sonuç:** Bağlı liste kullanarak yığını gerçekleştirirsek bu tür problemlerin önüne geçebiliriz.
