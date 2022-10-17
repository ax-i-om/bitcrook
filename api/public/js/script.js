let myFunction;

myFunction = () => {
  let tnav = document.getElementById("myTopnav");
  if (tnav.className === "topnav") {
    tnav.className += " responsive";
  } else {
    tnav.className = "topnav";
  }
}
