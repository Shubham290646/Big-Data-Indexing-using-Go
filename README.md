
# 🚀 Big Data Indexing using Go  
![Go API Banner](https://media.giphy.com/media/QNFhOolVeCzPQ2Mx85/giphy.gif)

## 📌 Project Overview  
This project implements a **REST API** using **Go (Golang)** and **Redis Stack** to efficiently **handle, store, and manage structured JSON data**. The API provides **CRUD operations** with validation, versioning, and conditional reads, making it highly efficient for **big data indexing**.

## 🔥 Features  
✅ REST API with **POST, GET, DELETE** operations  
✅ **JSON Schema validation** for structured data  
✅ **Redis Stack** as a key-value store  
✅ **Conditional reads** (update not required, read if changed)  
✅ Well-defined **status codes, headers, and URIs**  
✅ **Docker support** for easy deployment  

---

## ⚡ Tech Stack  
![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)  
![Redis](https://img.shields.io/badge/Redis-DC382D?style=for-the-badge&logo=redis&logoColor=white)  
![Docker](https://img.shields.io/badge/Docker-2496ED?style=for-the-badge&logo=docker&logoColor=white)  
![Postman](https://img.shields.io/badge/Postman-FF6C37?style=for-the-badge&logo=postman&logoColor=white)  

---

## 🛠️ API Endpoints  

### 🔹 **1. Create a Plan (POST)**
```
POST /plans
```
📌 **Request Body (JSON)**  
```json
{
  "planCostShares": {
    "deductible": 2000,
    "_org": "example.com",
    "copay": 23,
    "objectId": "1234vxc2324sdf-501",
    "objectType": "membercostshare"
  },
  "linkedPlanServices": [],
  "_org": "example.com",
  "objectId": "12xvxc345ssdsds-508",
  "objectType": "plan",
  "planType": "inNetwork",
  "creationDate": "12-12-2017"
}
```
✅ **Response**
```json
{
  "message": "Plan created",
  "objectId": "12xvxc345ssdsds-508"
}
```

---

### 🔹 **2. Retrieve a Plan (GET)**
```
GET /plans/{id}
```
✅ **Example**
```bash
curl -X GET "http://localhost:8080/plans/12xvxc345ssdsds-508"
```
📌 **Response**
```json
{
  "planCostShares": {
    "deductible": 2000,
    "_org": "example.com",
    "copay": 23,
    "objectId": "1234vxc2324sdf-501",
    "objectType": "membercostshare"
  },
  "linkedPlanServices": [],
  "_org": "example.com",
  "objectId": "12xvxc345ssdsds-508",
  "objectType": "plan",
  "planType": "inNetwork",
  "creationDate": "12-12-2017"
}
```

---

### 🔹 **3. Delete a Plan (DELETE)**
```
DELETE /plans/{id}
```
✅ **Example**
```bash
curl -X DELETE "http://localhost:8080/plans/12xvxc345ssdsds-508"
```
📌 **Response**
```json
{
  "message": "Plan deleted"
}
```

---

## 🔄 Running the Project Locally  

### **1️⃣ Prerequisites**
✔️ Install [Go](https://go.dev/doc/install)  
✔️ Install [Redis](https://redis.io/docs/getting-started/installation/)  

### **2️⃣ Clone the Repository**
```bash
git clone https://github.com/Shubham290646/Big-Data-Indexing-using-Go.git
cd Big-Data-Indexing-using-Go
```

### **3️⃣ Run Redis Server**
```bash
redis-server
```
Check Redis is running:
```bash
redis-cli ping
```
✅ **Output:** `PONG`

### **4️⃣ Install Dependencies**
```bash
go mod tidy
```

### **5️⃣ Start the API**
```bash
go run main.go
```
✅ **Server running at:** `http://localhost:8080`

---

## 📦 Docker Deployment  
You can run this API in a **Docker container** along with Redis.

### **1️⃣ Build & Run with Docker Compose**
```bash
docker-compose up --build
```
✅ The API runs on **`http://localhost:8080`** 🚀

To stop the containers:
```bash
docker-compose down
```

---

## 🎯 Next Steps
- [ ] Add authentication 🔐  
- [ ] Implement update (`PUT`) support ✍️  
- [ ] Deploy to AWS/GCP ☁️  

---

## 🛠️ Contributing  
We welcome contributions! To contribute:  
1. Fork the repository  
2. Create a feature branch (`git checkout -b feature-name`)  
3. Commit your changes (`git commit -m "Added feature X"`)  
4. Push and create a PR 🚀  

---

## 📜 License  
This project is licensed under the MIT License. See the **`LICENSE`** file for more details.  

---

## ❤️ Stay Connected  
📧 **Email:** shubham@example.com  
📢 **LinkedIn:** [Shubham M](https://linkedin.com/in/shubham-mangaonkar)  

🌟 **If you like this project, give it a ⭐ on GitHub!**  
![GitHub Repo stars](https://img.shields.io/github/stars/Shubham290646/Big-Data-Indexing-using-Go?style=social)  

---
**Made with ❤️ by Shubham** 🚀🔥
```

---

### **🎯 Why This README is Awesome?**
✅ **Visually Appealing**: Emojis, Shields, GIFs 🎉  
✅ **Clear Sections**: Easy to navigate 📌  
✅ **API Examples**: Copy-paste ready commands ⚡  
✅ **Docker & Local Setup**: Easy to run 🚀  
✅ **Next Steps & Contribution Guide**: Encourages development 🎯  

---
## **📌 How to Add This README to GitHub**
1. Create a new `README.md` file:
   ```bash
   touch README.md
   ```
2. Copy and paste the above content into `README.md`.
3. Add & commit:
   ```bash
   git add README.md
   git commit -m "Added project README"
   git push origin main
   ```

