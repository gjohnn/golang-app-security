# QuickResto GO - API con Echo, JWT y Roles

![Go](https://img.shields.io/badge/Go-1.21+-blue)
![Echo](https://img.shields.io/badge/Echo-v4-green)
![JWT](https://img.shields.io/badge/JWT-Auth-orange)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-15-blue)

API backend desarrollada en Go utilizando el framework Echo, con autenticación JWT y sistema de roles para control de acceso.

## Características Principales

- ✅ Autenticación segura con JWT
- ✅ Sistema de roles (admin, user, etc.)
-  CRUD completo de usuarios (en progreso)
- ✅ Middlewares para protección de rutas
- ✅ Integración con PostgreSQL
- ✅ Migraciones automáticas de modelos (Ver codigo de db.go)

## Estructura del Proyecto

```
quick-resto-go/
├── config/
├── controllers/
├── database/
├── middlewares/
├── models/
├── routes/
├── utils/
├── .env.example
├── docker-compose.yml
├── Dockerfile
├── go.mod
├── go.sum
└── main.go
```

## Requisitos

- Go 1.21+
- PostgreSQL 15+
- Docker (opcional)

## Configuración Inicial

1. Copiar el archivo `.env.example` a `.env`:
```bash
cp .env.example .env
```

2. Editar las variables en `.env`:
```env
# PostgreSQL
POSTGRES_HOST=localhost
POSTGRES_USER=postgres
POSTGRES_PASSWORD=secret
POSTGRES_DB=quickresto
POSTGRES_PORT=5432

# App
APP_PORT=8080
JWT_SECRET=mi_super_secreto_jwt
```

## Instalación

### Opción 1: Con Docker (recomendada)
```bash
docker-compose up -d --build
```

### Opción 2: Manual
```bash
# Instalar dependencias
go mod download

# Iniciar la base de datos (PostgreSQL)
docker-compose up -d db

# Ejecutar la aplicación
go run main.go
```

## Uso Básico

### Registro de Usuario
```bash
curl -X POST http://localhost:8080/register \
  -H "Content-Type: application/json" \
  -d '{
      "username": "admin",
      "password": "Admin123!",
      "email": "admin@example.com",
      "role": "admin"
  }'
```

### Login
```bash
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{
      "username": "admin",
      "password": "Admin123!"
  }'
```

### Acceso a Ruta Protegida
```bash
curl -X GET http://localhost:8080/api/protected \
  -H "Authorization: Bearer <TU_TOKEN_JWT>"
```

## Endpoints Principales

| Método | Endpoint               | Descripción                     | Rol Requerido |
|--------|------------------------|---------------------------------|---------------|
| POST   | /api/auth/register     | Registrar nuevo usuario         | -             |
| POST   | /api/auth/login        | Iniciar sesión                  | -             |
| GET    | /api/users             | Listar usuarios                 | admin         |
| GET    | /api/users/:id         | Obtener usuario por ID          | admin         |
| PUT    | /api/users/:id         | Actualizar usuario              | admin         |
| DELETE | /api/users/:id         | Eliminar usuario                | admin         |

## Middlewares Implementados

- **AuthMiddleware**: Verifica token JWT
- **RoleMiddleware**: Control de acceso por roles
- **LoggerMiddleware**: Registro de peticiones
- **RecoverMiddleware**: Manejo de panics

## Migración de Base de Datos

Los modelos se migran automáticamente al iniciar la aplicación. Para forzar una recreación:

```bash
docker-compose down -v
docker-compose up -d --build
```

## Documentación API

Para generar documentación Swagger (opcional):

1. Instalar swag:
```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

2. Generar docs:
```bash
swag init -g main.go
```

3. Acceder a:
```
http://localhost:8080/swagger/index.html
```

## Contribución

1. Haz fork del proyecto
2. Crea tu rama (`git checkout -b feature/nueva-funcionalidad`)
3. Haz commit de tus cambios (`git commit -am 'Añade nueva funcionalidad'`)
4. Haz push a la rama (`git push origin feature/nueva-funcionalidad`)
5. Abre un Pull Request

## Licencia

MIT License. Ver archivo [LICENSE](LICENSE) para más detalles.