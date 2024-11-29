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

        if (response.ok) {
            // Если успех, показываем тост с сообщением "Успешно"
            showToast('success', `${result.message}`);
        } else {
            // В противном случае показываем сообщение о неудаче
            showToast('info', `${result.message}`);
        }
    } catch (error) {
        console.error("Ошибка при выполнении запроса:", error);
        showToast('error', `${"Ошибка при выполнении запроса"}`);
    } finally {
        // Скрываем экран блокировки после завершения запроса
    }
}


function openModal(modalId) {
    document.getElementById(modalId).classList.remove('hidden');
  }

  // Функция для закрытия модального окна
  function closeModal(modalId) {
    document.getElementById(modalId).classList.add('hidden');
  }

  async function submitUrlForm(id) {
    const urlInput = document.getElementById("input-url").value;
    const data = { url: urlInput };

    const url = `/pc/api/link/${id}`;
    const response = await postRequestNew(url, data);
    closeModal('url-modal')

    if (response) {
        setTimeout(() => {
            location.reload();
        }, 2000);
    }
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