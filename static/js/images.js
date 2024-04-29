function formatDate(dateString) {
    const options = { month: 'long', day: 'numeric', year: 'numeric', hour: '2-digit', minute: '2-digit' };
    const date = new Date(dateString);
    return date.toLocaleDateString('en-US', options);
}

// Show Form in File Upload
document.getElementById("uploadBtn").addEventListener("click", function() {
    var form = document.getElementById("uploadForm");
    if (form.style.display == "none") {
        form.style.display = "block";
    } else {
        form.style.display = "none";
    }
});