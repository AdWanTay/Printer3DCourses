// ################################################################################################################
// ##                    Notifications
// ################################################################################################################

function showErr(text) {
    new Notify ({
        status: 'error',
        title: 'Ошибка',
        text: `${text}.`,
        effect: 'fade',
        speed: 300,
        showIcon: true,
        showCloseButton: true,
        autoclose: true,
        autotimeout: 10000,
        type: 'outline',
        position: `right top`,
    })
};

function showNotify(title, text) {
    new Notify ({
        status: 'success',
        title: `${title}.`,
        text: `${text}.`,
        effect: 'fade',
        speed: 300,
        customClass: '',
        customIcon: '',
        showIcon: true,
        showCloseButton: true,
        autoclose: true,
        autotimeout: 3000,
        notificationsGap: null,
        notificationsPadding: null,
        type: 'outline',
        position: 'right top',
    })
};


// ################################################################################################################
// ##                    Themes
// ################################################################################################################

document.addEventListener("DOMContentLoaded", () => {
    function detectColorScheme(){
        var theme="light";    //default to light

        if(localStorage.getItem("theme")){
            if(localStorage.getItem("theme") == "dark"){
                var theme = "dark";
            }
        } else if(!window.matchMedia) {
            return false;
        } else if(window.matchMedia("(prefers-color-scheme: dark)").matches) {
            var theme = "dark";
        }

        if (theme=="dark") {
            document.documentElement.setAttribute("data-theme", "dark");
            document.querySelector('#theme-switch').innerHTML = "Тема: Темная";
        }
    }
    try {
        detectColorScheme();
    } catch {}

    const toggleSwitch = document.querySelector('#checkbox-theme');
    const toggleSwitchText = document.querySelector('#theme-switch');

    function switchTheme(e) {
        if (e.target.checked) {
            localStorage.setItem('theme', 'dark');
            document.documentElement.setAttribute('data-theme', 'dark');
            toggleSwitchText.innerHTML  = "Тема: Темная";
            toggleSwitch.checked = true;
        } else {
            localStorage.setItem('theme', 'light');
            document.documentElement.setAttribute('data-theme', 'light');
            toggleSwitchText.innerHTML  = "Тема: Светлая";
            toggleSwitch.checked = false;
        }
    }
    try {
        toggleSwitch.addEventListener('change', switchTheme, false);
        if (document.documentElement.getAttribute("data-theme") == "dark"){
            toggleSwitch.checked = true;
        }
    } catch { }


    // ################################################################################################################
    // ##                    Profile
    // ################################################################################################################

    const profileBtn = document.getElementById('profileBtn');
    const dropdown = document.querySelector('.profile-dropdown');

    let hideTimeout;

    const showDropdown = () => {
        clearTimeout(hideTimeout);
        dropdown.classList.add('visible');
        profileBtn.classList.add('active');
    };

    const hideDropdown = () => {
        dropdown.classList.remove('visible');
        profileBtn.classList.remove('active');

        hideTimeout = setTimeout(() => {
            dropdown.style.visibility = 'hidden';
            dropdown.style.pointerEvents = 'none';
        }, 300); // такое же, как transition
    };

    const forceVisible = () => {
        dropdown.style.visibility = 'visible';
        dropdown.style.pointerEvents = 'auto';
        profileBtn.classList.add('active');
    };

    try {
        profileBtn.addEventListener('mouseenter', () => {
            forceVisible();
            showDropdown();
        });
        profileBtn.addEventListener('mouseleave', hideDropdown);
        dropdown.addEventListener('mouseenter', showDropdown);
        dropdown.addEventListener('mouseleave', hideDropdown);
    } catch { }
});

