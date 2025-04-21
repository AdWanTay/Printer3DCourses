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
