

document.getElementById("ClickButton").addEventListener("click", function() {
    fetch("http://localhost:3000/endpoint", {
        method: "POST", 
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify({ points: 1 }) 
    })
    .then(response => response.json())
    .then(data => {
        console.log("Успех:", data);
    })
    .catch((error) => {
        console.error("Ошибка:", error);
    });
    
    
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

const music = document.getElementById('music');
const cnopka = document.getElementById('sound-button');
const icon = document.getElementById('icon');

let isPlaying = false;

function SwitchIcon(){
    if (isPlaying) {
        music.volume = 0.1
        icon.src = '../web/assets/sound-svgrepo-com.svg';
    } else {
        music.volume = 0
        icon.src = '../web/assets/sound-off-svgrepo-com.svg';
    }
    isPlaying = !isPlaying; 
}
cnopka.addEventListener('click', SwitchIcon);
const tracks = [
    '../web/assets/Disfigure - Blank.mp3',
    '../web/assets/DifferentHeaven - MyHeart.mp3',
    '../web/assets/Elektronomia - SkyHigh.mp3',
    ];
let x = 0;


const clickButn = document.getElementById('ClickButton');
clickButn.addEventListener('click', () => {
        x++;
        TogglePlay();
});
function TogglePlay(){
    if (x == 1) {
        const randomIndex = Math.floor(Math.random() * tracks.length);
        music.src = tracks[randomIndex];
        
        music.play()
        music.volume = 0.1
    } else{
        
    }
}
const audio = document.getElementById('music');

function loadRandomTrack() {
        const randomIndex = Math.floor(Math.random() * tracks.length);
        audio.src = tracks[randomIndex];
        audio.play();
}

audio.addEventListener('ended', loadRandomTrack);

