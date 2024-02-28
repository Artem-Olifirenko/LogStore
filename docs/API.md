## Использование

Чтобы начать работу с LogStore в вашем проекте на Go, выполните следующие шаги:

1. **Импорт пакета LogStore**

   Добавьте LogStore в ваш проект, импортировав его:

   ```go
   import "github.com/Artem-Olifirenko/LogStore/store"
   ```
2. **Создание экземпляра хранилища**

   Инициализируйте новый экземпляр хранилища:

   ```go
   storage := store.NewStore()
   ```
3. **Добавление элемента с TTL**

   Для добавления элемента с определенным временем жизни (TTL) используйте метод `Set`:

    ```go
   storage.Set("myKey", "myValue", 30*time.Second)
   ```
    
   В этом примере элемент с ключом `"myKey"` и значением `"myValue"` будет храниться в течение 30 секунд.

4. **Получение элемента**

   Получите элемент по ключу с помощью метода Get:

   ```go
   value, found := storage.Get("myKey")
   if found {
        fmt.Println("Найдено значение:", value)
    } else {
        fmt.Println("Значение не найдено или срок его действия истек.")
    }
   ```

   Этот код проверяет наличие элемента и выводит его значение, если элемент найден и его TTL не истек.
   
5. **Удаление элемента**

   Удалите элемент из хранилища по ключу с помощью метода `Delete`:

   ```go
   storage.Delete("myKey")
   ```

6. Пример реализации

   Импортируйте пакет и создайте экземпляр хранилища:

    ```go
    package main
    
    import (
        "fmt"
        "time"
        "github.com/Artem-Olifirenko/LogStore/store" // Импорт пакета LogStore
    )
    
    func main() {
        storage := store.NewStore() // Создание экземпляра хранилища
    
        // Добавление элемента с TTL
        storage.Set("myKey", "myValue", 30*time.Second)
    
        // Получение элемента
        value, found := storage.Get("myKey")
        if found {
            fmt.Println("Найдено значение:", value)
        } else {
            fmt.Println("Значение не найдено или срок его действия истек.")
        }
    
        // Удаление элемента
        storage.Delete("myKey")
    }
