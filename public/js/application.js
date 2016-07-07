$(document).ready(function(){
  console.log("Hello");
  $('#logout').on('click', function(e){
    e.preventDefault();

    $.ajax({
      url: '/sessions',
      type: 'DELETE'
    });

    return false;
  });
});