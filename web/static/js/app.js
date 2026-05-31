// ShopOnline - Main JavaScript
$(document).ready(function() {
    console.log('ShopOnline app loaded');

    // HTMX event listeners
    document.body.addEventListener('htmx:afterSwap', function(event) {
        console.log('HTMX swap completed:', event.detail.target.id);
    });

    document.body.addEventListener('htmx:responseError', function(event) {
        console.error('HTMX error:', event.detail.xhr.status);
        alert('An error occurred. Please try again.');
    });
});
