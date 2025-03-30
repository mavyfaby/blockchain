# Technical Specifications

## Overview

This document outlines the technical specifications for the project.
It includes details about the architecture, components, and technologies used in the development process.

## Architecture

The architecture of the project is based on a microservices approach.
This allows for better scalability, maintainability, and deployment flexibility.

The main components of the architecture are:

- **Frontend**:
  - The user interface of the application, built using Nuxt as the framework for Web and Flutter for mobile app.
- **Backend**:
  - The server-side logic, built using Echo for the Go programming language.
- **Database**:
  - The data storage layer, using MariaDB as the database management system with `sqlc` for generating type-safe SQL queries.
- **The blockchain**:
  - The heart and core of the project, which manages the data and transactions in a decentralized manner.

## Components

The project consists of the following main components:

### Frontend

- **Nuxt**
  - Is a free and open source JavaScript library based on Vue.js, Nitro, and Vite.
  - Inspired by Next.js, which is a similar framework based on React rather than Vue.

- **Flutter**
  - Flutter is an open-source UI software development kit created by Google.
  - It can be used to develop cross-platform applications from a single codebase for the
    web, Fuchsia, Android, iOS, Linux, macOS, and Windows.

### Backend

- **Go**
  - Go is a high-level general purpose programming language that is statically typed and compiled.
  - It is designed for simplicity and efficiency, making it a popular choice for building web applications and microservices.
  
- **Echo**
  - Echo is a high-performance web framework for building robust and scalable applications in Go.

- **sqlc**
  - `sqlc` is a tool that generates type-safe SQL queries in Go.
  - It allows developers to write SQL queries and automatically generates the corresponding Go code,
    ensuring type safety and reducing the risk of runtime errors.

### Database

- **MariaDB**
  - MariaDB is a community-developed, commercially supported fork of the MySQL relational database management system (RDBMS).
  - Used as the database management system for the project for storing and managing user data, transactions, and other application-related information.

## References

- [Nuxt - Wikipedia](https://en.wikipedia.org/wiki/Nuxt)
- [Flutter (software) - Wikipedia](https://en.wikipedia.org/wiki/Flutter_(software))
- [Go (programming language) - Wikipedia](https://en.wikipedia.org/wiki/Go_(programming_language))
- [High performance, extensible, minimalist Go web framework | Echo](https://echo.labstack.com/)
- [Compile SQL to type-safe code | sqlc.dev](https://sqlc.dev/)
- [MariaDB - Wikipedia](https://en.wikipedia.org/wiki/MariaDB)