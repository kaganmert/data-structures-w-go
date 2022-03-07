#Linked List

- Dizilerde anladığımız üzere verilerin bellekte durumu önceden kesin olarak belirlidir. Ardışık indisli değerler ve ardışık bellek bölgelerinde bulunur. Bu nedenle çoğu zaman silme ve eklemede elemanların da topluca kaydırılması gerekir.

- Tam bu noktada bağlı listeler devreye girer. Elemanların bağlantılı şekilde gösterilmesinde her bir elemanı gösteren belirli bir işaretçi(pointer) bulunur. Ayrıca bu elemanın komşusunu gösteren işaretçi de kullanılır.

- Son olarak dizi başlangıcını ifade eden işaretçiye de gerek duyulur. Bu veri yapısı türüne bağlı liste denir.

- Özetle, farklı bellek bölgelerinde bulunan verilerin özel adres bağlantıları aracılığıyla bir araya getirildiği veri yapısıdır.

##Basit Bir Liste Yapısı

Basit bir bağlı liste yapısında, `Liste eleman değeri` ve `Bir sonraki liste elemanına bağlantıyı sağlayan işaretçi` olmak üzere iki parçadan oluşur.

----GÖRSEL

Şimdi bağlı listelerde arama, ekleme ve silme durumlarını yazalım.

##Bağlı Listelerde Arama

---

`SEARCH(L, k)` fonksiyonu `L` listesinde `k` anahtarına sahip ilk elemanı basit doğrusal arama ile bularak bu nesneyi gösteren bir işaretçi döndürür. Eğer listede `k` anahtarına sahip bir eleman yoksa, fonksiyon `NIL` değerini döndürür.

    SEARCH(L,k)
    x = L.head
    while x≠ NIL and x.key ≠ k
    	x = x.next
    return x

`n` elemana sahip bir listede, yukarıda yazdığımız arama fonksiyonu en kötü durumda listenin tamamında arama yapacağı için `θ(n)` sürede tamamlanır.

##Bağlı Listeye Eleman Ekleme

---

Key değerini önceden atanmış bir x nesnesi verildiğinde `INSERT` prosedürü x elemanını listenin önüne ekler ve listenin başı olarak belirlenir.

    INSERT(L, x)
    x.next = L.head
    if L.head ≠NIL
    	L.head.prev = x
    L.head = x
    x.prev = NIL

`n` elemanlı bir listede` INSERT` fonksiyonu `O(1)` sürede tamamlanır.

##Bağlı Listeden Eleman Silme

`DELETE` fonksiyonu bir `L` listesinden bir `x` elemanını siler. Bu işlem için `x` elemanını gösteren bir işaretçi verilmek zorundadır ve bu işlem listedeki işaretçileri güncelleyerek `x` elemanını listeden çıkarır.

Eğer listede değeri belli olan bir eleman silinmek istenirse öncelikle `SEARCH` fonksiyonu çağırılır ve silinecek elemanı gösteren pointeri belirleriz.

    DELETE(L, x)
    if x.prev ≠ NIL
    	x.prev.next = x.next
    else L.head = x.next
    if x.next ≠ NIL
    	x.next.prev = x.prev

`DELETE `işlemi `O(1)` sürede tamamlanır, ancak değeri belli olan bir eleman silinecek ise, en kötü durumda önce `SEARCH` çağırılır ve silinecek eleman liste üzerinden bulunması gerektiğinden `θ(n)` sürede tamamlanır.
