const myFunction = null;

myFunction = () => {
  const tnav = document.getElementById("myTopnav");
  if (tnav.className === "topnav") {
    tnav.className += " responsive";
  } else {
    tnav.className = "topnav";
  }
}
