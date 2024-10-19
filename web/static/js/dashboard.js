function updatePcOnlineBlock() {
    sendPostRequest('http://localhost:8385/pc/api/count', {}, function(result) {
        const countElement = document.getElementById('pc-online-count');
        const circleElement = document.getElementById('pc-online-circle');

        // Удаляем все классы, кроме базовых
        countElement.className = 'text-3xl font-bold'; // Базовые классы для текста
        circleElement.className = 'w-3 h-3 rounded-full pulse mr-2'; // Базовые классы для кружка

        // Динамически добавляем цвет текста и фона
        countElement.classList.add(`text-${result.color}`);  // Для текста, динамически добавляем класс цвета
        circleElement.classList.add(`bg-${result.color}`);    // Для кружка, динамически добавляем класс фона

        // Обновляем текст
        countElement.textContent = `${result.on} из ${result.count}`;
    });
}

function updateRoomOnlineBlock() {
    sendPostRequest('http://localhost:8385/room/api/count', {}, function(result) {
        const countElement = document.getElementById('room-online-count');
        const circleElement = document.getElementById('room-online-circle');  // Исправлено 'rooom' на 'room'

        // Удаляем все классы, кроме базовых
        countElement.className = 'text-3xl font-bold'; // Базовые классы для текста
        circleElement.className = 'w-3 h-3 rounded-full pulse mr-2'; // Базовые классы для кружка

        // Динамически добавляем цвет текста и фона
        countElement.classList.add(`text-${result.color}`);  // Для текста, динамически добавляем класс цвета
        circleElement.classList.add(`bg-${result.color}`);    // Для кружка, динамически добавляем класс фона

        // Обновляем текст
        countElement.textContent = `${result.on} из ${result.count}`;
    });
}

function updateStatOnlineBlock() {
    sendPostRequest('http://localhost:8385/setting/api/sys/status', {}, function(result) {
        const countElement = document.getElementById('stat-online-count');
        const circleElement = document.getElementById('stat-online-circle');  // Исправлено 'rooom' на 'room'

        // Удаляем все классы, кроме базовых
        countElement.className = 'text-3xl font-bold'; // Базовые классы для текста
        circleElement.className = 'w-3 h-3 rounded-full pulse mr-2'; // Базовые классы для кружка

        // Динамически добавляем цвет текста и фона
        countElement.classList.add(`text-${result.color}`);  // Для текста, динамически добавляем класс цвета
        circleElement.classList.add(`bg-${result.color}`);    // Для кружка, динамически добавляем класс фона

        // Обновляем текст
        countElement.textContent = `${result.system_status}`;
    });
}

// Вызываем функцию сразу после загрузки страницы
document.addEventListener("DOMContentLoaded", function() {
    updatePcOnlineBlock(); // Обновление при загрузке страницы
    updateRoomOnlineBlock();
    updateStatOnlineBlock();

    // Повторно обновляем каждые 10 секунд (10000 миллисекунд)
    // setInterval(updatePcOnlineBlock, 10000);  // 10 секунд
    // setInterval(updateRoomOnlineBlock, 10000);  // 10 секунд
});

function refresh() {
    showOverlay();  // Показываем overlay
    updatePcOnlineBlock();  // Обновляем блоки
    updateRoomOnlineBlock();
    updateStatOnlineBlock();
    
    // Ждём 2 секунды (2000 миллисекунд), затем скрываем overlay
    setTimeout(function() {
        hideOverlay();
    }, 1500);
}

// Функция для показа блока (меняет hidden на flex)
function showOverlay() {
    document.getElementById('overlay').classList.remove('hidden');
    document.getElementById('overlay').classList.add('flex');
}

// Функция для скрытия блока (меняет flex на hidden)
function hideOverlay() {
    document.getElementById('overlay').classList.remove('flex');
    document.getElementById('overlay').classList.add('hidden');
}
