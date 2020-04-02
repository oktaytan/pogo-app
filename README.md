# Pogo App

### Uygulama ve Servis Geliştirme bölümü 1. görev

#### Ön tarafta Vue.js, arka tarafta Golang kullanarak yazılmıştır.

---

1 - Kimliklendirme (authentication) JWT ile sağlanmıştır.

2 - Gönderiler, kullanıcı bilgileri gibi tüm veriler SQLite veritabanı üzerinden alınmaktadır.

- Kayıtlı üç kullanıcı vardır. Aşağıdaki kullanıcı bilgileri ile giriş yapılabilmektedir.

  | Kullanıcı Adı | Şifre |
  | ------------- | ----- |
  | cawis         | 1234  |
  | zidak         | 1111  |
  | maya          | 4567  |

---

## Install

```bash
  git clone https://github.com/oktaytan/pogo-app.git
```

### Backend bağımlılıkları

> Bağımlılıkları tek seferde yüklemek için

```bash
  cd api
  go get
```

> ya da manuel olarak yüklemek için

- HTTP router

```bash
  go get -u github.com/gorilla/mux
```

- HTTP middleware

```bash
  go get github.com/urfave/negroni
```

- CORS engellmeleri için

```bash
  go get github.com/rs/cors
```

- SQlite3 driver

```bash
  go get github.com/mattn/go-sqlite3
```

- Json web token implemantasyonu

```bash
  go get github.com/dgrijalva/jwt-go
```

- HTTP isteklerinde jwt kontrolü için

```bash
  go get github.com/auth0/go-jwt-middleware
```

- Şifreyi encrypt olarak kaydetmek için

```bash
  go get golang.org/x/crypto/bcrypt
```

### Frontend bağımlılıkları

- Vue.js
- vue-router (Vue için router)
- vuex ( State management )
- ant-design-vue ( Ant Design UI )
- animate.css ( Animasyonlar için )
- axios ( Http istekleri yapmak için )
- moment ( Tarih ve zaman formatlama için )
- vue-textarea-autosize ( Custom textarea )

```bash
  cd client
  npm install
```

---

## Start

- Projeyi başlatmak için kök dizindeyken

```bash
  npm start
```
