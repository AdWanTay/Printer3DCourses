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
    detectColorScheme();



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

    toggleSwitch.addEventListener('change', switchTheme, false);

    if (document.documentElement.getAttribute("data-theme") == "dark"){
        toggleSwitch.checked = true;
    }

});

