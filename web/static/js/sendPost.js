// Функция для отправки POST-запроса
function sendPostRequest(url, data, successCallback) {
    return fetch(url, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(data)
    })
    .then(response => response.json())
    .then(result => {
        // Если есть поле "success" и оно истинно
        if (result.success) {
            // Если есть поле "info", показываем тост
            if (result.info) {
                showToast(result.info);
            }

            // Если есть callback-функция successCallback, передаем туда результат
            if (successCallback) {
                successCallback(result);
            }

            return result;
        } else if (result.error) {
            // Если есть поле "error", показываем тост с ошибкой
            showErrorToast(result.error);
            throw new Error(result.error);
        } else {
            throw new Error("Unexpected response format.");
        }
    })
    .catch(error => {
        // Обработка ошибок, отображение тоста с ошибкой
        showErrorToast(error.message);
        console.error("Request failed:", error);
    });
}

// Функция для отображения информационного тоста
function showToast(message) {
    const toastElement = document.getElementById('info-toast');
    const toastText = toastElement.querySelector('.toast-message');
    toastText.textContent = message;
    toastElement.classList.remove('hidden');
    
    // Автоматическое скрытие через 5 секунд
    setTimeout(() => {
        hideToast(toastElement);
    }, 5000);
}

// Функция для отображения тоста с ошибкой
function showErrorToast(message) {
    const errorToastElement = document.getElementById('error-toast');
    const errorToastText = errorToastElement.querySelector('.toast-message');
    errorToastText.textContent = message;
    errorToastElement.classList.remove('hidden');
    
    // Автоматическое скрытие через 5 секунд
    setTimeout(() => {
        hideToast(errorToastElement);
    }, 5000);
}

// Функция для скрытия тоста (можно использовать для закрытия вручную)
function hideToast(toastElement) {
    toastElement.classList.add('hidden');
}
