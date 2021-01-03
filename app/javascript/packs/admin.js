require("@rails/ujs").start()
require("turbolinks").start()
require("@rails/activestorage").start()
require("channels")

import * as bootstrap from 'bootstrap/dist/js/bootstrap';
import "../stylesheets/admin";

document.addEventListener("DOMContentLoaded", function (event) {
  var popoverTriggerList = [].slice.call(
    document.querySelectorAll('[data-bs-toggle="popover"]')
  );
  var popoverList = popoverTriggerList.map(function (popoverTriggerEl) {
    return new bootstrap.Popover(popoverTriggerEl);
  });

  var tooltipTriggerList = [].slice.call(
    document.querySelectorAll('[data-bs-toggle="tooltip"]')
  );
  var tooltipList = tooltipTriggerList.map(function (tooltipTriggerEl) {
    return new bootstrap.Tooltip(tooltipTriggerEl);
  });

  document.querySelector('[data-bs-toggle="offcanvas"]').addEventListener('click', function () {
    document.querySelector('.offcanvas-collapse').classList.toggle('open')
  })

});
import "@fortawesome/fontawesome-free/js/all";

// const images = require.context('../images', true)
// const imagePath = (name) => images(name, true)

// $( document ).on('turbolinks:load', function() {
//   console.log("It works on each visit!")
//   console.log(document.querySelector('[data-toggle="offcanvas"]'))
// })