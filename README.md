# 🎵 F.Y.T.by Signal

## 🇫🇷 Version Française

F.Y.T.by Signal est une plateforme de vente de billets de concert en ligne.  
Ce repository contient le **backend développé en Go**, responsable de la gestion des utilisateurs, des artistes, des concerts et de l’authentification.

---

## 🚀 Stack technique

- **Langage** : Golang 1.25
- **Base de données** : PostgreSQL 16
- **Authentification** : JWT (JSON Web Token)
- **Sécurité** : Middleware JWT + Middleware CORS
- **Containerisation** : Docker & Docker Compose
- **Gestion BDD** : Migrations SQL
- **Outil de visualisation BDD** : DBeaver

---

## 🏗️ Architecture du projet

projet_gp_tracker_backend
|
|___cmd/
| |__api/ → Point d’entrée de l’API
| |__seed/ → Remplissage initial de la base
|
|___internal/
| |__database/ → Connexion et requêtes BDD
| |__handlers/ → Logique métier et routes
| |__middleware/→ JWT & CORS
|
|_migrations/ → Scripts SQL de création des tables
|
|.env → Variables d’environnement
|__docker-compose.yml → Orchestration des services
|__Dockerfile → Build de l’API


---

## 🔐 Sécurité

- Les mots de passe sont **hashés** avant stockage.
- Authentification via **JWT**.
- Middleware **CORS** pour limiter les accès à l’API.
- Variables sensibles stockées dans un fichier `.env`.

---

## 📦 Installation

### 1️⃣ Cloner le repository

```bash
git clone <repo_url>
cd projet_gp_tracker_backend

2️⃣ Créer un fichier .env

Créer un fichier .env en suivant le modèle du .env-exemple.

Exemple :

POSTGRES_USER=zeqzi
POSTGRES_PASSWORD=your_password
POSTGRES_DB=groupie_db
DATABASE_URL=postgres://zeqzi:your_password@db:5432/groupie_db
API_PORT=8080
JWT_SECRET=your_secret_key

▶️ Lancement du projet

Pour démarrer l’API et la base de données :

docker compose up --build


Pour arrêter et supprimer les volumes :

docker compose down -v


Le backend sera accessible sur :

http://localhost:8080

👤 Fonctionnalités principales

Inscription utilisateur

Connexion sécurisée

Modification du profil (pseudo & biographie)

Consultation des artistes

Consultation des concerts

Achat de billets

Visualisation des billets depuis le profil

🌐 Accès au Frontend

Ce repository contient uniquement le backend.

Pour accéder à l’interface utilisateur, vous devez également cloner le repository frontend du projet.

https://github.com/allanparis35/projet_gp_tracker.git

Le frontend possède son propre README.md qui explique comment l’installer et le lancer.

⚠️ Le backend doit être lancé avant le frontend.

/////////

## 🇬🇧 English Version

F.Y.T.by Signal is an online concert ticketing platform.
This repository contains the Go backend, responsible for managing users, artists, concerts, and authentication.

🚀 Tech Stack

Language: Golang 1.25

Database: PostgreSQL 16

Authentication: JWT (JSON Web Token)

Security: JWT Middleware + CORS Middleware

Containerization: Docker & Docker Compose

Database Management: SQL Migrations

Database Visualization Tool: DBeaver

🏗️ Project Architecture
projet_gp_tracker_backend
|
|___cmd/
|    |__api/        → API entry point
|    |__seed/       → Initial database seeding
|
|___internal/
|    |__database/   → Database connection & queries
|    |__handlers/   → Business logic & routes
|    |__middleware/ → JWT & CORS
|
|___migrations/     → SQL table creation scripts
|
|__.env                 → Environment variables
|__docker-compose.yml   → Service orchestration
|__Dockerfile           → API build configuration

🔐 Security

Passwords are hashed before being stored.

Authentication via JWT tokens.

CORS middleware to restrict API access.

Sensitive variables stored in a .env file.

📦 Installation
1️⃣ Clone the repository
git clone <repo_url>
cd projet_gp_tracker_backend

2️⃣ Create a .env file

Create a .env file following the .env-example template.

Example:

POSTGRES_USER=zeqzi
POSTGRES_PASSWORD=your_password
POSTGRES_DB=groupie_db
DATABASE_URL=postgres://zeqzi:your_password@db:5432/groupie_db
API_PORT=8080
JWT_SECRET=your_secret_key

▶️ Running the project

To start the API and the database:

docker compose up --build


To stop and remove volumes:

docker compose down -v


The backend will be available at:

http://localhost:8080

👤 Main Features

User registration

Secure login

Profile editing (username & bio)

Browse artists

Browse concerts

Purchase tickets

View purchased tickets in the profile

🌐 Frontend Access

This repository contains only the backend.

To access the user interface, you must also clone the project's frontend repository:

https://github.com/allanparis35/projet_gp_tracker.git


The frontend has its own README.md explaining how to install and run it.

⚠️ The backend must be running before starting the frontend.