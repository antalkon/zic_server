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
        showToast('success', `${message}`);

      } else {
      }
    } catch (error) {
      console.error("Ошибка при выполнении запроса:", error);
      showToast('error', 'Ошибка при выполнении запроса.');

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

        if (response.ok) {
            // Если успех, показываем тост с сообщением "Успешно"
            showToast('success', `${result.message}`);

        } else {
            // В противном случае показываем сообщение о неудаче
            showToast('error', 'Не удалось определить статус выполнения задачи.');

        }
    } catch (error) {
        showToast('error', 'Ошибка при выполнении запроса.');

        console.error("Ошибка при выполнении запроса:", error);
    } finally {
        // Скрываем экран блокировки после завершения запроса
        hideLockScreen();
    }
}



// Функции показа и скрытия экрана блокировки
function showLockScreen() {
    document.getElementById("lockScreen").classList.remove("hidden");
}

function hideLockScreen() {
    document.getElementById("lockScreen").classList.add("hidden");
}
