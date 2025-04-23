const navContainer = document.querySelector('.question-navigation');
const questionButtons = navContainer.querySelectorAll('.question-number');

// let isDragging = false;
// let startX;
// let scrollStart;
//
// questionButtons.forEach(button => {
//     button.addEventListener('click', (e) => {
//         button.classList.add('selected');
//     });
//
//     button.addEventListener('mousedown', (e) => {
//         isDragging = true;
//         startX = e.pageX;
//         scrollStart = navContainer.scrollLeft;
//         navContainer.classList.add('dragging');
//     });
//
//     button.addEventListener('mouseup', () => {
//         isDragging = false;
//         navContainer.classList.remove('dragging');
//     });
//
//     button.addEventListener('mouseleave', () => {
//         isDragging = false;
//         navContainer.classList.remove('dragging');
//     });
//
//     button.addEventListener('mousemove', (e) => {
//         if (!isDragging) return;
//         const x = e.pageX;
//         const walk = (x - startX) * 1.5;
//         navContainer.scrollLeft = scrollStart - walk;
//     });
// });



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

document.addEventListener('DOMContentLoaded', function() {
    // 1. Получаем ID теста из URL
    const pathParts = window.location.pathname.split('/').filter(Boolean);
    const testId = pathParts[pathParts.length - 1];

    if (!testId || isNaN(testId)) {
        console.error('Invalid test ID');
        return;
    }

    let testData = null;
    let userAnswers = {};
    let currentQuestionIndex = 0;
    let timerInterval = null;

    // 2. Модальное окно для начала теста
    function showStartModal() {
        const modal = document.createElement('div');
        modal.className = 'modal';
        modal.style.cssText = `
            position: fixed;
            top: 0;
            left: 0;
            width: 100%;
            height: 100%;
            background: rgba(0,0,0,0.5);
            display: flex;
            justify-content: center;
            align-items: center;
            z-index: 1000;
        `;

        modal.innerHTML = `
            <div style="background: white; padding: 20px; border-radius: 5px; max-width: 500px;">
                <h3>Готовы начать тест?</h3>
                <p>Тест содержит вопросы с одним или несколькими правильными ответами.</p>
                <div style="margin-top: 20px; text-align: right;">
                    <button id="start-test" style="padding: 8px 16px; background: #4CAF50; color: white; border: none; border-radius: 4px; cursor: pointer;">
                        Начать тест
                    </button>
                </div>
            </div>
        `;

        document.body.appendChild(modal);

        document.getElementById('start-test').addEventListener('click', function() {
            document.body.removeChild(modal);
            startTimer();
            showQuestion(0); // Показываем первый вопрос после старта
        });
    }

    // 3. Таймер
    function startTimer() {
        const timeElement = document.getElementById('time');
        let time = 30; // 30 минут в секундах

        timerInterval = setInterval(function() {
            time--;
            if (time <= 0) {
                clearInterval(timerInterval);
                finishTest();
                return;
            }

            const minutes = Math.floor(time / 60);
            const seconds = time % 60;
            timeElement.textContent = `${minutes}:${seconds < 10 ? '0' : ''}${seconds}`;
        }, 1000);
    }

    // 4. Загрузка данных теста
    async function loadTestData() {
        try {
            const response = await fetch(`/api/tests/${testId}`);
            if (!response.ok) throw new Error(`HTTP error! status: ${response.status}`);

            testData = await response.json();
            console.log('Test data loaded:', testData);

            if (!testData.questions || !testData.questions.length) {
                throw new Error('No questions in test');
            }

            document.querySelector('.test-title').textContent = testData.title;
            initializeTest();
            showStartModal();
        } catch (error) {
            console.error('Error loading test:', error);
            alert('Не удалось загрузить тест. Пожалуйста, попробуйте позже.');
        }
    }

    // 5. Инициализация теста
    function initializeTest() {
        const questionNav = document.querySelector('.question-navigation');
        questionNav.innerHTML = '';

        testData.questions.forEach((question, index) => {
            const label = document.createElement('label');
            label.className = 'question-label';
            if (index === 0) label.classList.add('active');

            label.innerHTML = `
                <input class="question-number" type="radio" name="question" value="${index}" ${index === 0 ? 'checked' : ''}>
                <span>${index + 1}</span>
            `;
            questionNav.appendChild(label);
        });

        // Обработчики для навигации
        document.querySelectorAll('.question-number').forEach(radio => {
            radio.addEventListener('change', function() {
                if (this.checked) {
                    saveCurrentAnswer();
                    currentQuestionIndex = parseInt(this.value);
                    showQuestion(currentQuestionIndex);
                }
            });
        });

        // Кнопка "Далее"
        const nextButton = document.getElementById('next-button');
        if (nextButton) {
            nextButton.addEventListener('click', function() {
                saveCurrentAnswer();
                if (currentQuestionIndex < testData.questions.length - 1) {
                    currentQuestionIndex++;
                    const nextRadio = document.querySelector(`.question-number[value="${currentQuestionIndex}"]`);
                    if (nextRadio) nextRadio.checked = true;
                    showQuestion(currentQuestionIndex);
                }
            });
        }
    }

    // 6. Показать вопрос
    function showQuestion(index) {
        if (!testData || !testData.questions || index >= testData.questions.length) {
            console.error('Invalid question index');
            return;
        }

        const question = testData.questions[index];

        // Обновляем текст вопроса
        const questionTextEl = document.querySelector('.question-text');
        if (questionTextEl) {
            questionTextEl.textContent = `${index + 1}. ${question.question_text}`;
        }

        // Обновляем варианты ответов
        const answersContainer = document.querySelector('.answers-box');
        if (answersContainer) {
            answersContainer.innerHTML = `
                <div class="answers-title">Варианты ответов (один ответ):</div>
                <ul class="answers-list">
                    ${question.answers.map(answer => `
                        <li>
                            <label>
                                <input type="radio" name="answer" value="${answer.id}"
                                    ${userAnswers[question.id] && userAnswers[question.id].includes(answer.id) ? 'checked' : ''}>
                                ${answer.answer_text}
                            </label>
                        </li>
                    `).join('')}
                </ul>
            `;
        }

        // Обновляем активные классы в навигации
        document.querySelectorAll('.question-label').forEach((label, i) => {
            label.classList.toggle('active', i === index);
        });

        // Обновляем состояние кнопки "Далее"
        const nextButton = document.getElementById('next-button');
        if (nextButton) {
            nextButton.style.display = index === testData.questions.length - 1 ? 'none' : 'block';
        }
    }

    // 7. Сохранить текущий ответ
    function saveCurrentAnswer() {
        if (!testData || currentQuestionIndex >= testData.questions.length) return;

        const questionId = testData.questions[currentQuestionIndex].id;
        const selected = document.querySelector('input[name="answer"]:checked');

        if (selected) {
            userAnswers[questionId] = [parseInt(selected.value)];
        } else {
            delete userAnswers[questionId];
        }

        console.log('Current answers:', userAnswers);
    }

    // 8. Завершение теста
    async function finishTest() {
        clearInterval(timerInterval);
        saveCurrentAnswer();

        const resultData = {
            result: testData.questions.map(question => ({
                question_id: question.id,
                answers_id: userAnswers[question.id] ? userAnswers[question.id][0] : null
            }))
        };

        console.log('Submitting results:', resultData);

        try {
            const response = await fetch('/sendresults', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(resultData)
            });

            if (response.ok) {
                const result = await response.json();
                alert('Тест завершен! Ваши результаты сохранены.');
                // Можно добавить перенаправление: window.location.href = '/results';
            } else {
                throw new Error(`Server returned ${response.status}`);
            }
        } catch (error) {
            console.error('Submission error:', error);
            alert('Ошибка при отправке результатов. Пожалуйста, попробуйте ещё раз.');
        }
    }

    // 9. Модальное окно подтверждения завершения
    window.forceFinishModal = function() {
        const modal = document.createElement('div');
        modal.className = 'modal';
        modal.style.cssText = `
            position: fixed;
            top: 0;
            left: 0;
            width: 100%;
            height: 100%;
            background: rgba(0,0,0,0.5);
            display: flex;
            justify-content: center;
            align-items: center;
            z-index: 1000;
        `;

        modal.innerHTML = `
            <div style="background: white; padding: 20px; border-radius: 5px; max-width: 500px;">
                <h3>Завершить тест?</h3>
                <p>Вы уверены, что хотите завершить тест? После завершения вы не сможете изменить ответы.</p>
                <div style="margin-top: 20px; text-align: right;">
                    <button id="confirm-finish" style="padding: 8px 16px; background: #f44336; color: white; border: none; border-radius: 4px; cursor: pointer; margin-right: 10px;">
                        Завершить
                    </button>
                    <button id="cancel-finish" style="padding: 8px 16px; background: #e7e7e7; border: none; border-radius: 4px; cursor: pointer;">
                        Отмена
                    </button>
                </div>
            </div>
        `;

        document.body.appendChild(modal);

        document.getElementById('confirm-finish').addEventListener('click', function() {
            document.body.removeChild(modal);
            finishTest();
        });

        document.getElementById('cancel-finish').addEventListener('click', function() {
            document.body.removeChild(modal);
        });
    };

    // Запускаем загрузку теста
    loadTestData();
});