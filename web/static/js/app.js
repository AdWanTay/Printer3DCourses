function showForm(type) {
    const loginForm = document.getElementById('loginForm');
    const registerForm = document.getElementById('registerForm');
    const btnLogin = document.getElementById('btn-login');
    const btnRegister = document.getElementById('btn-register');

    if (type === 'login') {
        loginForm.style.display = 'block';
        registerForm.style.display = 'none';
        btnLogin.checked = true;
    } else {
        loginForm.style.display = 'none';
        registerForm.style.display = 'block';
        btnRegister.checked = true;
    }
}

document.addEventListener("DOMContentLoaded", () => {
    try {
        const loginBtn = document.getElementById("loginBtn");
        loginBtn.addEventListener("click", () => openAuthModal("login"));

        const registerBtn = document.getElementById("registerBtn");
        registerBtn.addEventListener("click", () => openAuthModal("register"));
    } catch {
    }

    try {
        const logoutBtn = document.querySelector(".logoutBtn");
        logoutBtn.addEventListener("click", () => logout());
    } catch {
    }
});


function openAboutModal(num) {
    // todo переменная num участвует в выборе инфы по конкретному курсу

    fetch("/web/templates/modals/about-course.html")
        .then((res) => res.text())
        .then((html) => {
            const modalContainer = document.createElement("div");
            modalContainer.innerHTML = html;
            document.body.appendChild(modalContainer);
            document.querySelector(".course-description").innerHTML = "<p>"+num+"</p>"
            //todo ВЕЗДЕ СДЕЛАТЬ ТАК = Закрытие по клику вне модалки (доп)
            modalContainer.addEventListener("click", (e) => {
                if (e.target.classList.contains("modal")) {
                    modalContainer.remove();
                }
            });
        });
}

function openAuthModal(type) {
    fetch("/web/templates/modals/auth.html")
        .then((res) => res.text())
        .then((html) => {
            const modalContainer = document.createElement("div");
            modalContainer.innerHTML = html;
            document.body.appendChild(modalContainer);

            // Запускаем нужную форму
            if (type === "login") {
                showForm("login");
            } else {
                showForm("register");
            }

            //todo ВЕЗДЕ СДЕЛАТЬ ТАК = Закрытие по клику вне модалки (доп)
            modalContainer.addEventListener("click", (e) => {
                if (e.target.classList.contains("modal")) {
                    modalContainer.remove();
                }
            });
            addAuthListener(type, modalContainer)
        });
}

function addAuthListener(type, modalContainer) {
    if (type === "login") {

        const loginForm = document.getElementById("loginForm");
        const loginHandler = login(modalContainer);
        loginForm.addEventListener('submit', loginHandler);
    } else {
        const registerForm = document.getElementById("registerForm");
        const registerHandler = register(modalContainer);
        registerForm.addEventListener('submit', registerHandler);
    }
}

function login(modalContainer) {
    return async (event) => {
        event.preventDefault();
        const form = event.target;
        const formData = {
            email: form.email.value,
            password: form.password.value,
        };

        try {
            const response = await fetch('/api/auth/login', {
                method: 'POST',
                headers: {'Content-Type': 'application/json'},
                body: JSON.stringify(formData)
            });
            const result = await response.json();

            if (response.ok) {
                modalContainer.remove();
                location.reload();
            } else {
                showErr('Ошибка входа: ' + (result.error || 'Неизвестная ошибка'))
            }
        } catch (error) {
            showErr('Ошибка при отправке запроса')
        }
    }
}

function register(modalContainer) {
    return async (event) => {
        event.preventDefault();
        const form = event.target;
        const formData = {
            lastName: form.lastName.value,
            firstName: form.firstName.value,
            patronymic: form.patronymic.value,
            phoneNumber: form.phoneNumber.value,
            email: form.email.value,
            password: form.password.value,
        };

        try {
            const response = await fetch('/api/auth/register', {
                method: 'POST',
                headers: {'Content-Type': 'application/json'},
                body: JSON.stringify(formData)
            });

            const result = await response.json();

            if (response.ok) {
                modalContainer.remove()
                location.reload();
            } else {
                showErr('Ошибка входа: ' + (result.error || 'Неизвестная ошибка'))
            }
        } catch (error) {
            console.error('Ошибка:', error);
            showErr('Ошибка при отправке запроса')
        }
    }
}

async function logout() {
    showNotify("Уведомление", "Началось!!");

    try {
        const response = await fetch("/api/auth/logout", {
            method: "GET",
            credentials: "include", // Обязательно, чтобы cookie ушла
        });

        if (response.ok) {
            window.location.href = "/";
        } else {
            showErr("Ошибка в запросе выхода. Попробуйте позже")
        }
    } catch (err) {
        showErr("Ошибка при запросе выхода:", err);
    }
}

// async function register(event) {
//     event.preventDefault();
//     const form = event.target;
//     const formData = {
//         lastName: form.lastName.value,
//         firstName: form.firstName.value,
//         patronymic: form.patronymic.value,
//         phoneNumber: form.phoneNumber.value,
//         email: form.email.value,
//         password: form.password.value,
//     };
//
//     try {
//         const response = await fetch('/api/auth/register', {
//             method: 'POST',
//             headers: {'Content-Type': 'application/json'},
//             body: JSON.stringify(formData)
//         });
//
//         const result = await response.json();
//
//         if (response.ok) {
//             alert('Регистрация прошла успешно!');
//             // document.getElementById('registerModal').style.display = 'none';
//             // form.reset();
//         } else {
//             alert('Ошибка регистрации: ' + (result.error || 'Неизвестная ошибка'));
//         }
//     } catch (error) {
//         console.error('Ошибка:', error);
//         alert('Ошибка при отправке запроса');
//     }
// }




// Функция для открытия модального окна с нужным содержимым
function openModal(config) {
    // Создаем модальное окно из шаблона
    const modalHtml = `
        <div class="modal">
            <div class="modal-content main-modal">
                <span class="modal-close" onclick="this.closest('.modal').remove()">✖</span>
                <h2 class="modal-title">${config.title}</h2>
                <div class="modal-body">${config.body}</div>
                <div class="modal-description">${config.description}</div>
                <div class="modal-actions">
                    <button id="mainBtn" class="modal-button primary">${config.mainBtnText}</button>
                    <button class="modal-button cancel" onclick="this.closest('.modal').remove()">Отмена</button>
                </div>
            </div>
        </div>
    `;

    // Вставляем модальное окно в body
    document.body.insertAdjacentHTML('beforeend', modalHtml);

    // Назначаем обработчик для основной кнопки
    document.getElementById('mainBtn').addEventListener('click', config.mainBtnAction);
}

// Конфигурации для разных модальных окон
const modalConfigs = {
    emailEdit: {
        title: "Изменение почты",
        body: `
            <form id="emailForm" autocomplete="off">
                <input type="email" id="newEmail" autocomplete="off" placeholder="Новая почта" required>
                <input type="password" id="currentPasswordForEmail" autocomplete="off" placeholder="Текущий пароль" required>
            </form>
        `,
        description: "После изменения почты вам нужно будет подтвердить новый адрес",
        mainBtnText: "Сохранить",
        mainBtnAction: async function() {
            const newEmail = document.getElementById('newEmail').value;
            const currentPassword = document.getElementById('currentPasswordForEmail').value;

            if (!newEmail || !currentPassword) {
                showNotify("Ошибка", "Все поля должны быть заполнены", "error");
                return;
            }

            try {
                const response = await fetch('/api/auth/change-email', {
                    method: 'PATCH',
                    headers: {
                        'Content-Type': 'application/json',
                        'Authorization': `Bearer ${localStorage.getItem('token')}`
                    },
                    body: JSON.stringify({
                        new_email: newEmail,
                        current_password: currentPassword
                    })
                });

                if (!response.ok) {
                    const error = await response.json();
                    throw new Error(error.message || "Не удалось изменить email");
                }

                showNotify("Успех", "Email успешно изменён. Проверьте новую почту для подтверждения.", "success");
                this.closest('.modal').remove();
            } catch (error) {
                showNotify("Ошибка", error.message, "error");
            }
        }
    },
    passwordEdit: {
        title: "Изменение пароля",
        body: `
            <form id="passwordForm">
                <input type="password" id="currentPassword" autocomplete="off" placeholder="Текущий пароль" required>
                <input type="password" id="newPassword" autocomplete="off" placeholder="Новый пароль" required>
                <input type="password" id="repeatNewPassword" autocomplete="off" placeholder="Повторите новый пароль" required>
            </form>
        `,
        description: "Пароль должен содержать не менее 8 символов",
        mainBtnText: "Сохранить",
        mainBtnAction: async function() {
            const currentPassword = document.getElementById('currentPassword').value;
            const newPassword = document.getElementById('newPassword').value;
            const repeatNewPassword = document.getElementById('repeatNewPassword').value;

            if (!currentPassword || !newPassword || !repeatNewPassword) {
                showNotify("Ошибка", "Все поля должны быть заполнены", "error");
                return;
            }

            if (newPassword !== repeatNewPassword) {
                showNotify("Ошибка", "Новые пароли не совпадают", "error");
                return;
            }

            if (newPassword.length < 8) {
                showNotify("Ошибка", "Пароль должен содержать не менее 8 символов", "error");
                return;
            }

            try {
                const response = await fetch('/api/auth/change-password', {
                    method: 'PATCH',
                    headers: {
                        'Content-Type': 'application/json',
                        'Authorization': `Bearer ${localStorage.getItem('token')}`
                    },
                    body: JSON.stringify({
                        new_password: newPassword,
                        current_password: currentPassword
                    })
                });

                if (!response.ok) {
                    const error = await response.json();
                    throw new Error(error.message || "Не удалось изменить пароль");
                }

                showNotify("Успех", "Пароль успешно изменён", "success");
                this.closest('.modal').remove();
            } catch (error) {
                showNotify("Ошибка", error.message, "error");
            }
        }
    },
    nameEdit: {
        title: "Изменение ФИО",
        body: `
            <form id="nameForm">
                <input type="text" id="newLastName" autocomplete="off" placeholder="Фамилия" required>
                <input type="text" id="newFirstName" autocomplete="off" placeholder="Имя" required>
                <input type="text" id="newPatronymic" autocomplete="off" placeholder="Отчество (необязательно)">
            </form>
        `,
        description: "Укажите ваши реальные фамилию и имя",
        mainBtnText: "Сохранить",
        mainBtnAction: async function() {
            const lastName = document.getElementById('newLastName').value;
            const firstName = document.getElementById('newFirstName').value;
            const patronymic = document.getElementById('newPatronymic').value;

            if (!lastName || !firstName) {
                showNotify("Ошибка", "Фамилия и имя обязательны для заполнения", "error");
                return;
            }

            try {
                const response = await fetch('/api/auth/change-name', {
                    method: 'PATCH',
                    headers: {
                        'Content-Type': 'application/json',
                        'Authorization': `Bearer ${localStorage.getItem('token')}`
                    },
                    body: JSON.stringify({
                        new_last_name: lastName,
                        new_first_name: firstName,
                        new_patronymic: patronymic
                    })
                });

                if (!response.ok) {
                    const error = await response.json();
                    throw new Error(error.message || "Не удалось изменить ФИО");
                }

                showNotify("Успех", "ФИО успешно изменены", "success");
                this.closest('.modal').remove();

                // Обновляем отображение имени в интерфейсе
                if (typeof updateUserDisplay === 'function') {
                    updateUserDisplay();
                }
            } catch (error) {
                showNotify("Ошибка", error.message, "error");
            }
        }
    },
    deleteAccount: {
        title: "Удаление аккаунта",
        body: `
            <form id="deleteForm">
                <input type="password" id="deletePassword" autocomplete="off" placeholder="Ваш пароль" required>
                <input type="text" id="deleteConfirmation" autocomplete="off" placeholder="Напишите 'УДАЛИТЬ'" required>
            </form>
        `,
        description: "Это действие необратимо. Все ваши данные будут удалены.",
        mainBtnText: "Удалить аккаунт",
        mainBtnAction: async function() {
            const password = document.getElementById('deletePassword').value;
            const confirmation = document.getElementById('deleteConfirmation').value;

            if (!password || confirmation !== 'УДАЛИТЬ') {
                showNotify("Ошибка", "Пожалуйста, введите пароль и напишите 'УДАЛИТЬ' для подтверждения", "error");
                return;
            }

            try {
                const response = await fetch('/api/auth/delete-account', {
                    method: 'DELETE',
                    headers: {
                        'Content-Type': 'application/json',
                        'Authorization': `Bearer ${localStorage.getItem('token')}`
                    },
                    body: JSON.stringify({
                        password: password
                    })
                });

                if (!response.ok) {
                    const error = await response.json();
                    throw new Error(error.message || "Не удалось удалить аккаунт");
                }

                showNotify("Успех", "Аккаунт успешно удалён", "success");
                this.closest('.modal').remove();

                // Перенаправляем на страницу входа
                setTimeout(() => {
                    window.location.href = '/login';
                }, 1500);
            } catch (error) {
                showNotify("Ошибка", error.message, "error");
            }
        }
    }
};

// Функция для открытия модального окна с нужным содержимым
function openModal(config) {
    // Создаем модальное окно из шаблона
    const modalHtml = `
        <div class="modal">
            <div class="modal-content main-modal">
                <span class="modal-close" onclick="this.closest('.modal').remove()">✖</span>
                <h2 class="modal-title">${config.title}</h2>
                <div class="modal-body">${config.body}</div>
                <div class="modal-description">${config.description}</div>
                <div class="modal-actions">
                    <button id="mainBtn" class="modal-button primary">${config.mainBtnText}</button>
                    <button class="modal-button cancel" onclick="this.closest('.modal').remove()">Отмена</button>
                </div>
            </div>
        </div>
    `;

    // Вставляем модальное окно в body
    document.body.insertAdjacentHTML('beforeend', modalHtml);

    // Назначаем обработчик для основной кнопки
    document.getElementById('mainBtn').addEventListener('click', config.mainBtnAction.bind(document.querySelector('.modal')));
}
// // Назначаем обработчики на кнопки
// document.getElementById('emailEdit').addEventListener('click', () => openModal(modalConfigs.emailEdit));
// document.getElementById('passwordEdit').addEventListener('click', () => openModal(modalConfigs.passwordEdit));
// document.getElementById('deleteAccount').addEventListener('click', () => openModal(modalConfigs.deleteAccount));