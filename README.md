# Hürriyet Lite

[lite.cnn.io](http://lite.cnn.io/en)'dan esinlenilmiş, [Hürriyet Developer API](https://developers.hurriyet.com.tr) üzerine kurulu, sade ve hızlı olmayı hedefleyen haber okuma servisi.

## Nasıl Çalıştırılır

[Hürriyet Developer API](https://developers.hurriyet.com.tr)'a üye olup bir `API Key` oluşturulur ve aşağıdaki komutlar çalıştırılır.

```shell
export HURRIYET_API_KEY=<API Key>
```

### docker-compose

```shell
docker-compose-up
```

Uygulamaya http://localhost:8000 adresinden erişilebilir.

### docker swarm

```shell
docker stack deploy --compose-file docker-compose.prod.yml hurriyet
```