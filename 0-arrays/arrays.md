# Diziler (Arrays)

Bir indis aracılığıyla erişilen ardışıl olarak yerleştirilmiş ve aynı tipteki verilerin tutulduğu yapıya dizi denir.
Avantaj olarak diziler çok basit ifade edilir, kolay anlaşılır ve bellekte fazla yer işgal etmez. Ayrıca dizilerde verinin adresi ile indisi arasında kolaylıkla ilişkiler kurulur.
![Tek boyutlu dizi](https://miro.medium.com/max/768/0*TDrt1RUnaAF2JRD8.jpg)

**Dinamik diziler,** ise dizinin boyutunun tutulduğu ve dizi sınırlarındaki taşmaların kontrol edildiği ve dizinin dinamik olarak genişleyebildiği yapıdır.

![Dinamik diziler](https://thirumalblog.files.wordpress.com/2015/09/5adaf-dynamic2barrays.png)

Ardışıklık yapısının getirdiği en büyük **dezavantaj** ise belirli bir yere yeni elemanların eklenmesi istenirse sağ taraftaki elemanların kaydırılması gerekecektir.
Örneğin {1, 2, 3 ,5, 6} dizisi için 4'ün yerine koyulabilmesi için 5 ve 6 değerinin sağa kaydırılması gerekir. Sadece bununla kalmaz önceden belirlenen dizi boyutunun da bu işlemin boyutuna izin vermesi gerekir. Silme durumunda ise tam tersi olan sola kaydırma işlemi yapılmalıdır.
Zaman açısından elemanların işlenmesinde bu kaydırmalar istenmeyen durumlardır. Bu sebeple bağlı listelerin kullanılması daha uygundur.
