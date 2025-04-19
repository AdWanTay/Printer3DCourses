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
}

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
}