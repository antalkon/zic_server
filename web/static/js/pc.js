async function controlPC(pcId, action) {
    const url = `/pc/api/${action}/${pcId}`;
    
    // Показать экран блокировки

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

function openModal(modalId) {
    document.getElementById(modalId).classList.remove('hidden');
  }

  // Функция для закрытия модального окна
  function closeModal(modalId) {
    document.getElementById(modalId).classList.add('hidden');
  }

  function submitUrlForm(id) {
    const urlInput = document.getElementById("input-url").value;

    const data = { "url": urlInput };

    fetch(`/pc/api/link/${id}`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json"
      },
      body: JSON.stringify(data)
    })
    .then(response => {
      if (response.ok) {
        alert("URL отправлен успешно");
        document.getElementById("url-form").reset(); // Сброс формы
        closeModal('url-modal'); // Закрываем модальное окно
      } else {
        alert("Ошибка при отправке URL:", response.result.error);
      }
    })
    .catch(error => {
      console.error("Ошибка:", error);
      alert("Ошибка при отправке URL");
    });
  }


  function submitEditPcForm(id) {
    const name = document.getElementById("pc-name").value;
    const localIp = document.getElementById("local-ip").value;
    const publicIp = document.getElementById("public-ip").value;

    const data = {
        "id": id,
        "name": name,
        "lip": localIp,
        "pip": publicIp
    };

    fetch("/pc/api/edit", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(data)
    })
    .then(response => {
        if (response.ok) {
            alert("Данные ПК успешно обновлены");
            document.getElementById("edit-pc-form").reset();
            closeModal('edit-pc-modal');
        } else {
            alert("Ошибка при обновлении данных ПК");
        }
    })
    .catch(error => {
        console.error("Ошибка:", error);
        alert("Ошибка при обновлении данных ПК");
    });
}