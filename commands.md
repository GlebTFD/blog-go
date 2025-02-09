### **1. Создание поста**
```bash
curl -X POST -H "Content-Type: application/json" -d '{
    "title": "Мой первый пост",
    "content": "Это тестовый пост",
    "author_id": 1
}' http://localhost:8080/posts
```

### **2. Получение всех постов**
```bash
curl http://localhost:8080/posts
```

### **3. Получение поста по ID**
```bash
curl http://localhost:8080/posts/1
```

### **4. Обновление поста**
```bash
curl -X PUT -H "Content-Type: application/json" -d '{
    "title": "Обновленный пост",
    "content": "Новое содержимое"
}' http://localhost:8080/posts/1
```

### **5. Удаление поста**
```bash
curl -X DELETE http://localhost:8080/posts/1
```
