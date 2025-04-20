# 🛡️ Auth API - Microservicio de Autenticación con Go + MySQL + JWT

Este es un microservicio de autenticación simple desarrollado en **Go (Golang)**. Permite a los usuarios **registrarse** y **loguearse** utilizando **JWT** para generar tokens de sesión. Ideal como base para proyectos que requieren autenticación básica.

---

## 🚀 Tecnologías utilizadas

- [Golang](https://golang.org/)
- [Gin](https://github.com/gin-gonic/gin) - Framework HTTP
- [MySQL](https://www.mysql.com/)
- [GORM](https://gorm.io/index.html) - ORM para Go
- [JWT](https://github.com/golang-jwt/jwt) - Tokens para autenticación
- [Docker](https://www.docker.com/) + Docker Compose

---

## 📦 Instalación

### 1. Clona el repositorio

```bash
git clone https://github.com/dalthonmh/todoapp-auth.git
cd todoapp-auth
```

### 2. Configura el .env

Crea un archivo .env en la raíz del proyecton con el siguiente contenido

```bash
DB_DSN=root:password@tcp(mysql:3306)/authdb?charset=utf8mb4 parseTime=True&loc=Local
JWT_SECRET=supersecreto
```

### 3. Ejecuta el proyecto con Docker

```bash
docker compose up --build
```

Esto levantará:

- El backend en Go en http://localhost:8080
- Una instancia de MySQL en el puerto 3306

## 📚 Endpoints

### 📌 POST /register

Registra un nuevo usuario.
Body:

```json
{
  "username": "usuario1",
  "password": "123456"
}
```

respuesta:

```json
{
  "message": "Usuario registrado con éxito"
}
```

### 📌 POST /login

Autentica un usuario y devuelve un JWT.

Body:

```json
{
  "username": "usuario1",
  "password": "123456"
}
```

respuesta:

```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI..."
}
```
