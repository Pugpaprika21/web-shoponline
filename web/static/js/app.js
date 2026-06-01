// SweetAlert2 helpers
const Toast = Swal.mixin({
    toast: true, position: 'top-end', showConfirmButton: false, timer: 3000,
    timerProgressBar: true
});

function showSuccess(msg) { Toast.fire({ icon: 'success', title: msg }); }
function showError(msg) { Toast.fire({ icon: 'error', title: msg }); }
function showConfirm(title, text) {
    return Swal.fire({
        title: title, text: text, icon: 'warning', showCancelButton: true,
        confirmButtonColor: '#4F46E5', cancelButtonColor: '#6B7280',
        confirmButtonText: 'Confirm', cancelButtonText: 'Cancel'
    });
}

// AJAX helper with auto token
function api(method, url, data) {
    var opts = {
        url: url, method: method, contentType: 'application/json',
        headers: { 'Authorization': 'Bearer ' + (localStorage.getItem('token') || '') }
    };
    if (data) opts.data = JSON.stringify(data);
    return $.ajax(opts);
}
