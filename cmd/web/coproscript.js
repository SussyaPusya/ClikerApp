document.getElementById("ClickButton").addEventListener("click", function() {
    fetch("http://localhost:3000/endpoint", {
        method: "POST", // или "GET" в зависимости от вашего сервера
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify({ points: 1 }) // данные, которые вы хотите отправить
    })
    .then(response => response.json())
    .then(data => {
        console.log("Успех:", data);
    })
    .catch((error) => {
        console.error("Ошибка:", error);
    });
    console.log("+1 govno")
});

const numberElement = document.getElementById('count');
const button = document.getElementById('ClickButton');

let count = 0;

button.addEventListener('click', () => {
    count++;
    numberElement.textContent = count;
});

const openModalButton = document.getElementById('openModalButton');
const closeModalButton = document.getElementById('closeModalButton');
const modal = document.getElementById('modal');

openModalButton.addEventListener('click', function() {
    modal.classList.add('open');
});

closeModalButton.addEventListener('click', function() {
    modal.classList.remove('open');
});

// Закрытие модального окна при клике вне его
window.addEventListener('click', function(event) {
    if (event.target === modal) {
        modal.classList.remove('open');
    }
});

