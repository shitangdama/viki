$(document).on('turbolinks:load', function(){
    console.log(0)
    $("#menu-toggle").click(function(e) {
        console.log(1111)
        e.preventDefault();
        $("#wrapper").toggleClass("toggled");
      });
});
