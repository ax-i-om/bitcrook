/* Toggle between adding and removing the "responsive" class to topnav when the user clicks on the icon */
function myFunction() {
    let tnav = document.getElementById("myTopnav");
    if (tnav.className === "topnav") {
      tnav.className += " responsive";
    } else {
      tnav.className = "topnav";
    }
  }
  