// Функция отправки POST запроса и показа результатов
async function changeRoomStatus(roomId, action) {
    const url = `/room/api/room/${action}/${roomId}`;
    
    // Показать экран блокировки
    showLockScreen();
  
    try {
      const response = await fetch(url, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        }
      });
  
      // Получаем JSON-ответ от сервера
      const result = await response.json();

      if (response.ok) {
        // Формируем сообщение на основе ответа
        const message = `Успешно: ${result.success} | Неудачно: ${result.bad}`;
        showToast(message); // Показать тост с сообщением
      } else {
        showToast("Ошибка при выполнении операции. Попробуйте снова.");
      }
    } catch (error) {
      console.error("Ошибка при выполнении запроса:", error);
      showToast("Ошибка при соединении с сервером.");
    } finally {
      // Скрываем экран блокировки после завершения запроса
      hideLockScreen();
    }
  }

// Обновленная функция для отправки POST-запроса и показа результатов в controlPC
async function controlPC(pcId, action) {
    const url = `/pc/api/${action}/${pcId}`;
    
    // Показать экран блокировки
    showLockScreen();

    try {
        const response = await fetch(url, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            }
        });

        // Получаем JSON-ответ от сервера
        const result = await response.json();

        if (response.ok && result.success === true) {
            // Если успех, показываем тост с сообщением "Успешно"
            showToast("Успешно");
        } else {
            // В противном случае показываем сообщение о неудаче
            showToast("Не удалось узнать у ПК статус выполнения задачи");
        }
    } catch (error) {
        console.error("Ошибка при выполнении запроса:", error);
        showToast("Ошибка при соединении с сервером.");
    } finally {
        // Скрываем экран блокировки после завершения запроса
        hideLockScreen();
    }
}

// Функция показа тоста с сообщением
function showToast(message) {
    const toast = document.getElementById("toast");
    const toastMessage = document.getElementById("toastMessage");
    
    toastMessage.textContent = message;
    toast.classList.remove("hidden");

    // Автоматическое скрытие тоста через 3 секунды
    setTimeout(() => {
        hideToast();
    }, 3000);
}

// Функция скрытия тоста
function hideToast() {
    const toast = document.getElementById("toast");
    toast.classList.add("hidden");
}

// Функции показа и скрытия экрана блокировки
function showLockScreen() {
    document.getElementById("lockScreen").classList.remove("hidden");
}

function hideLockScreen() {
    document.getElementById("lockScreen").classList.add("hidden");
}
