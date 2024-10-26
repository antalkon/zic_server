// URL для получения версии
const url = '/setting/api/info';

function fetchAndCacheVersion() {
  // Проверяем, есть ли версия в localStorage
  let cachedVersion = localStorage.getItem('version');
  let version;

  if (cachedVersion) {
    // Если версия найдена в кэше, используем её сразу
    version = cachedVersion;
    console.log("Cached version:", version);
  }

  // Делаем запрос к API для получения актуальной версии
  fetch(url)
    .then(response => {
      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }
      return response.json(); // Парсим ответ как JSON
    })
    .then(data => {
      if (data.version && data.version !== version) {
        // Обновляем локальную переменную и localStorage
        version = data.version;
        localStorage.setItem('version', version);
        console.log("Updated version from API:", version);
      } else {
        console.log("Version is up-to-date, no changes required.");
      }
    })
    .catch(error => {
      console.error("Fetch error:", error);
    });
  
  return version; // Возвращаем текущую версию из локального хранилища или undefined
}

// Вызов функции
let currentVersion = fetchAndCacheVersion();
const version = document.getElementById('version')
version.textContent = currentVersion
console.log("Current version:", currentVersion);
