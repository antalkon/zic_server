function TgOnlineBlock() {
    sendPostRequest('/setting/api/tg', {}, function(result) {
        const token = document.getElementById('tgToken');
        const stat = document.getElementById('system-status');
        const task = document.getElementById('tasks');
        console.log(result)
        token.value = result.token
        if (result.sysSt == true){
            stat.checked = true;
        }
        if (result.newTask == true){
            task.checked = true;
        }
        // Удаляем все классы, кроме баз
    });
}
function SendPutTgUpdate() {
    const token = document.getElementById('tgToken').value;
    const stat = document.getElementById('system-status').checked;
    const task = document.getElementById('tasks').checked;

    // Создаем объект данных для отправки
    const data = {
        token: token,
        sysSt: stat,     // Используем значение прямо из checkbox
        newTask: task    // Используем значение прямо из checkbox
    };

    // Отправляем PUT-запрос с данными
    sendPutRequest('/setting/api/tg', data, function(result) {
        if (result.success) {
            showToast("Telegram настройки успешно обновлены!");
        } else {
            showErrorToast("Ошибка при обновлении настроек Telegram.");
        }
    });
}

document.addEventListener("DOMContentLoaded", function() {
    TgOnlineBlock()

    // Повторно обновляем каждые 10 секунд (10000 миллисекунд)
    // setInterval(updatePcOnlineBlock, 10000);  // 10 секунд
    // setInterval(updateRoomOnlineBlock, 10000);  // 10 секунд
});