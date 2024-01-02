# Microservices Project Repository

## Overview

Welcome to the Microservices Project Repository! This repository is a collection of projects centered around building scalable and modular microservices using Golang and Protocol Buffers. The primary components of this repository include an Event-Driven Microservice, a Pub&Sub system, an Encryption service, and an RPC client for seamless communication.

## Projects

### 1. Event-Driven Microservice (Pub&Sub)

This project focuses on creating a scalable and event-driven microservice architecture. It leverages Golang and Protocol Buffers for efficient communication between services. The Pub&Sub system allows seamless event handling and communication between different microservices.

- **Technology Stack:**
  - Golang
  - Protocol Buffers
  - Event-Driven Architecture

### 2. Encryption Service

The Encryption Service is designed to receive a string, encrypt it using the AES algorithm with Base64 encoding, and provide the capability to decrypt it on demand. This service ensures secure communication and data protection within the microservices ecosystem.

- **Features:**
  - String encryption using AES
  - Base64 encoding for secure data representation
  - On-demand decryption

- **Technology Stack:**
  - Golang
  - AES Encryption
  - Base64 Encoding

### 3. RPC Client for Encryption Service

To facilitate easy interaction with the Encryption Service, an RPC client has been developed. This client allows sending strings to the Encryption Service for encryption and retrieving decrypted strings as needed.

- **Features:**
  - Seamless communication with the Encryption Service
  - Simple API for sending and receiving encrypted strings

- **Technology Stack:**
  - Golang
  - Remote Procedure Call (RPC)

## Getting Started

To get started with these projects, follow the instructions provided in each project's respective directory. Ensure that you have the necessary dependencies installed, and the services are running as intended.

## Usage

Detailed usage guidelines for each project can be found in their respective directories. Make sure to refer to the documentation for each component to integrate them effectively into your microservices architecture.

## Contributions

Contributions are welcome! If you have any ideas, improvements, or bug fixes, feel free to open an issue or submit a pull request. Your contributions will help enhance the functionality and robustness of this microservices project.

## License

This repository is licensed under the [MIT License](LICENSE). Feel free to use, modify, and distribute the code as needed. Please refer to the license file for more details.

Happy coding!
