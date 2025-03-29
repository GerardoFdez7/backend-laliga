# ⚽ LaLiga Tracker

🌐 **LaLiga Tracker** es una API desarrollada en **Golang** que centraliza la administración de encuentros deportivos, ofreciendo operaciones completas de creación, consulta, modificación y eliminación de partidos.

![Demo de la API](assets/crear.jpg)

---

## 🔥 Funcionalidades Principales

🚩 **Gestión completa de partidos**  
   - Registra nuevos encuentros deportivos  
   - Consulta historial de partidos con filtros  
   - Actualiza información en tiempo real  
   - Elimina registros obsoletos  

🗃️ **Almacenamiento robusto**  
   - Motor PostgreSQL para gestión de datos  
   - Modelado relacional optimizado  
   - Escalabilidad garantizada  

🌍 **Acceso multiplataforma**  
   - Configuración CORS para integración frontal  
   - Compatibilidad con cualquier cliente REST  

---

## 🛠 Stack Tecnológico

### 🔧 Backend
- **Lenguaje**: Go 1.20+  
- **Framework**: Gin Gonic   

### 🗃️ Persistencia
- **Base de datos**: PostgreSQL 15  
- **ORM**: GORM v2  

---

## 🚀 Configuración Inicial

### 🐳 Docker Compose
```bash
# Clonar repositorio
git clone https://github.com/GerardoFdez7/backend-laliga.git
cd backend-laliga

# Levantar servicios con Docker
docker-compose up -d --build
```

### 🌐 Frontend
```bash
start LaLigaTracker.html
```

### 📚 Documentación API
[![Colección Postman](https://img.shields.io/badge/Postman-Colección_API-FF6C37?logo=postman&style=flat)](https://web.postman.co/workspace/de223e66-77d9-4d3d-87f3-aee8ba343d79/collection/41407945-ccd3f832-1d8d-4789-9d94-c4b3b2b83b54?share=true&origin=tab-menu)  
Accede a la documentación completa de endpoints y prueba la API directamente en Postman.
