function formatDate(dateString) {
    const options = { month: 'long', day: 'numeric', year: 'numeric', hour: '2-digit', minute: '2-digit' };
    const date = new Date(dateString);
    return date.toLocaleDateString('en-US', options);
}