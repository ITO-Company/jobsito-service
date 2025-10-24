# Jobsito Service

Servicio principal para ITO-Company.

## Requisitos

- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)
- (Opcional para desarrollo) [Go 1.24+](https://golang.org/dl/)

## Variables de entorno

Crea un archivo `.env` en la raíz del proyecto con al menos:

```
DATABASE_URL=postgres://usuario:password@host:port/dbname?sslmode=require
PORT=8000
ALLOW_ORIGINS=*
```

Asegúrate de que `DATABASE_URL` apunte a tu base de datos (por ejemplo, Neon).

---

## Levantar en Desarrollo

Usa Docker Compose para desarrollo, con hot reload y montando el código local:

```sh
docker compose -f docker-compose.dev.yml up --build
```

Esto:
- Usa `Dockerfile.dev`
- Expone el puerto 8000
- Recarga la app automáticamente al guardar cambios gracias a [Air](https://github.com/cosmtrek/air)
- Monta tu código local dentro del contenedor

**¿Debo reconstruir el contenedor cada vez que cambio el código?**  
No. Air detecta los cambios y recarga la app automáticamente.  
Solo usa `--build` si cambias dependencias (`go.mod`, `go.sum`) o el Dockerfile.

**¿Qué hago si apago mi máquina?**  
Cuando la prendas de nuevo, solo ejecuta:

```sh
docker compose -f docker-compose.dev.yml up
```

No necesitas `--build` a menos que hayas cambiado dependencias o el Dockerfile.

Puedes acceder a la app en [http://localhost:8000](http://localhost:8000).

---

## Levantar en Producción

Para producción, usa el Dockerfile principal:

```sh
docker build -t jobsito-service .
docker run --env-file .env -p 8000:8000 jobsito-service
```

Esto:
- Usa `Dockerfile` para construir un binario optimizado
- Solo necesitas el archivo `.env` con tus variables

---

## Notas

- El servicio lee la variable `PORT` para saber en qué puerto escuchar.
- Si usas Render, el puerto se inyecta automáticamente.
- Si necesitas migraciones o seeds, revisa la función `Migrate` en `config/config.go`.

---

## Comandos útiles

- **Detener los contenedores:**  
  `docker compose -f docker-compose.dev.yml down`

- **Ver logs:**  
  `docker logs cuent_ai_core`

---

## Estructura principal

- `cmd/main.go`: Punto de entrada de la app
- `config/`: Configuración y conexión a la base de datos
- `Dockerfile`, `Dockerfile.dev`: Imágenes para prod y dev
- `docker-compose.dev.yml`: Orquestación para desarrollo

---

¿Dudas? Abre un issue o contacta al equipo.