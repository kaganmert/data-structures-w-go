**KUYRUK** yapısı, ilk giren ilk çıkar veya FIFO (first in - first out) kuralıyla elemanlara erişebildiğimiz bir veri yapısıdır. Yani ilk eleman bizim işlem yapabileceğimiz yerdir.

-fifo resim

Bu yapıyı anlamak için, bir veznede ödeme yapmak için sıraya giren müşteri yapısıyla anlayabiliriz. Kuyrukta başlık(head) ve kuyruk(tail) nitelikleri söz konusudur. Yapıya bir eleman ekleyeceğimiz zaman işlemi **kuyruğun sonu**ndan, yani ödeme yapmak için vezneye gelen bir müşteri sıranın sonuna geçer. Ödemeyi yapan müşteri ise kuyruk yapısından yani **kuyruğun başı**ndan silinir.

Kuyrukta verilerin başını ve sonunu gösterecek **işaretçi**lere(pointer) ihtiyacımız vardır. Burada kuyruğun başını yani ilk elemanının indisini gösteren, **Q.head** niteliğine sahip olur. Kuyruğa en son eklenen eleman ise **Q.tail** niteliğidir. Kuyruğun elemanları, dizinin Q.head, Q.head+1, ..., Q.tail-1 indislerinde konumlandırılır.

Eğer bu yapıyı dizi halinde oluştursaydık; kuyruk boyuna ulaştığında taşma olacak ve/ya kuyruktan elemanlar silinmiş ise ön kısımlar boş olacaktı. Bu durumda yeni elemanların eklenmesine imkan sağlayabilmek için varolan elemanları kuyruğun başına kaydırabiliriz, ancak bunun karşılığında ** işlem yükü** olacaktır. İşlem yükü nedeniyle, yapının n indisli elemanını dizinin 1.elemanını takip edecek şekilde **dairesel** olarak sarmallayabiliriz.

-sarmal yapı resim

Bu dairesel yapıda Q.head = Q.tail ise kuyruk boştur. Yani başlangıç durumu Q.head = Q.tail = 1'dir. Bu durumlar göz önüne alındığında boş bir kuyruktan eleman silinmeye çalışılırsa, "alttan taşma" hatası ele alınmalıdır. Aynı şekilde Q.head = Q.tail+1 olduğunda ise yeni bir eleman eklendiğinde de "taşma" hatası oluşur.

Şimdi tüm bahsettiklerimizin sözde kodunu yazalım.

n = Q.length olmak üzere;

    ENQUEUE(Q,x)
    Q[Q.tail]=x
    if Q.tail == Q.length
    	Q.tail = 1
    else Q.tail = Q.tail +1

ve

    DEQUEUE(Q)
    x = Q[Q.head]
    if Q.head == Q.length
    	Q.head = 1
    else Q.head = Q.head +1
    return x

İki işlemde O(1) sürede tamamlanır.

**Not:** Bu sözde kodlarda taşma durumları ele alınmamıştır.

## Kuyruk veri yapısı nerelerde kullanılır?

1.  Yazıcılarda yazdırma sırasının ele alınmasında,
2.  Dosya yüklemelerinde,
3.  Labirent algolaritmalarında yolun bulunmasında,
4.  Arama algoritmalarının enine arama yöntemlerinde,
5.  İşletim Sistemlerinde; Semafor, FCFS Scheduling, Klavye gibi buffer işlemlerinde
6.  Network Alanında; Routers/Switches, Mail Queues
