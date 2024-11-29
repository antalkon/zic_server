
async function postRequestNew(url, body) {
      try {
          const response = await fetch(url, {
              method: 'POST',
              headers: {
                  'Content-Type': 'application/json',
              },
              body: JSON.stringify(body),
          });
  
          // Проверяем статус ответа
          if (!response.ok) {
              showToast('error', `Ошибка: ${response.json().error}`);
              return null;
          }
  
          // Преобразуем ответ в JSON
          const data = await response.json();
  
          showToast('success', `${data.message}`);
          return data;
      } catch (error) {
          showToast('error', 'Ошибка при выполнении запроса.');
          console.error('Ошибка во время POST-запроса:', error);
          return null;
      }
  }