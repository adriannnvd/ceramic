

//  // Get all side buttons
//  const sideButtons = document.querySelectorAll('.side-btn');

//  // Add click event listener to each button
//  sideButtons.forEach(button => {
//      button.addEventListener('click', () => {
//          // Remove 'active' class from previously active button
//          document.querySelector('.side-btn.active').classList.remove('active');
//          // Add 'active' class to the clicked button
//          button.classList.add('active');
//      });
//  });

 document.addEventListener("DOMContentLoaded",() => {
    var current = location.pathname;
    var navitem = document.querySelectorAll('.sidebar-buttons li a');

    navitem.forEach(e => {
        var currentlink = e.getAttribute("href");
        if (currentlink === current) {
            e.classList.add('active')
            if (currentlink.includes('dashboard')) {
                var dashboard = document.getElementById('dashboard');
                dashboard.classList.add('active')
            }
            if (currentlink.includes('company')) {
                var dashboard = document.getElementById('company');
                company.classList.add('active')
            }
            if (currentlink.includes('images')) {
                var images = document.getElementById('images');
                images.classList.add('active')
            }
            if (currentlink.includes('user')) {
                var user = document.getElementById('user');
                user.classList.add('active')
            }
        }
        else {
            e.classList.remove('active')
        }
    });
 });