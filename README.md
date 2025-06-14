[![Build Status](https://github.com/ishtiaqhimel/cms/workflows/CMS%20CI/badge.svg)](https://github.com/ishtiaqhimel/cms/actions?workflow=CMS%20CI)

## News Portal CMS

To deliver a fully functional content management backend that enables 
journalists and editors to create, manage, and publish news articles.

---

## Detailed Features

--TODO--

---

## How to Run Locally

To run this project locally, make sure you have the following installed on your machine:
- [Git](https://git-scm.com/)
- [Docker](https://www.docker.com/) (make sure you have `docker compose`)

### Steps

1. Clone the repository:
```bash
git clone git@github.com:ishtiaqhimel/cms.git
cd cms
```

2. Start the application:
```bash
make serve
```

---

The project is implemented by following [The Clean Code Blog by Robert C. Martin (Uncle Bob)](https://blog.cleancoder.com/uncle-bob/2011/11/22/Clean-Architecture.html)

---

## Tools Used
- [Go](https://go.dev/)
- [Postgres](https://www.postgresql.org/)
- [Minio](https://min.io/)
- [Docker](https://www.docker.com/)
- [Consul](https://developer.hashicorp.com/consul)
- All libraries listed in [go.mod](/go.mod)