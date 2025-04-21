const navContainer = document.querySelector('.question-navigation');
const questionButtons = navContainer.querySelectorAll('.question-number');

let isDragging = false;
let startX;
let scrollStart;

questionButtons.forEach(button => {
    button.addEventListener('click', (e) => {
        button.classList.add('selected');
    });

    button.addEventListener('mousedown', (e) => {
        isDragging = true;
        startX = e.pageX;
        scrollStart = navContainer.scrollLeft;
        navContainer.classList.add('dragging');
    });

    button.addEventListener('mouseup', () => {
        isDragging = false;
        navContainer.classList.remove('dragging');
    });

    button.addEventListener('mouseleave', () => {
        isDragging = false;
        navContainer.classList.remove('dragging');
    });

    button.addEventListener('mousemove', (e) => {
        if (!isDragging) return;
        const x = e.pageX;
        const walk = (x - startX) * 1.5;
        navContainer.scrollLeft = scrollStart - walk;
    });
});



function forceFinishModal() {
    fetch("/web/templates/test/modals/force-finish.html")
        .then((res) => res.text())
        .then((html) => {
            const modalContainer = document.createElement("div");
            modalContainer.innerHTML = html;
            document.body.appendChild(modalContainer);

            //todo ВЕЗДЕ СДЕЛАТЬ ТАК = Закрытие по клику вне модалки (доп)
            modalContainer.addEventListener("click", (e) => {
                if (e.target.classList.contains("modal")) {
                    modalContainer.remove();
                }
            });
            addAuthListener(type, modalContainer)
        });
}